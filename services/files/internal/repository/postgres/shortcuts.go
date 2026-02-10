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

const CreateShortcutQuery = `
INSERT INTO file_shortcuts (file_uuid, user_id)
VALUES ($1, $2)
RETURNING id, created_at
`

const DeleteShortcutQuery = `
DELETE FROM file_shortcuts
WHERE file_uuid = $1 AND user_id = $2
`

const FetchUsersShortcutsWithFilesQuery = `
SELECT
    f.file_uuid,
    f.user_id,
    f.file_name,
    f.mime_type,
    f.size_bytes,
    f.created_at
FROM file_shortcuts s
JOIN files f ON f.file_uuid = s.file_uuid
WHERE s.user_id = $1
  AND f.deleted_at IS NULL
ORDER BY f.created_at DESC
LIMIT $2 OFFSET $3;

`

type ShortcutsRepository struct {
	db *sql.DB
}

func NewShortcutsRepository(db *sql.DB) *ShortcutsRepository {
	return &ShortcutsRepository{db: db}
}

func (r *ShortcutsRepository) FetchUsersShortcutsWithFiles(
	ctx context.Context,
	userID uuid.UUID,
	limit int,
	offset int,
) ([]*model.FileSummary, error) {

	rows, err := r.db.QueryContext(
		ctx,
		FetchUsersShortcutsWithFilesQuery,
		userID,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*model.FileSummary

	for rows.Next() {
		var fs model.FileSummary

		err := rows.Scan(
			&fs.FileUUID,
			&fs.UserID,
			&fs.Name,
			&fs.MimeType,
			&fs.SizeBytes,
			&fs.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, &fs)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *ShortcutsRepository) CreateShortcut(
	ctx context.Context,
	fileID uuid.UUID,
	userID uuid.UUID,
) (*model.FileShortcut, error) {

	var sc model.FileShortcut

	err := r.db.QueryRowContext(
		ctx,
		CreateShortcutQuery,
		fileID,
		userID,
	).Scan(
		&sc.ID,
		&sc.CreatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, shared.ErrRowExists
		}
		return nil, err
	}

	sc.FileID = fileID
	sc.UserID = userID

	return &sc, nil
}

func (r *ShortcutsRepository) DeleteShortcut(
	ctx context.Context,
	fileUUID uuid.UUID,
	userID uuid.UUID,
) error {
	_, err := r.db.ExecContext(
		ctx,
		DeleteShortcutQuery,
		fileUUID,
		userID,
	)
	return err
}
