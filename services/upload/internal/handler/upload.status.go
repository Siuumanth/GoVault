package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) GetUploadStatus(w http.ResponseWriter, r *http.Request) {
	userID, err := userIDFromHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	uploadIDStr := r.URL.Query().Get("upload_uuid")
	if uploadIDStr == "" {
		http.Error(w, "missing upload_uuid", http.StatusBadRequest)
		return
	}

	uploadUUID, err := uuid.Parse(uploadIDStr)
	if err != nil {
		http.Error(w, "invalid upload_uuid", http.StatusBadRequest)
		return
	}

	session, err := h.uploadService.GetUploadStatus(uploadUUID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res := UploadStatusResponse{
		UploadUUID:  session.UploadUUID.String(),
		Status:      string(session.Status),
		TotalChunks: session.TotalChunks,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
