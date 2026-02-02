package handler

import (
	"fmt"
	"net/http"

	"download/internal/service"

	"github.com/google/uuid"
)

type Handler struct {
	downloadService *service.DownloadService
}

func NewDownloadHandler(s *service.DownloadService) *Handler {
	return &Handler{downloadService: s}
}

func userIDFromHeader(r *http.Request) (uuid.UUID, error) {
	id := r.Header.Get("X-User-ID")
	if id == "" {
		return uuid.Nil, fmt.Errorf("missing X-User-ID")
	}
	return uuid.Parse(id)
}
