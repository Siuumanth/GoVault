package utils

import (
	"context"
	"time"
)

type AuthContext struct {
	UserID   string
	Username string
	Email    string
	Expires  time.Time
}

// authContextKey
var authContextKey string = "auth"

func GetAuthContext(ctx context.Context) AuthContext {
	return ctx.Value(authContextKey).(AuthContext)
}
