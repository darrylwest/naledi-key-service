package keyservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	// "strings"
	"time"
)

func badRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bad request\r\n")
}

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		badRequestHandler(w, r)
		return
	}

	fmt.Fprintf(w, "expire session not implemented yet\r\n")
}

func ExpireSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		badRequestHandler(w, r)
		return
	}

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
