package share

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
type SharingService interface {
	AddFileShares(ctx context.Context, in *AddFileSharesInput) error
	UpdateFileShare(ctx context.Context, in *UpdateFileShareInput) error
	DeleteFileShare(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID, recipientUserID uuid.UUID) error
	ListFileShares(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) ([]*model.FileShare, error)
	AddPublicAccess(ctx context.Context, in *AddPublicAccessInput) error
	RemovePublicAccess(ctx context.Context, in *RemovePublicAccessInput) error
}
*/
