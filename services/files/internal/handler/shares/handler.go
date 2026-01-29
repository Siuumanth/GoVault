package shares

import "files/internal/service/shares"

type Handler struct {
	shares *shares.ShareService
}

func NewHandler(shares *shares.ShareService) *Handler {
	return &Handler{shares: shares}
}
