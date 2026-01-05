package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func MapClaims(claims jwt.MapClaims) (*AuthContext, error) {
	// only user ID is mandatory , otheres arent necessary
	uid, ok := claims["uid"].(string)
	if !ok {
		return nil, errors.New("uid missing or invalid")
	}

	email, _ := claims["email"].(string)
	username, _ := claims["user"].(string)

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
