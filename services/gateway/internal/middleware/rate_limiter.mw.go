package middleware

import (
	"gateway/internal/utils"
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

/*
Rn, all rate limiters must be middleware funcs
goal: create an interface that all rate limiters must implement
- reuse middleware funcs

NewMemoryRateLimiter will return a Middleware object , which can use any rate limiter as its wrapped in middleware Func

*/

type MemoryRateLimiter struct {
	mu        sync.Mutex
	visitors  map[string]int
	limit     int
	resetTime time.Duration
}

func NewMemoryRateLimiter(limit int, resetTime time.Duration) Middleware {
	rl := &MemoryRateLimiter{
		visitors:  make(map[string]int),
		limit:     limit,
		resetTime: resetTime,
	}

	go rl.resetVisitorCount()

	return utils.MiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			rl.mu.Lock()
			defer rl.mu.Unlock()

			var visitorKey string

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
	})
}

func (rl *MemoryRateLimiter) resetVisitorCount() {
	for {
		time.Sleep(rl.resetTime)
		rl.mu.Lock()
		// assign new map to clear old data
		rl.visitors = make(map[string]int)
		rl.mu.Unlock()
	}
}
