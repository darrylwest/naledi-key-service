package keyservice

import (
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/darrylwest/cassava-logger/logger"
	"gopkg.in/tylerb/graceful.v1"
	"fmt"
)

func ConfigureStandardRoutes() *http.ServeMux {
	log.Info("configure service routes")
	mux := http.NewServeMux()

	mux.HandleFunc("/status", StatusHandler)
	mux.HandleFunc("/ping", PingHandler)

	// now for the custom routes

	return mux
}

func ConfigureCustomRoutes(mux *http.ServeMux) {
	log.Info("configure custom routes")

}

func CreateServer(mux *http.ServeMux, ctx Context) *negroni.Negroni {
	server := negroni.New()

	// assign the standard middleware
	server.Use(negroni.NewRecovery())
	server.Use(logger.NewMiddlewareLogger(log))
	server.Use(NewProtoMiddleware(ctx.env))
	server.Use(NewAPIKeyMiddleware(ctx.apikey))

	// server.Use(gzip.Gzip(gzip.DefaultCompression))
	// server.Use(negroni.NewStatic(http.Dir(webroot)))

	server.UseHandler(mux)

	return server
}

func startServer(server *negroni.Negroni, port int) {
	log.Info("starting server at port: %d", port)
	graceful.Run(fmt.Sprintf(":%v", port), 0, server)
}
