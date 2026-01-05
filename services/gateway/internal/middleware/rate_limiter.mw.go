package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"
)

/*
GOAL of this Rate Limiter:
1. Find if the UserID exists for this requests
2. If UserID exists then use it as the key for counting requests
3. If UserID does not exist then increase the request count wrt IP address
*/

// TODO: switch rate limiter to Redis

type RateLimiter interface {
	RateLimiterMW(next http.Handler) http.Handler
}

type rateLimiter struct {
	mu        sync.Mutex
	visitors  map[string]int
	limit     int
	resetTime time.Duration
}

func NewRateLimiter(limit int, resetTime time.Duration) *rateLimiter {
	rl := &rateLimiter{
		visitors:  make(map[string]int),
		limit:     limit,
		resetTime: resetTime,
	}
	// start the reset routine
	go rl.resetVisitorCount()
	return rl
}

func (rl *rateLimiter) resetVisitorCount() {
	for {
		time.Sleep(rl.resetTime)
		rl.mu.Lock()
		// assign new map to clear old data
		rl.visitors = make(map[string]int)
		rl.mu.Unlock()
	}
}

func (rl *rateLimiter) RateLimiterMW(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rl.mu.Lock()
		defer rl.mu.Unlock()

		var visitorKey string
		// r.Context().Value("UserID") returns a value of type any, so we must assert it to string before using it as a key

		if userID, ok := r.Context().Value("UserID").(string); ok && userID != "" {
			visitorKey = userID
		} else {
			host, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				host = r.RemoteAddr
			}
			visitorKey = host
		}

		rl.visitors[visitorKey]++

		if rl.visitors[visitorKey] > rl.limit {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
