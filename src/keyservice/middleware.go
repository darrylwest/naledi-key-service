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
	if m.allowHTTP || strings.Contains(r.Host, "localhost") {
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
	m.apikey = ctx.apikey

	log.Info("skip: %v, apikey: %s", m.skip, m.apikey)

	return m
}

func (m *APIKeyMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	key := r.Header.Get("x-api-key")
	log.Info("api key check: %s, len: %d", key, len( key ))

	if m.apikey == key || (len( key ) >= 32 && m.skip) || strings.Contains(r.Host, "localhost")  {
		next(w, r)
	} else {
		log.Warn("request header does not have a recognized key: %s", key)

		headers := w.Header()
		headers.Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "invalid request...\r\n")
	}

}
