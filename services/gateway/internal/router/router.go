package router

import (
	"net/http"

	MW "gateway/internal/middleware"

	"github.com/go-chi/chi/v5"
)

// TODO: Refactor router to allow for different implementatoins of router

// the router is also just an http handler with some middlewares
// interface for swapping routers
type Router interface {
	Route(path string)
}

type ChiRouter struct {
	router chi.Router
}

func NewChiRouter(proxies *Proxies, authz MW.Middleware) http.Handler {
	r := chi.NewRouter()

	// Public routes (no JWT)
	r.Route("/auth", func(r chi.Router) {
		r.Mount("/auth", proxies.Auth)
		r.Mount("/preview", proxies.Preview)
	})

	// Protected routes (JWT applied at gateway level or here)
	r.Route("/api", func(r chi.Router) {
		r.Use(authz.Handle) // middleware to authorize secure routes
		r.Mount("/upload", proxies.Upload)
		r.Mount("/meta", proxies.Metadata)
		r.Mount("/share", proxies.Sharing)
	})

	return r
}

func (r ChiRouter) Route(path string) {
	r.router.Route(path, func(r chi.Router) {})
}
