package shortcuts

import (
	"net/http"
	"strconv"

	"files/internal/handler/common"
	"files/internal/handler/dto"
	"files/internal/service"

	"github.com/google/uuid"
)

// POST /{fileID}/shortcut
func (h *Handler) CreateShortcut(w http.ResponseWriter, r *http.Request) {
	actorID, err := common.GetActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid file id", http.StatusBadRequest)
		return
	}

	sc, err := h.shortcuts.CreateShortcut(r.Context(), &service.CreateShortcutInput{
		FileID:      fileID,
		ActorUserID: actorID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	common.RespondJSON(w, http.StatusCreated, dto.CreateShortcutResponse{
		ShortcutID: strconv.FormatInt(sc.ID, 10),
		FileID:     fileID.String(),
		CreatedAt:  sc.CreatedAt,
	})
}

// DELETE /{fileID}/shortcut
func (h *Handler) DeleteShortcut(w http.ResponseWriter, r *http.Request) {
	actorID, err := common.GetActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	shortcutID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid shortcut id", http.StatusBadRequest)
		return
	}

	err = h.shortcuts.DeleteShortcut(r.Context(), &service.DeleteShortcutInput{
		FileID:      shortcutID,
		ActorUserID: actorID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
