package utils

import (
	"context"
	"time"
)

type contextKey string

const (
	AuthContextKey contextKey = "auth"
)

// this is for claiming from the JWt
const (
	ClaimsUidKey      string = "uid"
	ClaimsExpKey      string = "exp"
	ClaimsEmailKey    string = "email"
	ClaimsUsernameKey string = "username"
)

type AuthContext struct {
	UserID   string
	Username string
	Email    string
	Expires  time.Time
}

func GetAuthContext(ctx context.Context) AuthContext {
	return ctx.Value(AuthContextKey).(AuthContext)
}
