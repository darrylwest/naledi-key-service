package keyservice

import (
	"fmt"
	"net/http"
	"strings"
)

type ProtoMiddleware struct {
	allowHTTP bool
}

func NewProtoMiddleware(env string) *ProtoMiddleware {
	return &ProtoMiddleware{allowHTTP: IsProduction(env) == false}
}

func (m *ProtoMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if m.allowHTTP || r.URL.Path == "/ping" || r.URL.Path == "/status" {
		next(w, r)
	} else {
		proto := r.Header.Get("X-Forwarded-Proto")

		if strings.HasPrefix(proto, "https") {
			next(w, r)
		} else {
			log.Warn("request was not https: %s from %s", proto, r.Header.Get("X-Forwarded-For"))

			headers := w.Header()
			headers.Set("Content-Type", "text/plain")
			fmt.Fprintf(w, "invalid request...\r\n")
		}
	}
}

type APIKeyMiddleware struct {
	apikey string
}

func NewAPIKeyMiddleware(key string) *APIKeyMiddleware {
	return &APIKeyMiddleware{ apikey:key }
}

func (m *APIKeyMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	key := r.Header.Get("x-api-key")
	log.Info("api key check: %s", key)

	if m.apikey == key || r.URL.Path == "/ping" || r.URL.Path == "/status" {
		next(w, r)
	} else {
		log.Warn("request header does not have a recognized key: %s", key)

		headers := w.Header()
		headers.Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "invalid request...\r\n")
	}

}
