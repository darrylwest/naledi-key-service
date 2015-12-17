package keyservice

import (
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/darrylwest/cassava-logger/logger"
	// "gopkg.in/tylerb/graceful.v1"
)

func ConfigureRoutes() *http.ServeMux {
	log.Info("configure service routes")
	mux := http.NewServeMux()

	mux.HandleFunc("/status", StatusHandler)
	mux.HandleFunc("/ping", PingHandler)

	// now for the custom routes

	return mux
}

func CreateService(mux *http.ServeMux, env string) *negroni.Negroni {
	server := negroni.New()

	// assign the standard middleware
	server.Use(negroni.NewRecovery())
	server.Use(logger.NewMiddlewareLogger(log))
	// server.Use(NewProtoMiddleware(env))

	// server.Use(gzip.Gzip(gzip.DefaultCompression))
	// server.Use(negroni.NewStatic(http.Dir(webroot)))

	server.UseHandler(mux)

	return server
}
