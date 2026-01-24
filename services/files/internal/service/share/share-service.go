package share

import (
	"context"
	"files/internal/model"
	"files/internal/repository"
	"files/internal/service"
	"files/internal/shared"
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
	type FileShare struct {
	ID               int64
	FileID           int64
	SharedWithUserID uuid.UUID
	Permission       string
	CreatedAt        time.Time
}

type File struct {
	ID         int64
	FileUUID   uuid.UUID
	SessionID  *int64
	UserID     uuid.UUID
	FileName   string
	MimeType   string
	SizeBytes  int64
	StorageKey string
	Checksum   *string
	CreatedAt  time.Time
	DeletedAt  *time.Time
}

type ShareRecipientInput struct {
	Email       string
	Permissions string
}

type AddFileSharesInput struct {
	FileID      uuid.UUID
	ActorUserID uuid.UUID
	Recipients  []ShareRecipientInput
}
s

*/

type ShareService struct {
	shareRepo repository.ShareRepository
	fileRepo  repository.FileRepository
}

func NewShareService(shareRepo repository.ShareRepository) *ShareService {
	return &ShareService{shareRepo: shareRepo}
}

func (s *ShareService) AddFileShares(ctx context.Context, in *service.AddFileSharesInput) error {
	// Verify if user is owner of file
	isOwner, err := s.isUserOwnerOfFile(ctx, &in.FileID, &in.ActorUserID)
	if err != nil {
		return err
	}
	if !isOwner {
		return shared.ErrUnauthorized
	}

	// verify if amount of shares are valid
	if len(in.Recipients) > shared.MAX_SHARES {
		return shared.ErrTooManyShares
	}

	// create shares
	// Extract User IDs of the emails of recipiemnts

	// TODO: Make this a Transaction
	n := len(in.Recipients)
	for i := 0; i < n; i++ {
		p := &model.FileShareParams{
			FileID:           in.FileID,
			SharedWithUserID: in.Recipients[i].Email,
			Permission:       in.Recipients[i].Permission,
		}
		_, err := s.shareRepo.CreateFileShare(ctx, p)
		if err != nil {
			return err
		}
	}

}
