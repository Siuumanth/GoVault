package shortcuts

import (
	"context"
	"files/internal/shared"

	"github.com/google/uuid"
)

// TODO: MAke a commmon check file access helper for files and shortcuts services n make logic more error handling

func (s *ShortcutService) checkFileAccess(
	ctx context.Context,
	fileID uuid.UUID,
	actorUserID uuid.UUID,
) (bool, error) {

	// 1,2 owner check (existence + not deleted)
	err := s.filesRepo.CheckFileOwnership(ctx, fileID, actorUserID)
	if err != nil {
		return false, err
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
