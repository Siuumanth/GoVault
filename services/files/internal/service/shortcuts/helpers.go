package shortcuts

import (
	"context"
	"files/internal/shared"

	"github.com/google/uuid"
)

// TODO: MAke a commmon check file access helper for files and shortcuts services

func (s *ShortcutService) checkFileAccess(
	ctx context.Context,
	fileID uuid.UUID,
	actorUserID uuid.UUID,
) (bool, error) {

	// 1 fetch file (existence + not deleted)
	isOwner, err := s.filesRepo.CheckFileOwnership(ctx, fileID, actorUserID)
	if err != nil {
		return false, err
	}

	// 2owner check
	if isOwner == true {
		return true, nil
	}

	// 3public access
	isPublic, err := s.sharesRepo.IsFilePublic(ctx, fileID)
	if err != nil {
		return false, err
	}
	if isPublic {
		return true, nil
	}

	// 4shared access
	isShared, err := s.sharesRepo.IsFileSharedWithUser(ctx, fileID, actorUserID)
	if err != nil {
		return false, err
	}
	if isShared {
		return true, nil
	}

	return false, shared.ErrUnauthorized
}
