package files

import (
	"errors"
	"files/internal/handler/common"
	"files/internal/model"
	"files/internal/shared"
	"net/http"

	"github.com/google/uuid"
)

// GET files/{fileID}/download (public)
func (h *Handler) GetDownload(w http.ResponseWriter, r *http.Request) {
	userID := common.GetOptionalActorID(r)

	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid uuid", http.StatusBadRequest)
		return
	}

	downloadInfo, err := h.files.GetDownloadDetails(r.Context(), fileID, userID)
	if err != nil {
		if errors.Is(err, shared.ErrUnauthorized) {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	common.RespondJSON(w, http.StatusOK, model.DownloadResponse{
		DownloadURL: downloadInfo.DownloadURL,
		ExpiresAt:   downloadInfo.ExpiresAt,
		FileName:    downloadInfo.FileName,
	})
}
