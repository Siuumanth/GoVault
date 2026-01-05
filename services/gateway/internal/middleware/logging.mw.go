package middleware

import "net/http"

func LoggingMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Add logging logic
		next.ServeHTTP(w, r)
	})
}
