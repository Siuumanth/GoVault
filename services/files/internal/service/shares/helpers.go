package shares

import (
	"context"
	"files/internal/shared"

	"github.com/google/uuid"
)

// MAybe 1 helper ws enuf
func (s *ShareService) isUserOwnerOfFile(ctx context.Context, fileID *uuid.UUID, userID *uuid.UUID) (bool, error) {
	// fetch file summary and check
	res, err := s.fileRepo.CheckFileOwnership(ctx, *fileID, *userID)

	if err != nil {
		return false, err
	}
	return res, err

}

func (s *ShareService) assertOwner(
	ctx context.Context,
	fileID *uuid.UUID,
	userID *uuid.UUID,
) error {
	ok, err := s.isUserOwnerOfFile(ctx, fileID, userID)
	if err != nil {
		return err
	}
	if !ok {
		return shared.ErrUnauthorized
	}
	return nil
}
