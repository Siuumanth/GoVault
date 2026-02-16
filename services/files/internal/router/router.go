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

	// identity extraction middleware
	r.Use(middleware.ActorContext)

	// ---------- private ----------
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.RequireActor) // 🔒 EVERYTHING inside requires auth

		// lists
		r.Get("/me/owned", filesH.ListOwnedFiles)
		r.Get("/me/shared", filesH.ListSharedFiles)
		r.Get("/me/shortcuts", shortcutsH.ListShortcuts)
	})

	// ---------- internal (Service-to-Service) ----------
	r.Route("/internal", func(r chi.Router) {
		// No ActorContext needed here usually, as this is machine-to-machine
		r.Post("/file", filesH.RegisterFile)
	})

	// ---------- public ----------
	r.Get("/health", healthH.HealthHandler)

	// ---------- files ----------
	// add at last or it will cause collision
	r.Route("/f/{fileID}", func(r chi.Router) {

		// public file access
		r.Get("/", filesH.GetSingleFileSummary)
		r.Get("/download", filesH.GetDownload)

		// ---------- private ----------
		r.Group(func(r chi.Router) {
			r.Use(middleware.RequireActor) // 🔒 EVERYTHING inside requires user

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
