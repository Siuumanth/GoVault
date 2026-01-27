package handler

import (
	"files/internal/service"
)

type Handler struct {
	registry *service.ServiceRegistry
}

func NewUploadHandler(regr *service.ServiceRegistry) *Handler {
	return &Handler{registry: regr}
}
