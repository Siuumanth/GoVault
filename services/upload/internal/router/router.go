package router

import (
	"log"
	"net/http"
	"upload/internal/handler"
	"upload/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterUploadRoutes(r chi.Router, h *handler.Handler) {
	r.Use(middleware.Logger)
	r.Use(middleware.ValidateUserID)

	// --- Public / Utility ---
	r.Get("/health", handler.HealthHandler)
	r.Get("/status", h.GetUploadStatus)

	// --- Proxy Flow (Backend-Chunked) ---
	r.Route("/proxy", func(r chi.Router) {
		r.Post("/session", h.CreateProxyUploadSession)
		r.Post("/chunk", h.UploadChunk)
	})

	// --- S3 Multipart Flow ---
	r.Route("/multipart", func(r chi.Router) {
		r.Post("/session", h.CreateMultipartSession)
		r.Post("/part", h.AddS3Part)
		r.Post("/complete", h.CompleteMultipart)
	})
}

/*
POST  /upload/proxy/session              Start session (uses backend 5MB chunk constant)
POST  /upload/proxy/chunk              Upload raw bytes (uses idx query param)

POST  /upload/multipart/session              Start session (frontend defines part_size_bytes)
POST  /upload/multipart/part              Record ETag after a successful direct S3 upload
POST  /upload/multipart/complete              Trigger backend to finalize S3 assembly

*/

func AfterStripLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)

		log.Printf("[AFTER] path=%s", r.URL.Path)
	})
}
