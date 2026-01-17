package router

import (
	"upload/internal/handler"

	"github.com/go-chi/chi/v5"
)

func RegisterUploadRoutes(r chi.Router, h *handler.Handler) {
	r.Route("/", func(r chi.Router) {
		r.Post("/session", h.CreateUploadSession)
		r.Post("/chunk", h.UploadChunk)
		r.Get("/status", h.GetUploadStatus)
	})
}
