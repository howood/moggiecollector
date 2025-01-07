package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/howood/moggiecollector/application/actor"
	"github.com/howood/moggiecollector/application/validator"
	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/di/uccluster"
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/infrastructure/requestid"
	"github.com/labstack/echo/v4"
)

const marshalIndent = "    "

// BaseHandler struct
type BaseHandler struct {
	UcCluster *uccluster.UsecaseCluster
}

//nolint:unparam
func (bh BaseHandler) errorResponse(ctx context.Context, c echo.Context, statudcode int, err error) error {
	if strings.Contains(strings.ToLower(err.Error()), dbcluster.RecordNotFoundMsg) {
		statudcode = http.StatusNotFound
	}

	c.Response().Header().Set(echo.HeaderXRequestID, fmt.Sprintf("%v", ctx.Value(echo.HeaderXRequestID)))
	return c.JSONPretty(statudcode, map[string]interface{}{"message": err.Error()}, "    ")
}

func (bh BaseHandler) validate(stc interface{}) error {
	val := validator.NewValidator()
	return val.Validate(stc)
}

//nolint:forcetypeassert
func (bh BaseHandler) getClaimsFromToken(c echo.Context) *entity.JwtClaims {
	user := c.Get(actor.JWTContextKey).(*jwt.Token)
	claims := user.Claims.(*entity.JwtClaims)
	return claims
}

func (bh BaseHandler) createToken(ctx context.Context, userID uint64, username string) string {
	tokenstr := actor.NewJwtOperator().CreateToken(ctx, userID, username, false, "moggiecollector-api")
	return tokenstr
}

func (bh BaseHandler) initalGenerateContext(c echo.Context) context.Context {
	ctx := context.WithValue(context.Background(), requestid.GetRequestIDKey(), c.Get(echo.HeaderXRequestID))

	return ctx
}
