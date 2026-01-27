package dto

import "time"

type AddFileSharesRequest struct {
	Recipients []ShareRecipientRequest `json:"recipients"`
}

type ShareRecipientRequest struct {
	Email      string `json:"email"`
	Permission string `json:"permission"`
}

type UpdateFileShareRequest struct {
	Permission string `json:"permission"`
}

type FileShareResponse struct {
	UserID     string    `json:"user_id"`
	Permission string    `json:"permission"`
	CreatedAt  time.Time `json:"created_at"`
}
