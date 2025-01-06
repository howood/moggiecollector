package handler

import (
	"net/http"
	"strconv"

	"github.com/howood/moggiecollector/domain/form"
	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/labstack/echo/v4"
)

// AccountHandler struct
type AccountHandler struct {
	BaseHandler
}

// GetUsers is get all users
func (ch AccountHandler) GetUsers(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := ch.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	withinactive := c.QueryParam("withinactive")
	users, err := ch.UcCluster.AccountUC.GetUsers(ctx, withinactive)
	if err != nil {
		return ch.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, users, marshalIndent)
}

// GetUser is get all users
func (ch AccountHandler) GetUser(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := ch.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	userid, _ := strconv.Atoi(c.Param("id"))
	user, err := ch.UcCluster.AccountUC.GetUser(ctx, userid)
	if err != nil {
		return ch.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, user, marshalIndent)
}

// CreateUser is get all users
func (ch AccountHandler) CreateUser(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := ch.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	form := form.CreateUserForm{}
	var err error
	if err == nil {
		err = c.Bind(&form)
	}
	if err == nil {
		err = ch.validate(form)
	}
	if err != nil {
		return ch.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	err = ch.UcCluster.AccountUC.CreateUser(ctx, form)
	if err != nil {
		return ch.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"message": "success"}, marshalIndent)
}

// CreateUser is get all users
func (ch AccountHandler) UpdateUser(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := ch.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	userid, _ := strconv.Atoi(c.Param("id"))
	form := form.CreateUserForm{}
	var err error
	if err == nil {
		err = c.Bind(&form)
	}
	if err == nil {
		err = ch.validate(form)
	}
	if err != nil {
		return ch.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	err = ch.UcCluster.AccountUC.UpdateUser(ctx, userid, form)
	if err != nil {
		return ch.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"message": "success"}, marshalIndent)
}

// InActiveUser is get all users
func (ch AccountHandler) InActiveUser(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := ch.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	userid, _ := strconv.Atoi(c.Param("id"))
	err := ch.UcCluster.AccountUC.InActiveUser(ctx, userid)
	if err != nil {
		return ch.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"message": "success"}, marshalIndent)
}

// Login is Login user
func (ch AccountHandler) Login(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := ch.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	form := form.LoginUserForm{}
	var token string
	var err error
	if err == nil {
		err = c.Bind(&form)
	}
	if err == nil {
		err = ch.validate(form)
	}
	if err != nil {
		return ch.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	user, err := ch.UcCluster.AccountUC.AuthUser(ctx, form)
	if err != nil {
		return ch.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	if err == nil {
		token, err = ch.createToken(ctx, user.UserID, user.Email)
	}
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"token": token}, marshalIndent)
}
