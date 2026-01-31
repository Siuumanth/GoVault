package shared

import "time"

const (
	S3UsersURI          = "s3://govault-files/users/"
	S3UsersPrefix       = "users/"
	MAX_SHARES      int = 5
	ActorIDKey          = "actor_id"
	PAGE_LIMIT          = 20
	PAGE_NO_DEFAULT     = 1

	DOWNLOAD_LINK_TTL = 300 * time.Second
)
