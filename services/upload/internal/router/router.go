package router

import (
	"log"
	"net/http"
	"upload/internal/handler"
	"upload/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterUploadRoutes(r chi.Router, h *handler.Handler) {
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.Logger)
		//r.Use(AfterStripLogger)
		r.Get("/health", handler.HealthHandler)
		r.Post("/session", h.CreateUploadSession)
		r.Post("/chunk", h.UploadChunk)
		r.Get("/status", h.GetUploadStatus)
	})
}

func AfterStripLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)

		log.Printf("[AFTER] path=%s", r.URL.Path)
	})
}
