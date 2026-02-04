package middleware

import (
	"context"
	"files/internal/shared"
	"net/http"

	"github.com/google/uuid"
)

// for routers
func RequireActor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, ok := r.Context().Value(shared.ActorIDKey).(uuid.UUID)
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(
			context.WithValue(r.Context(), shared.ActorIDKey, id),
		))
	})
}
