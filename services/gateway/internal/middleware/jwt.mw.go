package middleware

import (
	"context"
	"errors"
	"fmt"
	"gateway/internal/utils"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct{}

const authCtxKey string = "auth"

// http.Handler is an interface with 1 function - serveHTTP(w,r)
// w http.ResponseWriter is an interface so its automatically pass by ref, so * not needed
func (j JWT) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Extract the "Bearer" token from the request
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			return // no auth header
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return // invalid header format
		}
		tokenString := strings.TrimPrefix(authHeader, prefix)

		// now we decode jwt to check validity
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Parse the JWT token with the validation function
		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			// Validate the signing method is HMAC-SHA256
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Return the JWT secret for signature verification
			return []byte(jwtSecret), nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				http.Error(w, "Token Expired", http.StatusUnauthorized)
				return
			} else if errors.Is(err, jwt.ErrTokenMalformed) {
				http.Error(w, "Token Malformed", http.StatusUnauthorized)
			}
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Validate the parsed JWT token
		if parsedToken.Valid {
			log.Println("Valid JWT")
		} else {
			http.Error(w, "Invalid Login Token", http.StatusUnauthorized)
			log.Println("Invalid JWT:", tokenString)
			return
		}

		// Extract the claims of the parsed JWT token
		claims, ok := parsedToken.Claims.(jwt.MapClaims) // type assertion
		if !ok {
			http.Error(w, "Invalid Login Token", http.StatusUnauthorized)
			//log.Println("Invalid Login Token:", authHeader)
			return
		}

		authCtx, err := utils.MapClaims(claims)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, authCtxKey, authCtx)

		// Call the next handler with the enriched context
		// attaching old request with new contexrt and passing request
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
