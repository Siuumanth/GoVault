package handler

import (
	"net/http"
	"upload/internal/service/inputs"
	"upload/shared"

	"github.com/google/uuid"
)

// POST /upload/multipart/session
// services/upload/internal/handler/multipart.upload.go

// POST /upload/multipart/session
func (h *Handler) CreateMultipartSession(w http.ResponseWriter, r *http.Request) {
	userID, err := userIDFromHeader(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req CreateMultipartSessionRequest
	if err := decodeJSON(r, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Call service - catching the session AND the newly generated parts
	session, serviceParts, err := h.multipartUploadService.UploadSession(r.Context(), &inputs.UploadSessionInput{
		UserID:        userID,
		FileName:      req.FileName,
		FileSizeBytes: req.FileSizeBytes,
		PartSize:      req.PartSizeBytes,
	})
	if err != nil {
		// Log the error for debugging before sending to client
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Map service-layer parts to handler-layer DTOs
	var responseParts []PresignedPart
	for _, p := range serviceParts {
		responseParts = append(responseParts, PresignedPart{
			PartNumber: p.PartNumber,
			URL:        p.URL,
		})
	}

	// Return everything in one JSON response
	respondJSON(w, http.StatusOK, CreateMultipartUploadSessionResponse{
		UploadUUID:      session.UploadUUID,
		TotalParts:      session.TotalParts,
		StorageUploadID: session.StorageUploadID,
		Parts:           responseParts, // Frontend now has all URLs immediately!
	})
}

// 2. Add S3 Part Handler (Records the ETag from frontend)
// POST /multipart/part
func (h *Handler) AddS3Part(w http.ResponseWriter, r *http.Request) {
	var req AddS3PartRequest
	if err := decodeJSON(r, &req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	// extract user id form context and pass it to the service
	actorID := r.Context().Value(shared.ActorIDKey).(uuid.UUID)

	err := h.multipartUploadService.AddS3Part(r.Context(), req.UploadUUID, actorID, &inputs.AddPartInput{
		PartNumber: req.PartNumber,
		SizeBytes:  req.SizeBytes,
		Etag:       req.Etag,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// 3. Complete Multipart Handler (Triggers S3 assembly)
// POST /multipart/complete
func (h *Handler) CompleteMultipart(w http.ResponseWriter, r *http.Request) {
	var req CompleteMultipartRequest
	if err := decodeJSON(r, &req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	actorID := r.Context().Value(shared.ActorIDKey).(uuid.UUID)

	err := h.multipartUploadService.CompleteS3Multipart(r.Context(), req.UploadUUID, actorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

/*

// GET /upload/multipart/presign?uploadUUID=...
func (h *Handler) GenerateMultipartPartURLs(w http.ResponseWriter, r *http.Request) {

	uploadUUIDStr := r.URL.Query().Get("uploadUUID")
	if uploadUUIDStr == "" {
		http.Error(w, "missing uploadUUID", http.StatusBadRequest)
		return
	}

	uploadUUID, err := uuid.Parse(uploadUUIDStr)
	if err != nil {
		http.Error(w, "invalid uploadUUID", http.StatusBadRequest)
		return
	}

	serviceParts, err := h.multipartUploadService.GenerateAllPartURLs(
		r.Context(),
		uploadUUID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Map service → handler DTO
	var responseParts []PresignedPart
	for _, p := range serviceParts {
		responseParts = append(responseParts, PresignedPart{
			PartNumber: p.PartNumber,
			URL:        p.URL,
		})
	}

	respondJSON(w, http.StatusOK, GeneratePartURLsResponse{
		UploadUUID: uploadUUID,
		Parts:      responseParts,
	})
}

*/
