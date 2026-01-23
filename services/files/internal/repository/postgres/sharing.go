package postgres

import (
	"context"
	"database/sql"
	"errors"
	"files/internal/model"
	"files/internal/shared"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

/*
type ShareRepository interface {
	CreateFileShare(ctx context.Context, p *share.FileShareParams) (*model.FileShare, error)
	FetchUserFileShare(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) error
	DeleteFileShare(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) error
	UpdateFileShare(ctx context.Context, p *share.FileShareParams) error
	FetchFileShares(ctx context.Context, fileID uuid.UUID) ([]*model.FileShare, error)
	CreatePublicAccess(ctx context.Context, fileID uuid.UUID) error
	DeletePublicAccess(ctx context.Context, fileID uuid.UUID) error
}
*/

type FileShareRepository struct {
	db *sql.DB
}

const FetchUserFileShareQuery = `
SELECT id, file_id, shared_with_user_id, permission, created_at
FROM file_shares
WHERE file_id = $1 AND shared_with_user_id = $2`

func (r *FileShareRepository) FetchUserFileShare(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) (*model.FileShare, error) {

	var share model.FileShare

	err := r.db.QueryRowContext(
		ctx,
		FetchUserFileShareQuery,
		fileID,
		userID,
	).Scan(
		&share.ID,
		&share.FileID,
		&share.SharedWithUserID,
		&share.Permission,
		&share.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, shared.ErrRowNotFound
		}
		return nil, err
	}
	return &share, nil
}

const CreateFileShareQuery = `
INSERT INTO file_shares (file_id, shared_with_user_id, permission) VALUES ($1, $2, $3) RETURNING id, created_at
`

func (r *FileShareRepository) CreateFileShare(ctx context.Context, p *model.FileShareParams) (*model.FileShare, error) {
	var share model.FileShare
	err := r.db.QueryRowContext(
		ctx,
		CreateFileShareQuery,
		p.FileID,
		p.SharedWithUserID,
		p.Permission,
	).Scan(
		&share.ID,
		&share.CreatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, shared.ErrRowExists
		}
		return nil, err
	}

	return &share, nil
}

const DeleteFileShareQuery = `DELETE * from file_shares WHERE file_id = $1 && shared_with_user_id = $2`

func (r *FileShareRepository) DeleteFileShare(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) error {
	rows := r.db.QueryRowContext(
		ctx,
		DeleteFileShareQuery,
		fileID,
		userID,
	)

	return rows.Err()
}

// FetchFileShares(ctx context.Context, fileID uuid.UUID) ([]*model.FileShare, error)
const FetchFileSharesQuery = `
SELECT file_id, shared_with_user_id, permission from file_shares
WHERE file_id = $1
`

func (r *FileShareRepository) FetchAllFileShares(ctx context.Context, fileID uuid.UUID) ([]*model.FileShare, error) {
	var res []*model.FileShare
	rows, err := r.db.QueryContext(
		ctx,
		FetchFileSharesQuery,
		fileID,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		fs := new(model.FileShare)
		if err := rows.Scan(
			&fs.FileID,
			&fs.SharedWithUserID,
			&fs.Permission,
		); err != nil {
			return nil, err
		}
		res = append(res, fs)
	}
	return res, nil

}

const UpdateFileShareQuery = `
UPDATE file_shares 
SET permission = $1
WHERE 
    file_id = $2 AND
	shared_with_user_id = $3
`

func (r *FileShareRepository) UpdateFileShare(ctx context.Context, p *model.FileShareParams) error {
	_, err := r.db.ExecContext(
		ctx,
		UpdateFileShareQuery,
		p.Permission,
		p.FileID,
		p.SharedWithUserID,
	)
	return err
}

const IsFileSharedWithUserQuery = `
SELECT EXISTS (
	SELECT 1
	FROM file_shares
	WHERE file_id = $1 AND shared_with_user_id = $2
)
`

func (r *FileShareRepository) IsFileSharedWithUser(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) (bool, error) {
	var res bool
	err := r.db.QueryRowContext(
		ctx,
		IsFileSharedWithUserQuery,
		fileID,
		userID,
	).Scan(&res)
	return res, err
}

// PUBLIC ACCESS METHODS
const FetchPublicAccessQuery = `
SELECT EXISTS (
    SELECT 1
    FROM public_files
    WHERE file_id = $1
)
`

func (r *FileShareRepository) IsFilePublic(ctx context.Context, fileID uuid.UUID) (bool, error) {
	var res bool
	err := r.db.QueryRowContext(
		ctx,
		FetchPublicAccessQuery,
		fileID,
	).Scan(&res)
	return res, err
}

const CreatePublicAccessQuery = `
INSERT INTO file_shortcuts(file_id) 
VALUES($1)
`

func (r *FileShareRepository) CreatePublicAccess(ctx context.Context, fileID uuid.UUID) error {
	_, err := r.db.ExecContext(
		ctx,
		CreatePublicAccessQuery,
		fileID,
	)
	return err
}

const DeletePublicAccessQuery = `
DELETE from file_shortcuts
WHERE file_id = $1
`

func (r *FileShareRepository) DeletePublicAccess(ctx context.Context, fileID uuid.UUID) error {
	_, err := r.db.ExecContext(
		ctx,
		DeletePublicAccessQuery,
		fileID,
	)
	return err
}
