package shares

import "files/internal/service"

type Handler struct {
	shares service.SharesService
}

func New(shares service.SharesService) *Handler {
	return &Handler{shares: shares}
}
