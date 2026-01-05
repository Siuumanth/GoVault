package middleware

import (
	"gateway/internal/utils"
	"net/http"
	"time"
)

/*
Goal of this middleware:

type AuthContext struct {
	UserID   string
	Username string
	Email    string
	Expires  time.Time
}

get these context values and addd them to the header
- main thing is userID and
*/

type HeaderInjectionMW struct{}

func (h HeaderInjectionMW) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCtx := utils.GetAuthContext(r.Context())

		// we assume the previous MW have handled missing fields
		r.Header.Set("X-User-ID", authCtx.UserID)
		// username email are not needed right now, can add later

		if authCtx.Expires.IsZero() == false {
			r.Header.Set("X-Auth-Expires", authCtx.Expires.UTC().Format(time.RFC3339))
		}

		next.ServeHTTP(w, r)
	})
}
