package service

import "github.com/google/uuid"

// ---------- Files ----------
type UpdateFileNameInput struct {
	FileID      uuid.UUID
	ActorUserID uuid.UUID
	NewName     string
}

type MakeFileCopyInput struct {
	FileID      uuid.UUID
	ActorUserID uuid.UUID
}

type ListOwnedFilesInput struct {
	UserID uuid.UUID
	Limit  int
	Offset int
}

type ListSharedFilesInput struct {
	UserID uuid.UUID
	Limit  int
	Offset int
}

// ---------- Sharing ----------
type AddFileSharesInput struct {
	FileID      uuid.UUID
	ActorUserID uuid.UUID
	Recipients  []ShareRecipientInput
}

type ShareRecipientInput struct {
	Email      string
	Permission string
}

type UpdateFileShareInput struct {
	FileID          uuid.UUID
	ActorUserID     uuid.UUID
	RecipientUserID uuid.UUID
	Permission      string
}

// ---------- Public Access ----------
type AddPublicAccessInput struct {
	FileID      uuid.UUID
	ActorUserID uuid.UUID
}

type RemovePublicAccessInput struct {
	FileID      uuid.UUID
	ActorUserID uuid.UUID
}

// ---------- Shortcuts ----------
type CreateShortcutInput struct {
	FileID      uuid.UUID
	ActorUserID uuid.UUID
}

type DeleteShortcutInput struct {
	ShortcutID  uuid.UUID
	ActorUserID uuid.UUID
}
