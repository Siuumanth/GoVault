package router

import (
	"net/http"

	filesHandler "files/internal/handler/files"
	"files/internal/handler/health"
	sharesHandler "files/internal/handler/shares"
	shortcutsHandler "files/internal/handler/shortcuts"
	"files/internal/middleware"

	"github.com/go-chi/chi/v5"
)

// NewChiRouter wires all routes for the files service
func NewConfiguredChiRouter(
	filesH *filesHandler.Handler,
	sharesH *sharesHandler.Handler,
	shortcutsH *shortcutsHandler.Handler,
	healthH *health.Handler,
) http.Handler {

	r := chi.NewRouter()

	// inject actor id from x-user-id (non blocking)
	r.Use(middleware.ActorContext)

	// health
	r.Get("/health", healthH.HealthHandler)

	// lists
	r.Get("/owned", filesH.ListOwnedFiles)
	r.Get("/shared", filesH.ListSharedFiles)

	// file scoped routes
	r.Route("/{fileID}", func(r chi.Router) {
		// files
		r.Get("/", filesH.GetSingleFileSummary)
		r.Patch("/", filesH.UpdateFileName)
		r.Delete("/", filesH.SoftDeleteFile)
		r.Post("/copy", filesH.MakeFileCopy)

		// shares
		r.Route("/shares", func(r chi.Router) {
			r.Post("/", sharesH.AddFileShares)
			r.Get("/", sharesH.ListFileShares)
			r.Route("/{userID}", func(r chi.Router) {
				r.Patch("/", sharesH.UpdateFileShare)
				r.Delete("/", sharesH.RemoveFileShare)
			})
		})

		// public access (handled by shares)
		r.Post("/public", sharesH.AddPublicAccess)
		r.Delete("/public", sharesH.RemovePublicAccess)

		// shortcuts
		r.Post("/shortcut", shortcutsH.CreateShortcut)
		r.Delete("/shortcut", shortcutsH.DeleteShortcut)
	})

	return r
}
