package middleware

import (
	"log"
	"net/http"
)

// extracts x-user-id and puts it into context if present
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("[REQUEST] id=%s path=%s", r.Header.Get("X-Request-ID"), r.URL.Path)
	})
}
