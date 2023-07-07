package entity

import jwt "github.com/golang-jwt/jwt/v5"

// JwtClaims entity
type JwtClaims struct {
	Name       string `json:"name"`
	UserID     uint64 `json:"user_id"`
	Admin      bool   `json:"admin"`
	Identifier string `json:"identifier"`
	jwt.RegisteredClaims
}
