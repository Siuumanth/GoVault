package shortcuts

import (
	"context"
	"files/internal/model"
	"files/internal/repository"
	"files/internal/service/inputs"
	"files/internal/shared"
)

/*
type ShortcutRepository interface {
	CreateShortcut(ctx context.Context, fileID, userID uuid.UUID) (*model.FileShortcut, error)
	DeleteShortcut(ctx context.Context, shortcutID, userID uuid.UUID) error
}
*/

type ShortcutService struct {
	filesRepo     repository.FilesRepository
	sharesRepo    repository.SharesRepository
	shortcutsRepo repository.ShortcutsRepository
}

func NewShortcutService(
	filesRepo repository.FilesRepository,
	sharesRepo repository.SharesRepository,
	shortcutsRepo repository.ShortcutsRepository,
) *ShortcutService {
	return &ShortcutService{
		filesRepo:     filesRepo,
		sharesRepo:    sharesRepo,
		shortcutsRepo: shortcutsRepo,
	}
}

func (s *ShortcutService) CreateShortcut(ctx context.Context, in *inputs.CreateShortcutInput) (*model.FileShortcut, error) {

	// verify access (owner OR public OR shared)
	isAllowed, err := s.checkFileAccess(ctx, in.FileID, in.ActorUserID)
	if err != nil {
		return nil, err
	} else if !isAllowed {
		return nil, shared.ErrUnauthorized
	}

	return s.shortcutsRepo.CreateShortcut(
		ctx,
		in.FileID,
		in.ActorUserID,
	)
}

func (s *ShortcutService) DeleteShortcut(
	ctx context.Context,
	in *inputs.DeleteShortcutInput,
) error {

	// only shortcut owner can delete
	return s.shortcutsRepo.DeleteShortcut(
		ctx,
		in.FileID,
		in.ActorUserID,
	)
}
