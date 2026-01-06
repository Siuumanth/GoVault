package middleware

import (
	"fmt"
	"gateway/internal/utils"
	"net/http"
	"os"
)

// Allowed origins
var allowedOrigins = []string{
	os.Getenv("FRONTEND_URL"),
	os.Getenv("DEV_URL"),
	os.Getenv("OTHER_URL"),
}

func NewCORS() Middleware {
	return utils.MiddlewareFunc(func(next http.Handler) http.Handler {
		// Take that function value and treat it as a MiddlewareFunc
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // now tihs function becomes a http.handler because we are wrapping it in http.HanderFunc
			origin := r.Header.Get("Origin")
			// fmt.Println(origin)

			if isOriginAllowed(origin) {
				w.Header().Set("Access-Control-Allow-Origin", origin)

			} else {
				http.Error(w, "Not allowed by CORS", http.StatusForbidden)
				return
			}

			// Set other CORS headers
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Max-Age", "3600")

			// Handle preflight request
			if r.Method == http.MethodOptions {
				return
			}

			next.ServeHTTP(w, r)
			fmt.Println("Cors Middleware ends...")
		})
	})
}

func isOriginAllowed(origin string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}
	return false
}
