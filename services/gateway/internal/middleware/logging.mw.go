package middleware

import (
	"gateway/internal/metrics"
	"gateway/internal/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	body   []byte
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body = append(rw.body, b...)
	return rw.ResponseWriter.Write(b)
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

			duration := time.Since(start).Seconds()
			statusStr := strconv.Itoa(rw.status)

			// ---- METRICS ----
			metrics.HttpRequestsTotal.
				WithLabelValues(r.Method, r.URL.Path, statusStr).
				Inc()

			metrics.HttpRequestDuration.
				WithLabelValues(r.Method, r.URL.Path).
				Observe(duration)

			metrics.HttpInFlight.Dec()
			// ------------------

			reqID := GetRequestID(r.Context())

			if rw.status >= 400 {
				log.Printf(
					"request_id=%s method=%s path=%s status=%d duration=%s error=%s",
					reqID,
					r.Method,
					r.URL.Path,
					rw.status,
					time.Since(start),
					string(rw.body),
				)
			} else {
				log.Printf(
					"request_id=%s method=%s path=%s status=%d duration=%s",
					reqID,
					r.Method,
					r.URL.Path,
					rw.status,
					time.Since(start),
				)
			}

		})
	})
}
