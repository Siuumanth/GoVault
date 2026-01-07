package router

import (
	MW "gateway/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// DONE: Refactor router to allow for different implementatoins of router

// the router is also just an http handler with some middlewares
// interface for swapping routers

/*
dependency inject the router
*/
// any router than configure routes should implement this
type Router interface {
	http.Handler
	ConfigureRoutes(proxies *Proxies, authz MW.Middleware)
}

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
	// Public routes (no JWT)
	r.mux.Route("/auth", func(r chi.Router) {
		r.Mount("/auth", proxies.Auth)
		r.Mount("/preview", proxies.Preview)
	})

	// Protected routes (JWT applied at gateway level or here)
	r.mux.Route("/api", func(r chi.Router) {
		r.Use(authz.Handle) // middleware to authorize secure routes
		r.Mount("/upload", proxies.Upload)
		r.Mount("/meta", proxies.Metadata)
		r.Mount("/share", proxies.Sharing)
	})
}
