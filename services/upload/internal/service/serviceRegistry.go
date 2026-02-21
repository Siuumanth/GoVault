package service

import "upload/internal/service/backend-chunked"

type ServiceRegistry struct {
	backendChunked *backend.UploadService
}
