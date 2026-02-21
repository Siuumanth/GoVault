package postgres

import (
	"context"
	"database/sql"
	"upload/internal/model"

	"github.com/google/uuid"
)

/*
type UploadSessionRepository interface {
	CreateSession(session *model.UploadSession) error
	GetSessionByID(session_id int) (*model.UploadSession, error)
	GetSessionByUUID(upload_uuid uuid.UUID) (*model.UploadSession, error)
	UpdateUploadStatus(session_id int, status string) error
}
*/

type PGUploadSessionRepo struct {
	db *sql.DB
}

func NewUploadSessionRepo(db *sql.DB) *PGUploadSessionRepo {
	return &PGUploadSessionRepo{db: db}
}

const CreateSessionQuery = `
INSERT INTO upload_sessions 
(upload_uuid, user_id, file_name, file_size_bytes, total_chunks) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING id, upload_uuid, user_id, file_name, file_size_bytes, total_chunks, status, created_at
`

func (p *PGUploadSessionRepo) CreateSession(
	ctx context.Context,
	session *model.UploadSession,
) (*model.UploadSession, error) {

	err := p.db.QueryRowContext(
		ctx,
		CreateSessionQuery,
		session.UploadUUID,
		session.UserID,
		session.FileName,
		session.FileSize,
		session.TotalParts,
	).Scan(
		&session.ID,
		&session.UploadUUID,
		&session.UserID,
		&session.FileName,
		&session.FileSize,
		&session.TotalParts,
		&session.Status,
		&session.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return session, nil
}

const GetSesssionByIDQUery = `SELECT id, upload_uuid, user_id, file_name, file_size_bytes, total_chunks FROM upload_sessions WHERE id = $1`

func (p *PGUploadSessionRepo) GetSessionByID(ctx context.Context, session_id int64) (*model.UploadSession, error) {
	var session model.UploadSession
	err := p.db.QueryRowContext(
		ctx, GetSesssionByIDQUery, session_id).Scan(
		&session.ID,
		&session.UploadUUID,
		&session.UserID,
		&session.FileName,
		&session.FileSize,
		&session.TotalParts,
	)
	return &session, err
}

const GetSesionByUUIDQuery = `SELECT id, upload_uuid, user_id, file_name, file_size_bytes, total_chunks, status, created_at FROM upload_sessions WHERE upload_uuid = $1`

func (p *PGUploadSessionRepo) GetSessionByUUID(ctx context.Context, upload_uuid uuid.UUID) (*model.UploadSession, error) {
	var session model.UploadSession
	err := p.db.QueryRowContext(
		ctx, GetSesionByUUIDQuery, upload_uuid).Scan(
		&session.ID,
		&session.UploadUUID,
		&session.UserID,
		&session.FileName,
		&session.FileSize,
		&session.TotalParts,
		&session.Status,
		&session.CreatedAt,
	)
	return &session, err
}

const UpdateSessionStatusQuery = `UPDATE upload_sessions SET status = $1 WHERE id = $2`

func (p *PGUploadSessionRepo) UpdateSessionStatus(ctx context.Context, sessionID int64, status string) error {
	_, err := p.db.ExecContext(ctx, UpdateSessionStatusQuery, status, sessionID)
	return err
}

const DeleteSessionQuery = `DELETE FROM upload_sessions WHERE id = $1`

func (p *PGUploadSessionRepo) DeleteSessionChunks(ctx context.Context, sessionID int64) error {
	_, err := p.db.ExecContext(ctx, DeleteSessionQuery, sessionID)
	return err
}
