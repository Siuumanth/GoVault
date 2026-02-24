package middleware

import (
	"context"
	"net/http"
	"upload/shared"

	"github.com/google/uuid"
)

// extracts x-user-id and puts it into context if present

func ValidateUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIDHeader := r.Header.Get("X-User-ID")
		if userIDHeader == "" {
			http.Error(w, "Unauthorized: Missing User ID", http.StatusUnauthorized)
			return
		}

		uid, err := uuid.Parse(userIDHeader)
		if err != nil {
			http.Error(w, "Bad Request: Invalid User ID format", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), shared.ActorIDKey, uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
