package repository

import (
	"context"

	"github.com/google/uuid"
)

// JwtClaimRepository interface
type JwtClaimRepository interface {
	CreateToken(ctx context.Context, userID uuid.UUID, username string, admin bool, identifier string) string
}
