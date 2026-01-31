package files

import (
	"context"
	"files/internal/shared"

	"github.com/google/uuid"
)

func (s *FileService) checkFileAccess(
	ctx context.Context,
	fileID uuid.UUID,
	actorUserID uuid.UUID,
) (bool, error) {

	// 1 fetch file (existence + not deleted)
	isOwner, err := s.fileRepo.CheckFileOwnership(ctx, fileID, actorUserID)
	if err != nil {
		return false, err
	}

	// 2owner check
	if isOwner == true {
		return true, nil
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
	isOwner, err := s.fileRepo.CheckFileOwnership(ctx, fileID, userID)
	if err != nil {
		return false, err
	}
	if !isOwner {
		isEditor, err := s.shareRepo.IsFileEditableByUser(ctx, fileID, userID)
		if err != nil {
			return false, err
		}
		if !isEditor {
			return false, shared.ErrUnauthorized
		}
	}

	return true, nil
}
