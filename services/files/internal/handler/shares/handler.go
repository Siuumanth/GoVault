package shares

import "files/internal/service/shares"

type Handler struct {
	shares shares.ShareService
}

func New(shares shares.ShareService) *Handler {
	return &Handler{shares: shares}
}
