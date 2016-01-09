package keyservicetest

import (
	"fmt"
	"keyservice"
	"keyservice/models"
	"testing"

	"bytes"
	"encoding/json"
	"github.com/agl/ed25519"
	"github.com/codegangsta/negroni"
	"net/http"
	"net/http/httptest"

	. "github.com/franela/goblin"
)

func TestHandlers(t *testing.T) {
	g := Goblin(t)

	g.Describe("Handlers", func() {
		g.It("should create a new session", func() {
			keyservice.PurgeAllSessions()
			body := "2702a0c5ca762b9da8270c7fc4bd5ef73b24bafc7a2259bac85955a43b3d7aaf:4d6c4804d618d086670a258fcb09e3028adda4051aeffd533c1570c9a89ee7f5dcaba332611ba12283a2eb18d741e8211aff332114cad2b5df4d87865670f60b:1:44a840a2a69983b209f703ee097dd1080c07f7d8daa5fe4ec890ec8a658b8f40:587b2d753c8409bbf876e7f9dc682b01a411cdd2ce6f0c66046d69c6343c1a1d:881842cdbca0e3d394d63efae266d3cbe1c5e535570b2f43891c5c1be0b965a992ce9dd76c56fb35e8b8573db967af7b2e1d1ef54e6ca20797a9"

			mux := http.NewServeMux()
			mux.HandleFunc("/session/create", keyservice.CreateSessionHandler)

			server := negroni.New()
			server.UseHandler(mux)

			request, err := http.NewRequest("POST", "http://test.com/session/create", bytes.NewBufferString(body))
			g.Assert(err == nil).IsTrue()

			request.Header.Set("x-api-key", "my-api-key")
			request.Header.Set("Content-Type", "text/plain")
			request.Header.Set("Accept", "text/plain")

			recorder := httptest.NewRecorder()

			server.ServeHTTP(recorder, request)
			fmt.Println(recorder.Body.String())
			// log.Info("status response: %s", recorder.Body.String())

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body != nil).IsTrue()

			sessions := keyservice.GetSessions()
			g.Assert(sessions.Len()).Equal(1)

			// now decrypt the message to verify the session, expires, EncryptSymmetric
			msg, err := models.DecodeMessageFromString(recorder.Body.String())
			g.Assert(err == nil).IsTrue()
			g.Assert(msg != nil).IsTrue()

			g.Assert(msg.YourKey != nil).IsTrue()
			g.Assert(msg.MyKey != nil).IsTrue()
			g.Assert(msg.SignatureKey != nil).IsTrue()
			g.Assert(msg.Signature != nil).IsTrue()
			g.Assert(msg.EncryptedMessage != nil).IsTrue()
			g.Assert(msg.Number).Equal(1)

			g.Assert(ed25519.Verify(msg.SignatureKey, *msg.EncryptedMessage, msg.Signature)).IsTrue()

			// TODO : create a session and test decryption against sendNewSessionResponse
			// plain, err := keyservice.DecryptBox( msg.YourKey, )
		})

		g.It("should send a new session response with encrypted message")

		g.It("should expire an existing session", func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/session/expire", keyservice.CreateSessionHandler)

			server := negroni.New()
			server.UseHandler(mux)

			request, err := http.NewRequest("PUT", "http://test.com/session/expire", nil)
			if err != nil {
				panic(err)
			}

			recorder := httptest.NewRecorder()

			server.ServeHTTP(recorder, request)
			// fmt.Println( recorder.Body.String() )
			// log.Info("status response: %s", recorder.Body.String())

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body != nil).IsTrue()
		})

		g.It("should have a status handler that returns a json blob", func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/status", keyservice.StatusHandler)

			server := negroni.New()
			server.UseHandler(mux)

			request, err := http.NewRequest("GET", "http://test.com/status", nil)
			if err != nil {
				panic(err)
			}

			recorder := httptest.NewRecorder()

			server.ServeHTTP(recorder, request)
			// fmt.Println( recorder.Body.String() )
			// log.Info("status response: %s", recorder.Body.String())

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body != nil).IsTrue()

			wrapper := make(map[string]interface{})
			err = json.Unmarshal(recorder.Body.Bytes(), &wrapper)

			g.Assert(err == nil).IsTrue()
			status := wrapper["status"]
			g.Assert(status).Equal("ok")
		})

		g.It("should have a ping handler that returns pong", func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/ping", keyservice.PingHandler)

			server := negroni.New()
			server.UseHandler(mux)

			request, err := http.NewRequest("GET", "http://test.com/ping", nil)
			if err != nil {
				panic(err)
			}

			recorder := httptest.NewRecorder()

			server.ServeHTTP(recorder, request)
			// log.Info("status response: %s", recorder.Body.String())

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body != nil).IsTrue()
			g.Assert(recorder.Body.String()).Equal("pong\r\n")
		})

		g.It("should have a shutdown handler that fails if method is not a post", func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/shutdown", keyservice.ShutdownHandler)

			server := negroni.New()
			server.UseHandler(mux)

			request, err := http.NewRequest("GET", "http://test.com/shutdown", nil)
			if err != nil {
				panic(err)
			}

			recorder := httptest.NewRecorder()

			server.ServeHTTP(recorder, request)

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body != nil).IsTrue()
			g.Assert(recorder.Body.String()).Equal("shutdown requested...\r\nshutdown request denied...\r\n")
		})
	})
}

func createNewSessionRequest() []byte {
	// a known server box pub key
	peerpub := "587b2d753c8409bbf876e7f9dc682b01a411cdd2ce6f0c66046d69c6343c1a1d"

	return []byte("this is my test " + peerpub)
}
