package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// the router is also just an http handler with some middlewares
// interface for swapping routers
type Router interface {
	Route(path string)
}

type ChiRouter struct {
	router chi.Router
}

func NewChiRouter(proxies *Proxies) http.Handler {
	r := chi.NewRouter()

	// Public routes (no JWT)
	r.Route("/auth", func(r chi.Router) {
		r.Mount("/auth", proxies.Auth)

	})

	// Protected routes (JWT applied at gateway level or here)
	r.Route("/api", func(r chi.Router) {
		r.Mount("/upload", proxies.Upload)
		r.Mount("/meta", proxies.Metadata)
		r.Mount("/share", proxies.Sharing)
	})

	return r
}

func (r ChiRouter) Route(path string) {
	r.router.Route(path, func(r chi.Router) {})
}
