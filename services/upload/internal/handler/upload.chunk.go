package handler

import (
	"errors"
	"net/http"
	"strconv"
	"upload/internal/service"
	"upload/shared"
)

func (h *Handler) UploadChunk(w http.ResponseWriter, r *http.Request) {
	userID, err := userIDFromHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	var req UploadChunkRequest

	if err := decodeJSON(r, req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// get id from query parameter
	chunkID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "invalid query parameter", http.StatusBadRequest)
		return
	}

	err = h.uploadService.UploadChunk(r.Context(), &service.UploadChunkInput{
		UploadUUID: req.UploadUUID,
		UserID:     userID,
		ChunkID:    chunkID,
		CheckSum:   req.CheckSum,
		ChunkBytes: req.ChunkBytes,
	})
	// check if chunk already exists
	if !errors.Is(err, shared.ErrChunkAlreadyExists) && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
