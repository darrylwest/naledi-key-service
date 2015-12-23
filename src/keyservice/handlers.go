package keyservice

import (
	"github.com/agl/ed25519"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	// "strings"
	"time"
)

func findPrivateKey(pub *[KeySize]byte) (*[KeySize]byte,error) {
	// TODO lookup the key from the key store
	pk, _ := hex.DecodeString("1d4f58f4f1e40c72dc695836902119ac553b84693904efac931731ae2ea27b48")

	k := new([KeySize]byte)
	copy(k[:], pk)

	return k, nil
}

func badRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bad request\r\n")
}

func sendNewSessionResponse(w http.ResponseWriter, r *http.Request, session *Session) {
	w.Header().Set("Content-Type", "text/plain")

	sigpub, sigpriv, err := ed25519.GenerateKey( rand.Reader )
	if err != nil {
		log.Error("could not generate signature keys: ", err)
		RemoveSession(session.ssid)
		badRequestHandler(w, r)
		return
	}

	msg := new(Message)
	msg.MyKey = session.serverPub
	msg.YourKey = session.clientPub
	msg.SignatureKey = sigpub
	msg.Number = session.messageCount

	now := time.Now().Unix()
	// todo create a json blob for return message...
	str := fmt.Sprintf("ssid=%s,expires:%d\r\n", session.ssid,session.expires - now)
	enc, err := EncryptBox(session.clientPub, session.serverPriv, []byte( str ) )
	if err != nil {
		log.Error("could not box encode: ", err)
		RemoveSession(session.ssid)
		badRequestHandler(w, r)
		return
	}

	msg.EncryptedMessage = &enc
	msg.Signature = ed25519.Sign(sigpriv, *msg.EncryptedMessage)

	mbody, err := msg.EncodeToString()

	fmt.Fprintf(w, mbody)
}

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" || r.ContentLength < 420 {
		log.Warn("bad session request: ", r)
		badRequestHandler(w, r)
		return
	}

	body := make([]byte, r.ContentLength)
	n, err := r.Body.Read(body)
	if err != nil {
		log.Error("bad session request, body read error: ", err)
		badRequestHandler(w, r)
		return
	}

	log.Info("bytes read: %d from body: %s", n, body)

	msg, err := DecodeMessageFromString( string(body) )
	if err != nil {
		log.Warn("error decoding message: %v", err)
		badRequestHandler(w, r)
		return
	}

	if msg.Number != 1 {
		log.Warn("message number incorrect: %d", msg.Number)
		badRequestHandler(w, r)
		return
	}

	if !ed25519.Verify(msg.SignatureKey, *msg.EncryptedMessage, msg.Signature) {
		log.Warn("message signature incorrect: %s", msg.EncryptedMessage)
		badRequestHandler(w, r)
		return
	}

	priv, err := findPrivateKey(msg.YourKey)
	if err != nil {
		log.Warn("could not locate cached key %s", msg.YourKey)
		badRequestHandler(w, r)
		return
	}

	license, err := DecryptBox(msg.MyKey, priv, *msg.EncryptedMessage)
	if err != nil {
		log.Warn("could not decrypt box message %s", msg.EncryptedMessage)
		badRequestHandler(w, r)
		return
	}

	log.Info( string(license) )
	// TODO validate the license key...

	session := CreateSession(int64(0))
	if session == nil {
		log.Error("could not create a new session")
		badRequestHandler(w, r)
		return
	}
	session.clientPub = msg.MyKey
	session.messageCount = 1

	sendNewSessionResponse(w, r, session)
}

func ExpireSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		badRequestHandler(w, r)
		return
	}

	// pull the ssid and remove from sessions

	fmt.Fprintf(w, "expire session not implemented yet\r\n")
}

func ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "shutdown requested...\r\n")

	// check for post and token
	if r.Method == "POST" {
		log.Info("shutdown in a graceful way...\r\n")

		// TODO replace with internal signal listener
		cmd := exec.Command("kill", "-2", fmt.Sprintf("%d", os.Getpid()))
		cmd.Run()

		log.Info("shutdown running...")
	} else {
		log.Warn("shudown denied, method %s", r.Method)
		fmt.Fprintf(w, "shutdown request denied...\r\n")
	}
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	var m = map[string]interface{}{
		"status":  "ok",
		"ts":      time.Now().UnixNano() / 1000000,
		"version": "1.0",
		"webStatus": map[string]interface{}{
			"version":         version,
			"pid":             os.Getpid(),
			"proto":           r.Proto,
			"host":            r.Host,
			"path":            r.URL.Path,
			"agent":           r.UserAgent(),
			"remoteAddr":      r.RemoteAddr,
			"xForwardedFor":   r.Header.Get("X-Forwarded-For"),
			"xForwardedProto": r.Header.Get("X-Forwarded-Proto"),
		},
	}

	json, err := json.Marshal(m)

	if err != nil {
		fmt.Fprintf(w, "json error\r\n")
	} else {
		headers := w.Header()
		headers.Set("Content-Type", "application/json")
		log.Debug("headers: %v", headers)

		w.Write(json)
		w.Write([]byte("\r\n"))
	}
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\r\n"))
}
