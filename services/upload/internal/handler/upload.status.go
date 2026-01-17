package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) GetUploadStatus(w http.ResponseWriter, r *http.Request) {
	_, err := userIDFromHeader(r)
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

	session, err := h.uploadService.GetUploadStatus(uploadUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}
