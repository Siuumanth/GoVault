package handler

// creating handler registry

import (
	"files/internal/handler/files"
	"files/internal/handler/health"
	"files/internal/handler/shares"
	"files/internal/handler/shortcuts"
	"files/internal/service"
)

type HandlerRegistry struct {
	Files     *files.Handler
	Shares    *shares.Handler
	Shortcuts *shortcuts.Handler
	Health    *health.Handler
}

func NewHandlerRegistry(s *service.ServiceRegistry) *HandlerRegistry {
	return &HandlerRegistry{
		Files:     files.NewHandler(s.Files),
		Shares:    shares.NewHandler(s.Shares),
		Shortcuts: shortcuts.NewHandler(s.Shortcuts),
		Health:    health.NewHandler(),
	}
}
