package middleware

import (
	"context"
	"net/http"

	"files/internal/shared"

	"github.com/google/uuid"
)

// extracts x-user-id and puts it into context if present
func ActorContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIDHeader := r.Header.Get("X-User-ID")
		if userIDHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		uid, err := uuid.Parse(userIDHeader)
		if err != nil {
			// bad header, but dont block public routes
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), shared.ActorIDKey, uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
