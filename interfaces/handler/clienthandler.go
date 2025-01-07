package handler

import (
	"net/http"

	"github.com/howood/moggiecollector/domain/entity"
	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/howood/moggiecollector/interfaces/handler/response"
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
	return c.JSONPretty(http.StatusOK, ch.responseUser(user), marshalIndent)
}

func (ch ClientHandler) responseUser(user *entity.User) response.UserResponse {
	return response.UserResponse{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}
}
