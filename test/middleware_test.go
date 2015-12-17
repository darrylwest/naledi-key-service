package keyservicetest

import (
    "testing"
	"keyservice"

	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
	"net/http/httptest"
    "github.com/darrylwest/cassava-logger/logger"

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

			proto := keyservice.NewProtoMiddleware("production")

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

			proto := keyservice.NewProtoMiddleware("production")

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

		g.It("should accept a non-https request for status", func() {
			msg := "write success to service for status"
			recorder := httptest.NewRecorder()

			proto := keyservice.NewProtoMiddleware("production")

			server := negroni.New()
			server.Use(proto)

			server.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				// fmt.Println( req.Header.Get("X-Forwarded-Proto"))
				fmt.Fprintf(w, msg)
			}))

			request, err := http.NewRequest("GET", "https://bluelasso.com/status", nil)
			request.Header.Set("X-Forwarded-Proto", "http")

			g.Assert(err == nil).IsTrue()

			server.ServeHTTP(recorder, request)

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body.String()).Equal(msg)
		})

		g.It("should accept a non-https request for ping", func() {
			msg := "write success to service for ping"
			recorder := httptest.NewRecorder()

			proto := keyservice.NewProtoMiddleware("production")

			server := negroni.New()
			server.Use(proto)

			server.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				// fmt.Println( req.Header.Get("X-Forwarded-Proto"))
				fmt.Fprintf(w, msg)
			}))

			request, err := http.NewRequest("GET", "https://bluelasso.com/ping", nil)
			request.Header.Set("X-Forwarded-Proto", "http")

			g.Assert(err == nil).IsTrue()

			server.ServeHTTP(recorder, request)

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body.String()).Equal(msg)
		})

		g.It("should accept a non-https request for non-production", func() {
			msg := "write success to service for non-production env"
			recorder := httptest.NewRecorder()

			proto := keyservice.NewProtoMiddleware("test")

			server := negroni.New()
			server.Use(proto)

			server.UseHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				// fmt.Println( req.Header.Get("X-Forwarded-Proto"))
				fmt.Fprintf(w, msg)
			}))

			request, err := http.NewRequest("GET", "https://bluelasso.com/anypage", nil)
			request.Header.Set("X-Forwarded-Proto", "http")

			g.Assert(err == nil).IsTrue()

			server.ServeHTTP(recorder, request)

			g.Assert(recorder.Code).Equal(200)
			g.Assert(recorder.Body.String()).Equal(msg)
		})
	})
}
