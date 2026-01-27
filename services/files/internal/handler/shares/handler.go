package shares

import "files/internal/service"

type Handler struct {
	shares service.SharingService
}

func New(shares service.SharingService) *Handler {
	return &Handler{shares: shares}
}
