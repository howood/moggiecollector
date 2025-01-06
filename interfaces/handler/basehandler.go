package handler

import (
	"context"
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

func (bh BaseHandler) errorResponse(ctx context.Context, c echo.Context, statudcode int, err error) error {
	if strings.Contains(strings.ToLower(err.Error()), dbcluster.RecordNotFoundMsg) {
		statudcode = http.StatusNotFound
	}
	c.Response().Header().Set(echo.HeaderXRequestID, ctx.Value(echo.HeaderXRequestID).(string))
	return c.JSONPretty(statudcode, map[string]interface{}{"message": err.Error()}, "    ")
}

func (bh BaseHandler) setResponseHeader(c echo.Context, lastmodified, contentlength, xrequestud string) {
	c.Response().Header().Set(echo.HeaderLastModified, lastmodified)
	c.Response().Header().Set(echo.HeaderContentLength, contentlength)
	c.Response().Header().Set(echo.HeaderXRequestID, xrequestud)
}

func (bh BaseHandler) validate(stc interface{}) error {
	val := validator.NewValidator()
	return val.Validate(stc)
}

func (bh BaseHandler) getClaimsFromToken(c echo.Context) *entity.JwtClaims {
	user := c.Get(actor.JWTContextKey).(*jwt.Token)
	claims := user.Claims.(*entity.JwtClaims)
	return claims
}

func (bh BaseHandler) createToken(ctx context.Context, userId uint64, username string) (string, error) {
	tokenstr := actor.NewJwtOperator().CreateToken(ctx, userId, username, false, "moggiecollector-api")
	return tokenstr, nil
}

func (bh BaseHandler) initalGenerateContext(c echo.Context) context.Context {
	ctx := context.WithValue(context.Background(), requestid.GetRequestIDKey(), c.Get(echo.HeaderXRequestID))

	return ctx
}
