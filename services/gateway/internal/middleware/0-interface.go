package middleware

import "net/http"

type Middleware interface {
	Handle(http.Handler) http.Handler
}

// all middlewares will implement this interface
