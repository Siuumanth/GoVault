package handler

import (
	"errors"
	"files/internal/service"
	"files/internal/shared"
	"net/http"

	"github.com/google/uuid"
)

// POST /{fileID}/public
func (h *Handler) AddPublicAccess(w http.ResponseWriter, r *http.Request) {
	actorID, err := h.getActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid file id", http.StatusBadRequest)
		return
	}

	err = h.registry.Sharing.AddPublicAccess(r.Context(), &service.AddPublicAccessInput{
		FileID:      fileID,
		ActorUserID: actorID,
	})
	if err != nil {
		switch {
		case errors.Is(err, shared.ErrUnauthorized):
			http.Error(w, err.Error(), http.StatusForbidden)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DELETE /{fileID}/public
func (h *Handler) RemovePublicAccess(w http.ResponseWriter, r *http.Request) {
	actorID, err := h.getActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid file id", http.StatusBadRequest)
		return
	}

	err = h.registry.Sharing.RemovePublicAccess(r.Context(), &service.RemovePublicAccessInput{
		FileID:      fileID,
		ActorUserID: actorID,
	})
	if err != nil {
		switch {
		case errors.Is(err, shared.ErrUnauthorized):
			http.Error(w, err.Error(), http.StatusForbidden)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
