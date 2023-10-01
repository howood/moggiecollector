package actor

import (
	"context"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/domain/repository"
	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/howood/moggiecollector/library/utils"
)

// tokenExpired is token's expired
var tokenExpired = utils.GetOsEnv("TOKEN_EXPIED", "3600")

// TokenSecret define token secrets
var TokenSecret = utils.GetOsEnv("TOKEN_SECRET", "secretsecretdsfdsfsdfdsfsdf")

// JWTContextKey is context key name
const JWTContextKey = "moggiecollector"

// JwtOperator struct
type JwtOperator struct {
	repository.JwtClaimsRepository
}

// NewJwtOperator creates a new JwtClaimsRepository
func NewJwtOperator(ctx context.Context, userId uint64, username string, admin bool, identifier string) *JwtOperator {
	expired, _ := strconv.ParseInt(tokenExpired, 10, 64)
	return &JwtOperator{
		&jwtCreator{
			jwtClaims: &entity.JwtClaims{
				Name:       username,
				UserID:     userId,
				Admin:      admin,
				Identifier: identifier,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expired))),
				},
			},
			ctx: ctx,
		},
	}
}

// jwtCreator struct
type jwtCreator struct {
	jwtClaims *entity.JwtClaims
	ctx       context.Context
}

// CreateToken creates a new token
func (jc *jwtCreator) CreateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jc.jwtClaims)
	tokenstring, err := token.SignedString([]byte(TokenSecret))
	if err != nil {
		log.Error(jc.ctx, err.Error())
	}
	return tokenstring
}
