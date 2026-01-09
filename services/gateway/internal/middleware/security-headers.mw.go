package middleware

import (
	"gateway/internal/utils"
	"net/http"
)

func NewSecurityHeaders() Middleware {
	return utils.MiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
			w.Header().Set("Cache-Control", "no-store")

			next.ServeHTTP(w, r)
		})
	})
}
