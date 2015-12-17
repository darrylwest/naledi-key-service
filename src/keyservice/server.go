package keyservice

import (
	"net/http"
)

func ConfigureRoutes() *http.ServeMux {
	log.Info("configure service routes")
	mux := http.NewServeMux()

	mux.HandleFunc("/status", StatusHandler)
	mux.HandleFunc("/ping", PingHandler)

	return mux
}
