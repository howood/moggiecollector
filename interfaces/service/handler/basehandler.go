package handler

import (
	"context"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/howood/moggiecollector/application/actor"
	"github.com/howood/moggiecollector/application/validator"
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/interfaces/service/config"
	"github.com/labstack/echo/v4"
)

const marshalIndent = "    "

// BaseHandler struct
type BaseHandler struct {
	ctx context.Context
}

func (bh BaseHandler) errorResponse(c echo.Context, statudcode int, err error) error {
	if strings.Contains(strings.ToLower(err.Error()), config.RecordNotFoundMsg) {
		statudcode = http.StatusNotFound
	}
	c.Response().Header().Set(echo.HeaderXRequestID, bh.ctx.Value(echo.HeaderXRequestID).(string))
	return c.JSONPretty(statudcode, map[string]interface{}{"message": err.Error()}, "    ")
}

func (bh BaseHandler) setResponseHeader(c echo.Context, lastmodified, contentlength, xrequestud string) {
	c.Response().Header().Set(echo.HeaderLastModified, lastmodified)
	c.Response().Header().Set(echo.HeaderContentLength, contentlength)
	c.Response().Header().Set(echo.HeaderXRequestID, xrequestud)
}

func (bh BaseHandler) validate(stc interface{}) error {
	val := validator.NewValidator(bh.ctx)
	return val.Validate(stc)
}

func (bh BaseHandler) getClaimsFromToken(c echo.Context) *entity.JwtClaims {
	user := c.Get(actor.JWTContextKey).(*jwt.Token)
	claims := user.Claims.(*entity.JwtClaims)
	return claims
}

func (ch AccountHandler) createToken(userId uint64, username string) (string, error) {
	jwtinstance := actor.NewJwtOperator(ch.ctx, userId, username, false, "moggiecollector-api")
	tokenstr := jwtinstance.CreateToken(actor.TokenSecret)
	return tokenstr, nil
}
