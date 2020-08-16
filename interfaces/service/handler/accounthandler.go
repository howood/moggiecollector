package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/howood/moggiecollector/domain/entity"
	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/howood/moggiecollector/infrastructure/requestid"
	"github.com/howood/moggiecollector/interfaces/service/config"
	"github.com/labstack/echo/v4"
)

// ClientHandler struct
type AccountHandler struct {
	BaseHandler
}

// GetUsers is get all users
func (ch AccountHandler) GetUsers(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	xRequestID := requestid.GetRequestID(c.Request())
	ch.ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, xRequestID)
	log.Info(ch.ctx, "========= START REQUEST : "+requesturi)
	log.Info(ch.ctx, c.Request().Method)
	log.Info(ch.ctx, c.Request().Header)
	withinactive := c.QueryParam("withinactive")
	var err error
	var users []entity.User
	if withinactive == "true" {
		users, err = config.GetDataStore().User.GetAllWithInActive()
	} else {
		users, err = config.GetDataStore().User.GetAll()
	}
	if err != nil {
		return ch.errorResponse(c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, users, marshalIndent)
}

// GetUser is get all users
func (ch AccountHandler) GetUser(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	xRequestID := requestid.GetRequestID(c.Request())
	ch.ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, xRequestID)
	log.Info(ch.ctx, "========= START REQUEST : "+requesturi)
	log.Info(ch.ctx, c.Request().Method)
	log.Info(ch.ctx, c.Request().Header)
	userid, _ := strconv.Atoi(c.Param("id"))
	user, err := config.GetDataStore().User.Get(uint64(userid))
	if err != nil {
		return ch.errorResponse(c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, user, marshalIndent)
}

// CreateUser is get all users
func (ch AccountHandler) CreateUser(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	xRequestID := requestid.GetRequestID(c.Request())
	ch.ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, xRequestID)
	log.Info(ch.ctx, "========= START REQUEST : "+requesturi)
	log.Info(ch.ctx, c.Request().Method)
	log.Info(ch.ctx, c.Request().Header)
	form := entity.CreateUserForm{}
	var err error
	if err == nil {
		err = c.Bind(&form)
	}
	if err == nil {
		err = ch.validate(form)
	}
	if err != nil {
		return ch.errorResponse(c, http.StatusBadRequest, err)
	}
	if err := config.GetDataStore().User.Create(form.Name, form.Email, form.Password); err != nil {
		return ch.errorResponse(c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"message": "success"}, marshalIndent)
}

// CreateUser is get all users
func (ch AccountHandler) UpdateUser(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	xRequestID := requestid.GetRequestID(c.Request())
	ch.ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, xRequestID)
	log.Info(ch.ctx, "========= START REQUEST : "+requesturi)
	log.Info(ch.ctx, c.Request().Method)
	log.Info(ch.ctx, c.Request().Header)
	userid, _ := strconv.Atoi(c.Param("id"))
	form := entity.CreateUserForm{}
	var err error
	if err == nil {
		err = c.Bind(&form)
	}
	if err == nil {
		err = ch.validate(form)
	}
	if err != nil {
		return ch.errorResponse(c, http.StatusBadRequest, err)
	}
	if err := config.GetDataStore().User.Update(uint64(userid), form.Name, form.Email, form.Password); err != nil {
		return ch.errorResponse(c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"message": "success"}, marshalIndent)
}

// InActiveUser is get all users
func (ch AccountHandler) InActiveUser(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	xRequestID := requestid.GetRequestID(c.Request())
	ch.ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, xRequestID)
	log.Info(ch.ctx, "========= START REQUEST : "+requesturi)
	log.Info(ch.ctx, c.Request().Method)
	log.Info(ch.ctx, c.Request().Header)
	userid, _ := strconv.Atoi(c.Param("id"))
	if err := config.GetDataStore().User.InActive(uint64(userid)); err != nil {
		return ch.errorResponse(c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"message": "success"}, marshalIndent)
}

// Login is Login user
func (ch AccountHandler) Login(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	xRequestID := requestid.GetRequestID(c.Request())
	ch.ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, xRequestID)
	log.Info(ch.ctx, "========= START REQUEST : "+requesturi)
	log.Info(ch.ctx, c.Request().Method)
	log.Info(ch.ctx, c.Request().Header)
	form := entity.LoginUserForm{}
	var err error
	if err == nil {
		err = c.Bind(&form)
	}
	if err == nil {
		err = ch.validate(form)
	}
	if err != nil {
		return ch.errorResponse(c, http.StatusBadRequest, err)
	}
	user, err := config.GetDataStore().User.Auth(form.Email, form.Password)
	if err != nil {
		return ch.errorResponse(c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, user, marshalIndent)
}
