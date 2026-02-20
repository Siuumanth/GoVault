package postgres

import (
	"context"
	"database/sql"
	"errors"
	model "files/internal/model"
	"files/internal/shared"

	"github.com/google/uuid"
)

/*
type FilesRepository interface {
	GetSingleFile(fileID uuid.UUID) (*model.FileSummary, error)
	ListOwnedFiles(userID uuid.UUID, limit int, offset int) ([]*model.FileSummary, error)
	ListSharedFiles(userID uuid.UUID) ([]*model.FileSummary, error)
	CreateFile(file *model.CreateFileParams) (*model.File, error)
	FetchOwnerIDByFileID(ctx context.Context, fileID uuid.UUID) error
}

type FileSummary struct {
    FileUUID        uuid.UUID
	UserID    uuid.UUID
	Name      stringh
	MimeType  string
	Size      int64
	CreatedAt time.Time
}
*/

type FilesRepository struct {
	db *sql.DB
}

func NewFilesRepository(db *sql.DB) *FilesRepository {
	return &FilesRepository{db: db}
}

const CheckFileOwnershipQuery = `
SELECT EXISTS (
	SELECT 1
	FROM files
	WHERE file_uuid = $1
	  AND user_id = $2
	  AND deleted_at IS NULL
)
`

func (r *FilesRepository) CheckFileOwnership(
	ctx context.Context,
	fileID uuid.UUID,
	userID uuid.UUID,
) error {

	var exists bool
	err := r.db.QueryRowContext(
		ctx,
		CheckFileOwnershipQuery,
		fileID,
		userID,
	).Scan(&exists)

	if err != nil {
		if !exists {
			return shared.ErrUnauthorized
		}
		return err
	}
	return nil

}

const UpdateFileNameQuery = `
		UPDATE files
		SET file_name = $1
		WHERE file_uuid = $2
		AND deleted_at IS NULL
	`

func (r *FilesRepository) UpdateFileName(ctx context.Context, fileUUID uuid.UUID, newName string) (bool, error) {

	result, err := r.db.ExecContext(
		ctx,
		UpdateFileNameQuery,
		newName,
		fileUUID,
	)
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil // file not found or soft-deleted
	}

	return true, nil
}

const GetSingleFileQuery = `
SELECT 
    f.file_uuid, 
    f.user_id, 
    f.file_name, 
    f.mime_type, 
    f.size_bytes, 
    f.created_at,
    (p.file_uuid IS NOT NULL) as is_public
FROM files f
LEFT JOIN public_files p ON f.file_uuid = p.file_uuid
WHERE f.file_uuid = $1
AND f.deleted_at IS NULL`

func (r *FilesRepository) FetchFileSummaryByID(ctx context.Context, fileID uuid.UUID) (*model.FileSummary, error) {
	var fs model.FileSummary

	err := r.db.QueryRowContext(
		ctx,
		GetSingleFileQuery,
		fileID,
	).Scan(
		&fs.FileUUID,
		&fs.UserID,
		&fs.Name,
		&fs.MimeType,
		&fs.SizeBytes,
		&fs.CreatedAt,
		&fs.IsPublic,
	)

	if err != nil {
		return nil, err
	}

	return &fs, nil
}

const CreateFileQuery = `
INSERT INTO files
(
    file_uuid,
    upload_uuid,
    user_id,
    file_name,
    mime_type,
	size_bytes,
	storage_key, 
	checksum
)
VALUES ($1,$2, $3, $4, $5, $6, $7, $8)
RETURNING id, file_uuid, upload_uuid, user_id, file_name,
          mime_type, size_bytes, storage_key, checksum,
          created_at, deleted_at
`

func (r *FilesRepository) CreateFile(ctx context.Context, p *model.CreateFileParams) (*model.File, error) {
	var file model.File

	err := r.db.QueryRowContext(
		ctx,
		CreateFileQuery,
		p.FileUUID,
		p.UploadUUID,
		p.UserID,
		p.Name,
		p.MimeType,
		p.SizeBytes,
		p.StorageKey,
		p.Checksum,
	).Scan(
		&file.ID,
		&file.FileUUID,
		&file.UploadUUID,
		&file.UserID,
		&file.FileName,
		&file.MimeType,
		&file.SizeBytes,
		&file.StorageKey,
		&file.Checksum,
		&file.CreatedAt,
		&file.DeletedAt,
	)

	return &file, err
}

const DeleteFileQuery = `
UPDATE files
SET deleted_at = now()
WHERE file_uuid = $1
`

func (r *FilesRepository) SoftDeleteFile(ctx context.Context, fileID uuid.UUID) error {
	_, err := r.db.ExecContext(
		ctx,
		DeleteFileQuery,
		fileID,
	)

	return err
}

const GetFullFileByIDQuery = `
SELECT id, upload_uuid,file_uuid, user_id, file_name, mime_type, size_bytes, storage_key, checksum, created_at, deleted_at
FROM files
WHERE file_uuid = $1
AND deleted_at IS NULL
`

func (r *FilesRepository) FetchFullFileByID(ctx context.Context, fileID uuid.UUID) (*model.File, error) {
	var file model.File

	err := r.db.QueryRowContext(
		ctx,
		GetFullFileByIDQuery,
		fileID,
	).Scan(
		&file.ID,
		&file.UploadUUID,
		&file.FileUUID,
		&file.UserID,
		&file.FileName,
		&file.MimeType,
		&file.SizeBytes,
		&file.StorageKey,
		&file.CreatedAt,
		&file.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, shared.ErrRowNotFound) {
			return nil, shared.ErrRowNotFound
		} else {
			return nil, err
		}
	}

	return &file, nil
}
