package handler

import (
	"errors"
	"files/internal/handler/dto"
	"files/internal/service"
	"files/internal/shared"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

// POST /{fileID}/shortcut
func (h *Handler) CreateShortcut(w http.ResponseWriter, r *http.Request) {
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

	sc, err := h.registry.Shortcuts.CreateShortcut(r.Context(), &service.CreateShortcutInput{
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

	resp := dto.CreateShortcutResponse{
		ShortcutID: strconv.FormatInt(sc.ID, 10),
		FileID:     fileID.String(),
		CreatedAt:  sc.CreatedAt,
	}

	respondJSON(w, http.StatusCreated, resp)
}

// DELETE /shortcuts/{shortcutID}
func (h *Handler) DeleteShortcut(w http.ResponseWriter, r *http.Request) {
	actorID, err := h.getActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	shortcutID, err := uuid.Parse(r.PathValue("shortcutID"))
	if err != nil {
		http.Error(w, "invalid shortcut id", http.StatusBadRequest)
		return
	}

	err = h.registry.Shortcuts.DeleteShortcut(r.Context(), &service.DeleteShortcutInput{
		ShortcutID:  shortcutID,
		ActorUserID: actorID,
	})
	if err != nil {
		switch {
		case errors.Is(err, shared.ErrUnauthorized):
			http.Error(w, err.Error(), http.StatusForbidden)
		case errors.Is(err, shared.ErrRowNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
