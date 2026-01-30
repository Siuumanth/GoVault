package router

import (
	"download/internal/handler"

	"github.com/go-chi/chi/v5"
)

func RegisterDownloadRoutes(r chi.Router, h *handler.Handler) {
	r.Route("/", func(r chi.Router) {
		r.Get("/download/url", h.GetDownloadURL)
	})
}
