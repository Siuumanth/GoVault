package service

import (
	"fmt"
	"os"
	"path/filepath"
	"upload/internal/model"
	"upload/shared"

	"github.com/google/uuid"
)

func (s *UploadService) UploadSession(inputs *UploadSessionInput) (*model.UploadSession, error) {
	/*
	   - calculate total chunks
	   - insert session row to uploadSession table
	   - get session ID , make folder for tat session
	   - return UUID, 200, saying that session is created
	*/

	var session model.UploadSession
	// assume there are no missing fields
	// fill upload session model
	session.UploadUUID = uuid.New()
	session.FileName = inputs.FileName
	session.FileSize = inputs.FileSizeBytes
	session.UserID = inputs.UserID
	session.TotalChunks = calculateTotalChunks(inputs.FileSizeBytes)

	// insert session into database
	err := s.registry.Sessions.CreateSession(&session)
	if err != nil {
		return nil, err
	}

	// create session folder
	_, err = createSessionDir(session.ID)
	if err != nil {
		// delete row if fail
		s.registry.Sessions.UpdateSessionStatus(session.ID, "failed")
		return nil, err
	}

	return &session, nil
}

func calculateTotalChunks(fileSize int64) int {
	return int((fileSize + shared.ChunkSizeBytes - 1) / shared.ChunkSizeBytes)
}

func createSessionDir(sessionID int) (string, error) {
	dir := filepath.Join(shared.UploadBasePath, fmt.Sprintf("%d", sessionID))

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	return dir, nil
}

func deleteSessionDir(sessionID int) error {
	dir := filepath.Join(shared.UploadBasePath, fmt.Sprintf("%d", sessionID))

	return os.RemoveAll(dir)
}
