package handler

import (
	"errors"
	"io"
	"net/http"
	"strconv"
	"upload/internal/service/inputs"
	"upload/shared"

	"github.com/google/uuid"
)

/*
Headers:
- Upload-UUID
- Checksum
- ChunkID  // in query parameter
- X-User-ID
*/

func (h *Handler) UploadChunk(w http.ResponseWriter, r *http.Request) {
	userID, err := userIDFromHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Read metadata (NOT from body)
	uploadUUIDStr := r.Header.Get("Upload-UUID")
	if uploadUUIDStr == "" {
		http.Error(w, "missing Upload-UUID header", http.StatusBadRequest)
		return
	}

	uploadUUID, err := uuid.Parse(uploadUUIDStr)
	if err != nil {
		http.Error(w, "missing Upload-UUID", http.StatusBadRequest)
		return
	}

	checksum := r.Header.Get("Checksum")

	// Chunk index from query
	chunkIDStr := r.URL.Query().Get("id")
	chunkID, err := strconv.Atoi(chunkIDStr)
	if err != nil || chunkID < 0 {
		http.Error(w, "invalid chunk id", http.StatusBadRequest)
		return
	}

	// limit chunk size
	const maxChunkSize = shared.ChunkSizeBytes
	r.Body = http.MaxBytesReader(w, r.Body, maxChunkSize)
	defer r.Body.Close()

	// Call service with RAW STREAM
	err = h.uploadService.UploadChunk(
		r.Context(),
		&inputs.UploadChunkInput{
			UserID:     userID,
			UploadUUID: uploadUUID,
			ChunkID:    chunkID,
			CheckSum:   checksum,
			ChunkBytes: r.Body,
		})
	// io.ReadCloser is an io.Reader that returns an error after n Read calls return an error.
	// This is useful for reading a stream that is known to be closed (e.g. HTTP response body).
	readCloser := io.NopCloser(r.Body)
	defer readCloser.Close()

	// idempotency handling
	if errors.Is(err, shared.ErrChunkAlreadyExists) {
		w.WriteHeader(http.StatusOK)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
