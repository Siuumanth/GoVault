package model

import (
	"time"

	"github.com/google/uuid"
)

// these are models returned by the repository and the actual schema

type UploadSession struct {
	ID              int64     // internal BIGSERIAL
	UploadUUID      uuid.UUID // public ID
	UserID          uuid.UUID
	FileName        string
	FileSize        int64
	TotalParts      int64 // TODO: Change to total parts
	Status          string
	UploadMethod    string  // proxy or multipart
	StorageUploadID *string // nullable so ptr
	CreatedAt       time.Time
}

type UploadChunk struct {
	ID         int64
	SessionID  int64 // upload sessions id
	ChunkIndex int
	SizeBytes  int64
	CheckSum   string
	UploadedAt time.Time
}

type UploadPart struct {
	ID         int64
	SessionID  int64 // upload sessions id
	PartNumber int
	SizeBytes  int64
	Etag       string
	UploadedAt time.Time
}

// type File struct {
// 	ID         int
// 	FileUUID   uuid.UUID
// 	UserID     uuid.UUID
// 	SessionID  int
// 	Name       string
// 	MimeType   string
// 	SizeBytes  int64
// 	CheckSum   *string
// 	StorageKey string
// 	CreatedAt  time.Time
// }
