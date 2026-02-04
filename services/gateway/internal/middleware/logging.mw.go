package middleware

import (
	"context"
	"gateway/internal/utils"
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func NewLogger() Middleware {
	return utils.MiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()

			rw := &responseWriter{
				ResponseWriter: w,
				status:         http.StatusOK, // default
			}

			next.ServeHTTP(rw, r)

			reqID := GetRequestID(r.Context())

			log.Printf(
				"request_id=%s method=%s path=%s status=%d duration=%s",
				reqID,
				r.Method,
				r.URL.Path,
				rw.status,
				time.Since(start),
			)
		})
	})
}

func GetRequestID(ctx context.Context) string {
	if val := ctx.Value(utils.RequestIDKey); val != nil {
		if id, ok := val.(string); ok {
			return id
		}
	}
	return "unknown"
}
