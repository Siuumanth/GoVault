package middleware

import "net/http"

type Logger struct{}

func (l Logger) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Add logging logic
		next.ServeHTTP(w, r)
	})
}
