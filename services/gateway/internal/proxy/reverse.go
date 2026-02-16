package proxy

import (
	"fmt"
	"gateway/internal/middleware"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/sony/gobreaker"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rw *statusRecorder) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func NewReverseProxy(target string, serviceName string) http.Handler {
	u, err := url.Parse(target)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(u)

	settings := gobreaker.Settings{
		Name:        serviceName,
		MaxRequests: 5,                // max trial req in half open state
		Interval:    10 * time.Second, // count in intervals of this sec
		Timeout:     30 * time.Second, // wait this much after open to go to half open
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			if counts.Requests < 10 {
				return false
			}
			ratio := float64(counts.TotalFailures) / float64(counts.Requests)
			return ratio >= 0.6
		},
		// trip when 60 percent of requests fail
		OnStateChange: func(name string, from, to gobreaker.State) {
			log.Printf("[CB:%s] %s → %s", name, from.String(), to.String())
		}, // log wen state change
	}

	cb := gobreaker.NewCircuitBreaker(settings)

	// Capture upstream transport errors (dial, timeout, etc.)
	// Upstream error happens when the gateway cannot even successfully talk to the backend service.
	// default respnose is 502, we just handle it and log it how we want here
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		reqID := middleware.GetRequestID(r.Context())
		log.Printf("[req_id=%s] [CB:%s] upstream error: %v", reqID, serviceName, err)

		fmt.Fprintf(w, "Upstream error")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Fail fast if breaker is OPEN
		if cb.State() == gobreaker.StateOpen {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "%s unavailable (circuit open)", serviceName)
			return
		}

		rw := &statusRecorder{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		// execute forwarding logic
		_, err := cb.Execute(func() (interface{}, error) {
			proxy.ServeHTTP(rw, r)

			// Trip breaker only on 5xx
			if rw.status >= 500 {
				return nil, fmt.Errorf("upstream 5xx: %d", rw.status)
			}

			return nil, nil
		})

		// returning proper request if circuit breaker is open, not just 5xx
		if err != nil {
			if err == gobreaker.ErrOpenState {
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Fprintf(w, "%s unavailable (circuit open)", serviceName)
				return
			}

			// Upstream already responded; just log
			log.Printf("[CB:%s] execution error: %v", serviceName, err)
		}
	})
}
