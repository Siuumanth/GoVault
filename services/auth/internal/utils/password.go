package utils

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// hash password using sha256
func HashPassword(password string) (string, error) {
	// SHA-256 doesn't actually return an error on Write, but for clean code:
	h := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", h), nil
}

// switch to bcrypt
//
//	func HashPassword(password string) (string, error) {
//		bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//		return string(bytes), err
//	}
//
// verify if password correct
func VerifyPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
