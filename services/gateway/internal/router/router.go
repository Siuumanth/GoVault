package router

import (
	"gateway/internal/health"
	MW "gateway/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// DONE: Refactor router to allow for different implementatoins of router

// the router is also just an http handler with some middlewares

type Router interface {
	http.Handler
	ConfigureRoutes(proxies *Proxies, authz MW.Middleware)
}

/*
.
.
.
.
MAKING CHI ROUTER:
.
.
.
.
*/

// container for holding a chi router
type ChiRouter struct {
	mux *chi.Mux // type of http handler
}

func NewChiRouter() *ChiRouter {
	return &ChiRouter{
		mux: chi.NewRouter(),
	}
}

// implement ServeHTTP to implement http.Handler
func (r *ChiRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r ChiRouter) ConfigureRoutes(proxies *Proxies, authz MW.Middleware) {
	// Health check
	r.mux.Get("/health", health.HealthCheckHandler)

	// Public routes (no JWT)
	r.mux.Route("/auth", func(r chi.Router) {
		r.Mount("/", http.StripPrefix("/auth", proxies.Auth))
	})

	r.mux.Route("/preview", func(r chi.Router) {
		r.Mount("/", http.StripPrefix("/preview", proxies.Preview))
	})

	// Protected routes (JWT applied at gateway level or here)
	r.mux.Route("/api", func(r chi.Router) {
		r.Use(authz.Handle) // middleware to authorize secure routes
		r.Mount("/upload", http.StripPrefix("/upload", proxies.Upload))
		r.Mount("/meta", http.StripPrefix("/meta", proxies.Upload))
		r.Mount("/share", http.StripPrefix("/upload", proxies.Upload))
	})
}
