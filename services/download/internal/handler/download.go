package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) GetDownloadURL(w http.ResponseWriter, r *http.Request) {
	userID, err := userIDFromHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fileIDStr := r.URL.Query().Get("file_uuid")
	if fileIDStr == "" {
		http.Error(w, "missing file_uuid", http.StatusBadRequest)
		return
	}

	fileID, err := uuid.Parse(fileIDStr)
	if err != nil {
		http.Error(w, "invalid file_uuid", http.StatusBadRequest)
		return
	}

	resp, err := h.downloadService.GetDownloadURL(
		r.Context(),
		userID,
		fileID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
