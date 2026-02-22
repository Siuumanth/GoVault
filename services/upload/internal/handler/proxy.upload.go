package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"upload/internal/service/inputs"
	"upload/shared"

	"github.com/google/uuid"
)

// methods to be implemented
func (h *Handler) CreateProxyUploadSession(w http.ResponseWriter, r *http.Request) {
	userID, err := userIDFromHeader(r)
	if err != nil {
		fmt.Println("UserID error: ", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var req CreateProxyUploadSessionRequest

	if err := decodeJSON(r, &req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	session, err := h.proxyUploadService.UploadSession(
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
		UploadUUID: session.UploadUUID,
		TotalParts: session.TotalParts,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

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
	err = h.proxyUploadService.UploadChunk(
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

	session, err := h.proxyUploadService.GetUploadStatus(
		r.Context(),
		uploadUUID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res := UploadStatusResponse{
		UploadUUID: session.UploadUUID.String(),
		Status:     string(session.Status),
		TotalParts: session.TotalParts,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
