package actor

import (
	"context"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/domain/repository"
	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/howood/moggiecollector/library/utils"
)

// tokenExpired is token's expired
//
//nolint:gochecknoglobals
var tokenExpired = utils.GetOsEnv("TOKEN_EXPIED", "3600")

// TokenSecret define token secrets
//
//nolint:gochecknoglobals
var TokenSecret = utils.GetOsEnv("TOKEN_SECRET", "secretsecretdsfdsfsdfdsfsdf")

// JWTContextKey is context key name
const JWTContextKey = "moggiecollector"

// JwtOperator struct
type JwtOperator struct {
	repository.JwtClaimRepository
}

// NewJwtOperator creates a new JwtClaimsRepository
func NewJwtOperator() *JwtOperator {
	return &JwtOperator{
		&jwtCreator{},
	}
}

// jwtCreator struct
type jwtCreator struct{}

// CreateToken creates a new token
func (jc *jwtCreator) CreateToken(ctx context.Context, userID uuid.UUID, username string, admin bool, identifier string) string {
	expired, err := strconv.ParseInt(tokenExpired, 10, 64)
	if err != nil {
		log.Error(ctx, err.Error())
		return ""
	}
	jwtClaims := &entity.JwtClaims{
		Name:       username,
		UserID:     userID,
		Admin:      admin,
		Identifier: identifier,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expired))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenstring, err := token.SignedString([]byte(TokenSecret))
	if err != nil {
		log.Error(ctx, err.Error())
	}
	return tokenstring
}
