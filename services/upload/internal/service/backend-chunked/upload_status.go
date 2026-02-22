package backend

import (
	"context"
	"errors"
	"upload/internal/model"

	"github.com/google/uuid"
)

// Get upload status handler
func (s *ProxyUploadService) GetUploadStatus(ctx context.Context, upload_uuid uuid.UUID, user_id uuid.UUID) (*model.UploadSession, error) {

	session, err := s.registry.Sessions.GetSessionByUUID(ctx, upload_uuid)
	if err != nil {
		return nil, err
	}

	if session.UserID != user_id {
		return nil, errors.New("Unauthorized")
	}

	return session, nil
}
