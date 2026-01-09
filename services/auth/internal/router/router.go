package router

import (
	"auth/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h handler.UserHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", h.SignupHandler)
		r.Post("/login", h.LoginHandler)
		r.Get("/healtheck", h.HealthCheckHandler)
	})
	return r
}
