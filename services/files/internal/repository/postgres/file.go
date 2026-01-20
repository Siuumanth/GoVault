package postgres

import (
	"context"
	"database/sql"
	"files/internal/model"
	"files/internal/service/files"

	"github.com/google/uuid"
)

/*
type FilesRepository interface {
	GetSingleFile(fileID uuid.UUID) (*files.FileSummary, error)
	ListOwnedFiles(userID uuid.UUID, limit int, offset int) ([]*files.FileSummary, error)
	ListSharedFiles(userID uuid.UUID) ([]*files.FileSummary, error)
	CreateFile(file *files.CreateFileParams) (*model.File, error)
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

const GetSingleFileQuery = `
SELECT file_uuid, user_id, name, mime_type, size, created_at
FROM files
WHERE file_uuid = $1
AND deleted_at IS NULL`

func (r *FilesRepository) GetSingleFile(ctx context.Context, fileID uuid.UUID) (*files.FileSummary, error) {
	var fs files.FileSummary

	err := r.db.QueryRowContext(
		ctx,
		GetSingleFileQuery,
		fileID,
	).Scan(
		&fs.FileUUID,
		&fs.UserID,
		&fs.Name,
		&fs.MimeType,
		&fs.Size,
		&fs.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // not founnd or no access
		}
		return nil, err
	}
	return &fs, nil
}

const ListOwnedFilesQuery = ` 
SELECT file_uuid, user_id, name, mime_type, size, created_at
FROM files
WHERE user_id = $1
AND deleted_at IS NULL
LIMIT $2
OFFSET $3
`

func (r *FilesRepository) ListOwnedFiles(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*files.FileSummary, error) {
	var fs []*files.FileSummary

	rows, err := r.db.QueryContext(ctx, ListOwnedFilesQuery, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var f files.FileSummary
		if err := rows.Scan(
			&f.FileUUID,
			&f.UserID,
			&f.Name,
			&f.MimeType,
			&f.Size,
			&f.CreatedAt,
		); err != nil {
			return nil, err
		}
		fs = append(fs, &f)
	}
	return fs, nil
}

const ListSharedFilesQuery = `
SELECT file_uuid, user_id, name, mime_type, size, created_at
FROM files
JOIN file_shares FS ON
FS.file_uuid = files.file_uuid
WHERE FS.user_id = $1
AND files.deleted_at IS NULL
LIMIT $2
OFFSET $3`

func (r *FilesRepository) ListSharedFiles(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*files.FileSummary, error) {
	var fs []*files.FileSummary

	rows, err := r.db.QueryContext(ctx, ListSharedFilesQuery, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var f files.FileSummary
		if err := rows.Scan(
			&f.FileUUID,
			&f.UserID,
			&f.Name,
			&f.MimeType,
			&f.Size,
			&f.CreatedAt,
		); err != nil {
			return nil, err
		}
		fs = append(fs, &f)
	}
	return fs, nil
}

const CreateFileQuery = `
INSERT INTO files
(
    file_uuid,
    session_id,
    user_id,
    file_name,
    mime_type,
	size_bytes,
	storage_key, 
	checksum
)
VALUES ($1,$2, $3, $4, $5, $6, $7, $8)
RETURNING created_at, deleted_at
`

func (r *FilesRepository) CreateFile(
	ctx context.Context,
	p *files.CreateFileParams,
) (*model.File, error) {

	var file model.File

	err := r.db.QueryRowContext(
		ctx,
		CreateFileQuery,
		p.FileUUID,
		p.SessionID,
		p.UserID,
		p.Name,
		p.MimeType,
		p.SizeBytes,
		p.StorageKey,
		p.Checksum,
	).Scan(
		&file.ID,
		&file.FileUUID,
		&file.SessionID,
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
