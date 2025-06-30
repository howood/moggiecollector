package handler

import (
	"net/http"

	"github.com/howood/moggiecollector/domain/dto"
	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/howood/moggiecollector/interfaces/handler/request"
	"github.com/labstack/echo/v4"
)

// AuthHandler struct
type AuthHandler struct {
	BaseHandler
}

// Login is Login user
func (ah AuthHandler) Login(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := ah.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	form := request.LoginUserForm{}
	if err := c.Bind(&form); err != nil {
		return ah.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	if err := ah.validate(form); err != nil {
		return ah.errorResponse(ctx, c, http.StatusBadRequest, err)
	}

	user, err := ah.UcCluster.AuthUC.AuthUser(ctx, ah.convertToLoginrDto(form))
	if err != nil {
		return ah.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	token := ah.createToken(ctx, user.ID, user.Email)

	return c.JSONPretty(http.StatusOK, map[string]interface{}{"token": token}, marshalIndent)
}

func (ah AuthHandler) convertToLoginrDto(user request.LoginUserForm) *dto.LoginDto {
	return &dto.LoginDto{
		Email:    user.Email,
		Password: user.Password,
	}
}
