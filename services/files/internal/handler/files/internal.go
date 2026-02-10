package files

import (
	"files/internal/handler/common"
	"files/internal/handler/dto"
	"files/internal/model"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// POST /internal/file
func (h *Handler) RegisterFile(w http.ResponseWriter, r *http.Request) {
	// 1. Decode the request body (DTO)
	var req dto.CreateFileRequest
	if err := common.DecodeJSON(r, &req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// 2. Validate mandatory fields (Basic integrity check)
	if req.FileUUID == uuid.Nil || req.UserID == uuid.Nil || req.StorageKey == "" {
		http.Error(w, "missing required file metadata", http.StatusBadRequest)
		return
	}

	// 3. Map DTO to Service Model (CreateFileParams)
	// Using the pointer logic similar to your UpdateFileName pattern
	_, err := h.files.CreateFile(r.Context(), &model.CreateFileParams{
		FileUUID:   req.FileUUID,
		UserID:     req.UserID,
		Name:       req.Name,
		SizeBytes:  req.SizeBytes,
		MimeType:   req.MimeType,
		Checksum:   req.CheckSum,
		StorageKey: req.StorageKey,
	})

	if err != nil {
		// Handle database or unique constraint errors
		log.Printf("[ERROR] Failed to register file: %v", err)
		http.Error(w, "failed to save file metadata", http.StatusInternalServerError)
		return
	}

	// 4. Respond with 201 Created
	w.WriteHeader(http.StatusCreated)
}
