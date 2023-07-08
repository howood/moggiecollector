package handler

import (
	"context"
	"net/http"

	"github.com/howood/moggiecollector/application/usecase"
	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/howood/moggiecollector/infrastructure/requestid"
	"github.com/labstack/echo/v4"
)

// ClientHandler struct
type ClientHandler struct {
	BaseHandler
}

// GetProfile is get logined user
func (ch ClientHandler) GetProfile(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	xRequestID := requestid.GetRequestID(c.Request())
	ch.ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, xRequestID)
	log.Info(ch.ctx, "========= START REQUEST : "+requesturi)
	log.Info(ch.ctx, c.Request().Method)
	log.Info(ch.ctx, c.Request().Header)
	claims := ch.getClaimsFromToken(c)
	user, err := usecase.ClientUsecase{}.GetUserByToken(claims)
	if err != nil {
		return ch.errorResponse(c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, user, marshalIndent)
}
