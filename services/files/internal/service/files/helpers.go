package files

import (
	"context"
	"errors"
	"files/internal/shared"

	"github.com/google/uuid"
)

func (s *FileService) checkFileAccess(
	ctx context.Context,
	fileID uuid.UUID,
	actorUserID uuid.UUID,
) (bool, error) {

	// 1,2 fetch file (existence + not deleted)
	err := s.fileRepo.CheckFileOwnership(ctx, fileID, actorUserID)
	if err != nil {
		return false, err
	}

	// 3public access
	isPublic, err := s.shareRepo.IsFilePublic(ctx, fileID)
	if err != nil {
		return false, err
	}
	if isPublic {
		return true, nil
	}

	// 4shared access
	isShared, err := s.shareRepo.IsFileSharedWithUser(ctx, fileID, actorUserID)
	if err != nil {
		return false, err
	}
	if isShared {
		return true, nil
	}

	return false, shared.ErrUnauthorized
}

func (s *FileService) canUserEditFile(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) (bool, error) {
	err := s.fileRepo.CheckFileOwnership(ctx, fileID, userID)

	if errors.Is(err, shared.ErrUnauthorized) {
		isEditor, err := s.shareRepo.IsFileEditableByUser(ctx, fileID, userID)
		if err != nil {
			return false, err
		}
		if !isEditor {
			return false, shared.ErrUnauthorized
		}
	} else if err != nil {
		return false, err
	}

	return true, nil
}
