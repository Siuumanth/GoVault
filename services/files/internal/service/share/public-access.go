package share

import (
	"context"
	"files/internal/service"
)

func (s *ShareService) AddPublicAccess(ctx context.Context, in *service.AddPublicAccessInput) error {
	// check owner
	if err := s.assertOwner(ctx, &in.FileID, &in.ActorUserID); err != nil {
		return err
	}

	return s.shareRepo.CreatePublicAccess(ctx, in.FileID)
}

func (s *ShareService) RemovePublicAccess(ctx context.Context, in *service.RemovePublicAccessInput) error {
	// check owner
	if err := s.assertOwner(ctx, &in.FileID, &in.ActorUserID); err != nil {
		return err
	}

	return s.shareRepo.DeletePublicAccess(ctx, in.FileID)
}
