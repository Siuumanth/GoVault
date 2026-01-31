package files

import (
	"files/internal/handler/common"
	"files/internal/service/files"
	"net/http"

	"github.com/google/uuid"
)

type InternalHandler struct {
	service *files.FileService
}

func NewInternalHandler(s *files.FileService) *InternalHandler {
	return &InternalHandler{service: s}
}

// GET files/{fileID}/download
func (h *InternalHandler) GetDownload(w http.ResponseWriter, r *http.Request) {
	userID, err := common.GetActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}

	file, err := h.service.GetDownloadDetails(r.Context(), fileID, userID)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	// Returns the full model (internal model.File)
	// This includes the storage_key which the Public API hides.
	common.RespondJSON(w, http.StatusOK, file)
}
