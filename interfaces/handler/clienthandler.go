package handler

import (
	"net/http"

	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/labstack/echo/v4"
)

// ClientHandler struct
type ClientHandler struct {
	BaseHandler
}

// GetProfile is get logined user
func (ch ClientHandler) GetProfile(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := ch.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	claims := ch.getClaimsFromToken(c)
	user, err := ch.UcCluster.ClientUC.GetUserByToken(ctx, claims)
	if err != nil {
		return ch.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, user, marshalIndent)
}
