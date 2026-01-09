package middleware

import (
	"gateway/internal/utils"
	"net/http"
)

type Logger struct{}

func NewLogger() Middleware {
	return utils.MiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Add logging logic
			next.ServeHTTP(w, r)
		})
	})

}
