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

	// identity extraction (optional)
	r.Use(middleware.ActorContext)

	// ---------- internal (Service-to-Service) ----------
	r.Route("/internal", func(r chi.Router) {
		// No ActorContext needed here usually, as this is machine-to-machine
		r.Post("/file", filesH.RegisterFile) // This matches your existing AddFile logic
	})

	// ---------- public ----------
	r.Get("/health", healthH.HealthHandler)

	// public file access
	r.Route("/{fileID}", func(r chi.Router) {
		r.Get("/", filesH.GetSingleFileSummary)
	})
	r.Route("/{fileID}/download", func(r chi.Router) {
		r.Get("/", filesH.GetDownload)
	})

	// ---------- private ----------
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.RequireActor) // 🔒 EVERYTHING inside requires user

		// lists
		r.Get("/owned", filesH.ListOwnedFiles)
		r.Get("/shared", filesH.ListSharedFiles)
		r.Get("/shortcuts", shortcutsH.ListShortcuts)

		// file scoped private ops
		r.Route("/{fileID}", func(r chi.Router) {
			// files
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

			// public access toggles (owner-only)
			r.Post("/public", sharesH.AddPublicAccess)
			r.Delete("/public", sharesH.RemovePublicAccess)

			// shortcuts
			r.Post("/shortcut", shortcutsH.CreateShortcut)
			r.Delete("/shortcut", shortcutsH.DeleteShortcut)
		})
	})

	return r
}
