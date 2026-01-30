package service

import (
	"context"

	"github.com/google/uuid"
)

// internal
func (p *AuthService) ResolveUserIDs(ctx context.Context, emails []string) (map[string]uuid.UUID, error) {
	return p.dao.ResolveUserIDsByEmails(ctx, emails)
}
