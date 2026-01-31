package internalhandler

import (
	"files/internal/handler/common"
	"net/http"

	"github.com/google/uuid"
)

type InternalHandler struct {
	service *internal.InternalFileService
}

func NewInternalHandler(s *internal.InternalFileService) *InternalHandler {
	return &InternalHandler{service: s}
}

// GET /internal/files/{fileID}/metadata
func (h *InternalHandler) GetInternalMetadata(w http.ResponseWriter, r *http.Request) {
	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}

	file, err := h.service.GetMetadata(r.Context(), fileID)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	// Returns the full model (internal model.File)
	// This includes the storage_key which the Public API hides.
	common.RespondJSON(w, http.StatusOK, file)
}
