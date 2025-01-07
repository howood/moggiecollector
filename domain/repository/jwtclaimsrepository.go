package repository

import "context"

// JwtClaimsRepository interface
type JwtClaimsRepository interface {
	CreateToken(ctx context.Context, userID uint64, username string, admin bool, identifier string) string
}
