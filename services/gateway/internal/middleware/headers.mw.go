package middleware

import (
	"gateway/internal/utils"
	"net/http"
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
*/

// Context Keys 
var (
    userID   = "UserID"
	username = "Username"
	email    = "Email"
	expires  = "Expires"
)


var authCtx utils.AuthContext

func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r* http.Request){
        authCtx = utils.GetAuthContext(r.Context())
		if()
		next.ServeHTTP(w, r)
	})
}