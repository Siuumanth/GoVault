package handler

import (
	"encoding/json"
	"net/http"
	"upload/internal/service"
)

func (h *Handler) UploadChunk(w http.ResponseWriter, r *http.Request) {
	userID, err := userIDFromHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req UploadChunkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	err = h.uploadService.UploadChunk(r.Context(), &service.UploadChunkInput{
		UploadUUID: req.UploadUUID,
		UserID:     userID,
		ChunkID:    req.ChunkID,
		CheckSum:   req.CheckSum,
		ChunkBytes: req.ChunkBytes,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
