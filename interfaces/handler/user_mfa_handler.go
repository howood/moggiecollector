package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/howood/moggiecollector/domain/dto"
	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/howood/moggiecollector/interfaces/handler/request"
	"github.com/howood/moggiecollector/interfaces/handler/response"
	"github.com/labstack/echo/v4"
)

// UserMfaHandler struct
type UserMfaHandler struct {
	BaseHandler
}

func (uh UserMfaHandler) InitialUserAuthenticator(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := uh.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	userid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	secret, err := uh.UcCluster.UserMfaUC.GetAuthenticatorSecret(ctx, userid)
	if err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, response.InitialAuthenticatorResponse{Secret: secret}, marshalIndent)
}

func (uh UserMfaHandler) UpsertUserAuthenticator(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := uh.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	userid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	var form request.UpsertUserAuthenticatorForm
	if err := c.Bind(&form); err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	if err := uh.validate(form); err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	if err := uh.UcCluster.UserMfaUC.UpsertAuthenticator(ctx, &dto.UserMfaTotpDto{
		UserID:    userid,
		Secret:    form.Secret,
		Passcode:  form.Passcode,
		IsDefault: form.IsDefault,
	}); err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"message": "success"}, marshalIndent)
}
