package service

import (
	"upload/internal/model"

	"github.com/google/uuid"
)

// Get upload status handler
func (s *UploadService) GetUploadStatus(upload_uuid uuid.UUID) (*model.UploadSession, error) {
	return s.registry.Sessions.GetSessionByUUID(upload_uuid)
}
