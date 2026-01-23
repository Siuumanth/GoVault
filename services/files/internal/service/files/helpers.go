package files

import (
	"context"
	"files/internal/model"
	"files/internal/shared"

	"github.com/google/uuid"
)

func (s *FilesService) checkFileAccess(
	ctx context.Context,
	fileID uuid.UUID,
	actorUserID uuid.UUID,
) (*model.File, error) {

	// 1 fetch file (existence + not deleted)
	file, err := s.fileRepo.FetchFullFileByID(ctx, fileID)
	if err != nil {
		return nil, err
	}

	// 2owner check
	if file.UserID == actorUserID {
		return file, nil
	}

	// 3public access
	isPublic, err := s.shareRepo.IsFilePublic(ctx, fileID)
	if err != nil {
		return nil, err
	}
	if isPublic {
		return file, nil
	}

	// 4shared access
	isShared, err := s.shareRepo.IsFileSharedWithUser(ctx, fileID, actorUserID)
	if err != nil {
		return nil, err
	}
	if isShared {
		return file, nil
	}

	return nil, shared.ErrUnauthorized
}
