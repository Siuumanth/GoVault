package shares

import (
	"context"
	"files/internal/clients"
	"files/internal/model"
	"files/internal/repository"
	"files/internal/service/inputs"
	"files/internal/shared"
	"log"

	"github.com/google/uuid"
)

/*
Repo functions available:
type ShareRepository interface {
	CreateFileShare(ctx context.Context, p *model.FileShareParams) (*model.FileShare, error)
	FetchUserFileShare(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) (*model.FileShare, error)
	DeleteFileShare(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) error
	UpdateFileShare(ctx context.Context, p *model.FileShareParams) error
	FetchAllFileShares(ctx context.Context, fileID uuid.UUID) ([]*model.FileShare, error)
	IsFileSharedWithUser(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) (bool, error)

	// Public Access Methods
	CreatePublicAccess(ctx context.Context, fileID uuid.UUID) error
	DeletePublicAccess(ctx context.Context, fileID uuid.UUID) error
	IsFilePublic(ctx context.Context, fileID uuid.UUID) (bool, error)
}

Share service interface :
type ShareService interface {
	AddFileShares(ctx context.Context, in *AddFileSharesInput) error
	UpdateFileShare(ctx context.Context, in *UpdateFileShareInput) error
	RemoveFileShare(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID, recipientUserID uuid.UUID) error
	ListFileShares(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) ([]*model.FileShare, error)
	AddPublicAccess(ctx context.Context, in *AddPublicAccessInput) error
	RemovePublicAccess(ctx context.Context, in *RemovePublicAccessInput) error
}

*/

type ShareService struct {
	shareRepo  repository.SharesRepository
	fileRepo   repository.FilesRepository
	authClient *clients.AuthClient
}

func NewShareService(shareRepo repository.SharesRepository, fileRepo repository.FilesRepository, authClient *clients.AuthClient) *ShareService {
	return &ShareService{
		shareRepo:  shareRepo,
		fileRepo:   fileRepo,
		authClient: authClient,
	}
}

func (s *ShareService) AddFileShares(ctx context.Context, in *inputs.AddFileSharesInput) error {
	// owner check
	if err := s.fileRepo.CheckFileOwnership(ctx, in.FileID, in.ActorUserID); err != nil {
		return err
	}

	// share count check
	if len(in.Recipients) > shared.MAX_SHARES {
		return shared.ErrTooManyShares
	}

	// collect emails
	emails := make([]string, 0, len(in.Recipients))
	for _, r := range in.Recipients {
		emails = append(emails, r.Email)
	}
	// Done: user Auth microservice for this
	// bulk resolve emails to userIDs to insert shares
	log.Println("Resolving emails to userID")
	emailToUserID, err := s.authClient.ResolveEmails(ctx, emails)
	if err != nil {
		return err
	}
	log.Println("Emails resolved to userID ", emailToUserID)

	// TODO: make tis a transaction
	// create shares
	for _, r := range in.Recipients {
		userID, ok := emailToUserID[r.Email]
		if !ok {
			return shared.ErrRowNotFound // or ErrUserNotFound
		}

		p := &model.FileShareParams{
			FileID:           in.FileID,
			SharedWithUserID: userID,
			Permission:       r.Permission,
		}

		if _, err := s.shareRepo.CreateFileShare(ctx, p); err != nil {
			return err
		}
	}

	return nil
}

func (s *ShareService) UpdateFileShare(ctx context.Context, in *inputs.UpdateFileShareInput) error {
	// validate if owner, only owner can update
	if err := s.fileRepo.CheckFileOwnership(ctx, in.FileID, in.ActorUserID); err != nil {
		return err
	}

	// update share
	p := &model.FileShareParams{
		FileID:           in.FileID,
		SharedWithUserID: in.RecipientUserID,
		Permission:       in.Permission,
	}

	return s.shareRepo.UpdateFileShare(ctx, p)
}

func (s *ShareService) RemoveFileShare(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID, recipientUserID uuid.UUID) error {
	// validate if owner
	if err := s.fileRepo.CheckFileOwnership(ctx, fileID, actorUserID); err != nil {
		return err
	}

	// delete share
	return s.shareRepo.DeleteFileShare(ctx, fileID, recipientUserID)
}

func (s *ShareService) ListFileShares(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) ([]*model.FileShare, error) {
	// validate if owner, only owner can list shares
	if err := s.fileRepo.CheckFileOwnership(ctx, fileID, actorUserID); err != nil {
		return nil, err
	}
	return s.shareRepo.FetchAllFileShares(ctx, fileID)
}
