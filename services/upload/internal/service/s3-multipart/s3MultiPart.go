package multipart

import (
	"context"
	"upload/internal/model"
	"upload/internal/service/inputs"

	"github.com/google/uuid"
)

// type ServiceMethods interface {
// 	UploadSession(ctx context.Context, inputs *UploadSessionInput) (*model.UploadSession, error)
// 	UploadChunk(ctx context.Context, inputs *UploadChunkInput) (*model.UploadChunk, error)
// 	GetUploadStatus(ctx context.Context, upload_uuid uuid.UUID) (*model.UploadSession, error)
// }

func (s *UploadService) UploadSession(ctx context.Context, in *inputs.UploadSessionInput) (*model.UploadSession, error) {
	// form model
	var session model.UploadSession
	n := calculateTotalParts(in.FileSizeBytes, in.PartSize)
	session.UploadUUID = uuid.New()
	session.FileName = in.FileName
	session.FileSize = in.FileSizeBytes
	session.UserID = in.UserID
	session.TotalParts = n
	// insert into database
	_, err := s.registry.Sessions.CreateSession(ctx, &session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}
