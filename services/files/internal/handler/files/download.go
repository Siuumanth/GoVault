package files

import (
	"files/internal/handler/common"
	"files/internal/handler/dto"
	"net/http"

	"github.com/google/uuid"
)

// GET files/{fileID}/download (public)
func (h *Handler) GetDownload(w http.ResponseWriter, r *http.Request) {
	userID := common.GetOptionalActorID(r)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusUnauthorized)
	// 	return
	// }

	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}

	downloadInfo, err := h.files.GetDownloadDetails(r.Context(), fileID, userID)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	// Returns the full model (internal model.File)
	// This includes the storage_key which the Public API hides.
	common.RespondJSON(w, http.StatusOK, dto.DownloadResponse{
		StorageKey: downloadInfo.StorageKey,
		ExpiresAt:  downloadInfo.ExpiresAt,
	})
}
