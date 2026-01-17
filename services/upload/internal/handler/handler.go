package handler

import (
	"fmt"
	"net/http"
	"upload/internal/service"

	"github.com/google/uuid"
)

type Handler struct {
	uploadService *service.UploadService
}

func NewUploadHandler(uploadService *service.UploadService) *Handler {
	return &Handler{uploadService: uploadService}
}

// helpers
func userIDFromHeader(r *http.Request) (uuid.UUID, error) {
	id := r.Header.Get("X-User-ID")
	if id == "" {
		return uuid.Nil, fmt.Errorf("missing X-User-ID")
	}
	return uuid.Parse(id)
}
