package service

import (
	"errors"
	"upload/internal/model"

	"github.com/google/uuid"
)

// Get upload status handler
func (s *UploadService) GetUploadStatus(upload_uuid uuid.UUID, user_id uuid.UUID) (*model.UploadSession, error) {

	session, err := s.registry.Sessions.GetSessionByUUID(upload_uuid)
	if err != nil {
		return nil, err
	}

	if session.UserID != user_id {
		return nil, errors.New("Unauthorized")
	}

	return session, nil
}
