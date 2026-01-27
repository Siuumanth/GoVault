package router

import (
	"net/http"

	"files/internal/handler"
	"files/internal/middleware"

	"github.com/go-chi/chi/v5"
)

// mounts all file-service routes
func NewFileRouter(h *handler.Handler) http.Handler {
	r := chi.NewRouter()

	// attach actor middleware once
	r.Use(middleware.ActorContext)

	// --------------------
	// FILES (mixed access)
	// --------------------
	r.Route("/{fileID}", func(r chi.Router) {
		r.Get("/", h.GetSingleFileSummary)
		r.Patch("/", h.UpdateFileName)
		r.Delete("/", h.SoftDeleteFile)
		r.Post("/copy", h.MakeFileCopy)

		// shares
		r.Route("/shares", func(r chi.Router) {
			r.Post("/", h.AddFileShares)
			r.Get("/", h.ListFileShares)

			r.Route("/{userID}", func(r chi.Router) {
				r.Patch("/", h.UpdateFileShare)
				r.Delete("/", h.UpdateFileShare) // if you later add RemoveFileShare handler, swap this
			})
		})

		// public access
		r.Post("/public", h.AddPublicAccess)
		r.Delete("/public", h.RemovePublicAccess)

		// shortcuts
		r.Post("/shortcut", h.CreateShortcut)
	})

	// --------------------
	// LISTS
	// --------------------
	r.Get("/moved", h.ListOwnedFiles)
	r.Get("/shared", h.ListSharedFiles)

	// --------------------
	// SHORTCUTS
	// --------------------
	r.Delete("/shortcuts/{shortcutID}", h.DeleteShortcut)

	return r
}
