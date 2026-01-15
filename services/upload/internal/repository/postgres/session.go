package postgres

import (
	"database/sql"
	"upload/internal/model"

	"github.com/google/uuid"
)

/*
type UploadSessionRepository interface {
	CreateUploadSession(session *model.UploadSession) error
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

const CreateSessionQuery = `INSERT INTO upload_sessions (upload_uuid, user_id, file_name, file_size_bytes, total_chunks) VALUES ($1, $2, $3, $4, $5) RETURNING id`

func (p *PGUploadSessionRepo) CreateSession(session *model.UploadSession) error {
	err := p.db.QueryRow(
		CreateSessionQuery,
		session.UploadUUID,
		session.UserID,
		session.FileName,
		session.FileSize,
		session.TotalChunks,
	).Scan(&session.ID)
	return err
}

const GetSesssionByIDQUery = `SELECT id, upload_uuid, user_id, file_name, file_size_bytes, total_chunks FROM upload_sessions WHERE id = $1`

func (p *PGUploadSessionRepo) GetSessionByID(session_id int) (*model.UploadSession, error) {
	var session model.UploadSession
	err := p.db.QueryRow(GetSesssionByIDQUery, session_id).Scan(
		&session.ID,
		&session.UploadUUID,
		&session.UserID,
		&session.FileName,
		&session.FileSize,
		&session.TotalChunks,
	)
	return &session, err
}

const GetSesionByUUIDQuery = `SELECT id, upload_uuid, user_id, file_name, file_size_bytes, total_chunks, status, created_at FROM upload_sessions WHERE upload_uuid = $1`

func (p *PGUploadSessionRepo) GetSessionByUUID(upload_uuid uuid.UUID) (*model.UploadSession, error) {
	var session model.UploadSession
	err := p.db.QueryRow(GetSesionByUUIDQuery, upload_uuid).Scan(
		&session.ID,
		&session.UploadUUID,
		&session.UserID,
		&session.FileName,
		&session.FileSize,
		&session.TotalChunks,
		&session.Status,
		&session.CreatedAt,
	)
	return &session, err
}

const UpdateSessionStatusQuery = `UPDATE upload_sessions SET status = $1 WHERE id = $2`

func (p *PGUploadSessionRepo) UpdateSessionStatus(sessionID int, status string) error {
	_, err := p.db.Exec(UpdateSessionStatusQuery, status, sessionID)
	return err
}

const DeleteSessionQuery = `DELETE FROM upload_sessions WHERE id = $1`

func (p *PGUploadSessionRepo) DeleteSessionChunks(sessionID int) error {
	_, err := p.db.Exec(DeleteSessionQuery, sessionID)
	return err
}
