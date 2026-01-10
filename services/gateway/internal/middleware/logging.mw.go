package middleware

import (
	"fmt"
	"gateway/internal/utils"
	"net/http"
)

type Logger struct{}

func NewLogger() Middleware {
	return utils.MiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Add logging logic
			fmt.Printf("Request URL: %s\n", r.URL)

			next.ServeHTTP(w, r)
		})
	})

}
