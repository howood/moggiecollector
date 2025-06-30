package handler

import (
	"errors"
	"net/http"

	"github.com/howood/moggiecollector/domain/dto"
	"github.com/howood/moggiecollector/domain/entity"
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
	var form request.LoginUserForm
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
	verifyMfa, mfaType, err := ah.UcCluster.UserMfaUC.IsUseMfa(ctx, user)
	if err != nil {
		return ah.errorResponse(ctx, c, http.StatusInternalServerError, err)
	}

	switch mfaType {
	case entity.MfaTypeTOTP:
		return c.JSONPretty(http.StatusOK, map[string]interface{}{"identifier": verifyMfa.Identifier, "mfa_type": mfaType}, marshalIndent)
	default:
		token := ah.createToken(ctx, user.ID, user.Email)
		return c.JSONPretty(http.StatusOK, map[string]interface{}{"token": token}, marshalIndent)
	}
}

func (ah AuthHandler) VerifyAuthenticator(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := ah.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	var form request.VerifyAuthenticatorForm
	if err := c.Bind(&form); err != nil {
		log.Error(ctx, "Failed to bind form: ", err)
		return ah.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	if err := ah.validate(form); err != nil {
		log.Error(ctx, "Validation error: ", err)
		return ah.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	isValid, user, err := ah.UcCluster.UserMfaUC.ValidateAuthenticatorCode(ctx, &dto.VerifyMfaAuthenticator{
		Identifier: form.Identifier,
		Passcode:   form.Passcode,
	})
	if err != nil {
		log.Error(ctx, "Failed to validate authenticator code: ", err)
		return ah.errorResponse(ctx, c, http.StatusInternalServerError, err)
	}
	if !isValid {
		return ah.errorResponse(ctx, c, http.StatusUnauthorized, errors.New("invalid passcode"))
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
