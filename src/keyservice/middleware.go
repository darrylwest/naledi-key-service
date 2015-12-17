package keyservice

import (
	"fmt"
	"net/http"
	"strings"
)

type ProtoMiddleware struct {
	allowHTTP bool
}

func NewProtoMiddleware(ctx *Context) *ProtoMiddleware {
	return &ProtoMiddleware{allowHTTP: IsProduction(ctx.env) == false}
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
	skip   bool
	apikey string
}

func NewAPIKeyMiddleware(ctx *Context) *APIKeyMiddleware {
	m := &APIKeyMiddleware{}

	m.skip = IsProduction(ctx.env) == false
	if ctx.config != nil {
		m.apikey = ctx.config.appkey
		log.Info("apikey: %s", m.apikey)
	} else {
		m.apikey = "c2b4d9bf-652e-4915-ab23-7a0e0e32e362"
		log.Warn("using the default api key: %s", m.apikey)
	}

	return m
}

func (m *APIKeyMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	key := r.Header.Get("x-api-key")
	log.Info("api key check: %s, len: %d", key, len( key ))

	if m.apikey == key || (len( key ) >= 32 && m.skip)  {
		next(w, r)
	} else {
		log.Warn("request header does not have a recognized key: %s", key)

		headers := w.Header()
		headers.Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "invalid request...\r\n")
	}

}
