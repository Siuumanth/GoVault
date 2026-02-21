package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"upload/internal/service/inputs"
)

// methods to be implemented
func (h *Handler) CreateUploadSession(w http.ResponseWriter, r *http.Request) {
	userID, err := userIDFromHeader(r)
	if err != nil {
		fmt.Println("UserID error: ", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req CreateUploadSessionRequest

	if err := decodeJSON(r, &req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	session, err := h.uploadService.UploadSession(
		r.Context(),
		&inputs.UploadSessionInput{
			UserID:        userID,
			FileName:      req.FileName,
			FileSizeBytes: req.FileSizeBytes,
		})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := CreateUploadSessionResponse{
		UploadUUID:  session.UploadUUID,
		TotalChunks: session.TotalChunks,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
