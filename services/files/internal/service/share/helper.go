package share

import (
	"context"

	"github.com/google/uuid"
)

func (s *ShareService) isUserOwnerOfFile(ctx context.Context, fileID *uuid.UUID, userID *uuid.UUID) (bool, error) {
	// fetch file summary and check
	res, err := s.fileRepo.CheckFileOwnership(ctx, *fileID, *userID)

	if err != nil {
		return false, err
	}
	return res, err

}
