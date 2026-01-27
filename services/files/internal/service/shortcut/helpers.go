package shortcut

// TODO: update names to plural
import (
	"context"
	"files/internal/model"
	"files/internal/shared"

	"github.com/google/uuid"
)

func (s *ShortcutsService) checkFileAccess(
	ctx context.Context,
	fileID uuid.UUID,
	actorUserID uuid.UUID,
) (*model.File, error) {

	// 1 fetch file (existence + not deleted)
	file, err := s.filesRepo.FetchFullFileByID(ctx, fileID)
	if err != nil {
		return nil, err
	}

	// 2owner check
	if file.UserID == actorUserID {
		return file, nil
	}

	// 3public access
	isPublic, err := s.sharesRepo.IsFilePublic(ctx, fileID)
	if err != nil {
		return nil, err
	}
	if isPublic {
		return file, nil
	}

	// 4shared access
	isShared, err := s.sharesRepo.IsFileSharedWithUser(ctx, fileID, actorUserID)
	if err != nil {
		return nil, err
	}
	if isShared {
		return file, nil
	}

	return nil, shared.ErrUnauthorized
}
