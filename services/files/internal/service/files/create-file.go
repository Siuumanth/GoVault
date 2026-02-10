package files

import (
	"context"
	"files/internal/model"
)

func (s *FileService) CreateFile(ctx context.Context, in *model.CreateFileParams) (*model.File, error) {
	return s.fileRepo.CreateFile(ctx, in)
}
