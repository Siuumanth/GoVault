package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"upload/internal/service/backend-chunked"
	multipart "upload/internal/service/s3-multipart"

	"github.com/google/uuid"
)

// tempoerory
type Handler struct {
	proxyUploadService     *backend.ProxyUploadService
	multipartUploadService *multipart.MultipartUploadService
}

func NewUploadHandler(proxyUploadService *backend.ProxyUploadService, multipartUploadService *multipart.MultipartUploadService) *Handler {
	return &Handler{
		proxyUploadService:     proxyUploadService,
		multipartUploadService: multipartUploadService,
	}
}

// helpers

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("Error encoding JSON: %v", err)
	}
}

func userIDFromHeader(r *http.Request) (uuid.UUID, error) {
	id := r.Header.Get("X-User-ID")
	if id == "" {
		return uuid.Nil, fmt.Errorf("missing X-User-ID")
	}
	return uuid.Parse(id)
}

func decodeJSON(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return err
	}

	// ensure only ONE JSON object
	if dec.More() {
		return errors.New("multiple JSON objects not allowed")
	}

	return nil
}
