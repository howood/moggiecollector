package entity

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JwtClaims entity
type JwtClaims struct {
	Name       string    `json:"name"`
	UserID     uuid.UUID `json:"user_id"`
	Admin      bool      `json:"admin"`
	Identifier string    `json:"identifier"`
	jwt.RegisteredClaims
}
