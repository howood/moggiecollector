package repository

import (
	"context"

	"github.com/google/uuid"
)

// JwtClaimsRepository interface
type JwtClaimsRepository interface {
	CreateToken(ctx context.Context, userID uuid.UUID, username string, admin bool, identifier string) string
}
