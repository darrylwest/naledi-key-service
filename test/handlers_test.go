package keyservicetest

import (
	"keyservice"
	"testing"
	"fmt"

	// "code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"net/http"
	"net/http/httptest"
	"net/url"
	// "bytes"

	. "github.com/franela/goblin"
)

func TestHandlers(t *testing.T) {
	g := Goblin(t)



	g.Describe("Handlers", func() {
		g.It("should create a new session", func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/session/create", keyservice.CreateSessionHandler)

			server := negroni.New()
			server.UseHandler(mux)

			request, err := http.NewRequest("POST", "http://test.com/session/create", nil)
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

		g.It("should create a new session", func() {
			mux := http.NewServeMux()
			mux.HandleFunc("/KeyService/session/create", keyservice.CreateSessionHandler)

			server := negroni.New()
			server.UseHandler( mux )

			values := url.Values{}
			values.Add("image", "flarb")

			request, err := http.NewRequest("POST", "http://test.com/KeyService/session/create", nil)
			if err != nil {
				panic( err )
			}

			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			recorder := httptest.NewRecorder()

			server.ServeHTTP(recorder, request)

			fmt.Println( recorder.Body.String() )

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body != nil).IsTrue()
		})
	})
}

func createNewSessionRequest() []byte {
	// a known server box pub key
	peerpub := "587b2d753c8409bbf876e7f9dc682b01a411cdd2ce6f0c66046d69c6343c1a1d"

	return []byte("this is my test " + peerpub)
}
