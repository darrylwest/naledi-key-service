package keyservicetest

import (
	"keyservice"
	"testing"

	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/darrylwest/cassava-logger/logger"
	"net/http"
	"net/http/httptest"

	. "github.com/franela/goblin"
)

func TestMiddleware(t *testing.T) {
	g := Goblin(t)

	g.Describe("Middleware", func() {
		log := func() *logger.Logger {
			ctx := keyservice.NewContextForEnvironment("test")
			return ctx.CreateLogger()
		}()

		g.It("should accept a https proto request", func() {
			log.Info("test the https proto request")

			msg := "write success to service"
			recorder := httptest.NewRecorder()
			ctx := keyservice.NewDefaultContext()

			proto := keyservice.NewProtoMiddleware(ctx)

			server := negroni.New()
			server.Use(proto)

			server.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				// fmt.Println( req.Header.Get("X-Forwarded-Proto"))
				fmt.Fprintf(w, msg)
			}))

			request, err := http.NewRequest("GET", "https://bluelasso.com/mypage", nil)
			request.Header.Set("X-Forwarded-Proto", "https")

			g.Assert(err == nil).IsTrue()

			server.ServeHTTP(recorder, request)

			// fmt.Printf("%d - %s", recorder.Code, recorder.Body.String())
			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body.String()).Equal(msg)
		})

		g.It("should reject a non-https request", func() {
			msg := "write success to service"
			recorder := httptest.NewRecorder()
			ctx := keyservice.NewDefaultContext()

			proto := keyservice.NewProtoMiddleware(ctx)

			server := negroni.New()
			server.Use(proto)

			server.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				// fmt.Println( req.Header.Get("X-Forwarded-Proto"))
				fmt.Fprintf(w, msg)
			}))

			request, err := http.NewRequest("GET", "https://bluelasso.com/mypage", nil)
			request.Header.Set("X-Forwarded-Proto", "http")

			g.Assert(err == nil).IsTrue()

			server.ServeHTTP(recorder, request)

			// fmt.Printf("%d - %s", recorder.Code, recorder.Body.String())
			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body.String() != msg)
		})

		g.It("should accept a non-https request for localhost", func() {
			msg := "write success to service for status"
			recorder := httptest.NewRecorder()
			ctx := keyservice.NewDefaultContext()

			proto := keyservice.NewProtoMiddleware(ctx)

			server := negroni.New()
			server.Use(proto)

			server.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				// fmt.Println( req.Header.Get("X-Forwarded-Proto"))
				fmt.Fprintf(w, msg)
			}))

			request, err := http.NewRequest("GET", "https://localhost:3434/status", nil)
			request.Header.Set("X-Forwarded-Proto", "http")

			g.Assert(err == nil).IsTrue()

			server.ServeHTTP(recorder, request)

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body.String()).Equal(msg)
		})

		g.It("should accept a non-https request for non-production", func() {
			msg := "write success to service for non-production env"
			recorder := httptest.NewRecorder()
			ctx := keyservice.NewContextForEnvironment("test")

			proto := keyservice.NewProtoMiddleware(ctx)

			server := negroni.New()
			server.Use(proto)

			server.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				// fmt.Println( req.Header.Get("X-Forwarded-Proto"))
				fmt.Fprintf(w, msg)
			}))

			request, err := http.NewRequest("GET", "https://test.com/anypage", nil)
			request.Header.Set("X-Forwarded-Proto", "http")

			g.Assert(err == nil).IsTrue()

			server.ServeHTTP(recorder, request)

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body.String()).Equal(msg)
		})

		g.It("should accept known apikey", func() {
			msg := "write success to service with apikey"
			recorder := httptest.NewRecorder()
			ctx := keyservice.NewDefaultContext()

			akm := keyservice.NewAPIKeyMiddleware(ctx)

			server := negroni.New()
			server.Use(akm)

			server.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				// fmt.Println( req.Header.Get("X-Forwarded-Proto"))
				fmt.Fprintf(w, msg)
			}))

			request, err := http.NewRequest("GET", "http://test.com/session", nil)
			request.Header.Set("x-api-key", "c2b4d9bf-652e-4915-ab23-7a0e0e32e362")

			g.Assert(err == nil).IsTrue()

			server.ServeHTTP(recorder, request)

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body.String()).Equal(msg)
		})

		g.It("should reject unknown apikey", func() {
			msg := "invalid request...\r\n"
			recorder := httptest.NewRecorder()
			ctx := keyservice.NewDefaultContext()

			akm := keyservice.NewAPIKeyMiddleware(ctx)

			server := negroni.New()
			server.Use(akm)

			server.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				// fmt.Println( req.Header.Get("X-Forwarded-Proto"))
				fmt.Fprintf(w, msg)
			}))

			request, err := http.NewRequest("GET", "http://test.com/session", nil)
			request.Header.Set("x-api-key", "c-bad-key2")

			g.Assert(err == nil).IsTrue()

			server.ServeHTTP(recorder, request)

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body.String()).Equal(msg)
		})
	})
}
