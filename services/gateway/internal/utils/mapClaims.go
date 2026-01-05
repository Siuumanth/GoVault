package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func MapClaims(claims jwt.MapClaims) (*AuthContext, error) {
	uid, ok := claims["uid"].(string)
	if !ok {
		return nil, errors.New("uid missing or invalid")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("email missing or invalid")
	}

	username, ok := claims["user"].(string)
	if !ok {
		return nil, errors.New("user missing or invalid")
	}

	expFloat, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("exp missing or invalid")
	}

	return &AuthContext{
		UserID:   uid,
		Username: username,
		Email:    email,
		Expires:  time.Unix(int64(expFloat), 0),
	}, nil
}
