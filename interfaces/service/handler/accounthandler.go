package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/howood/moggiecollector/domain/entity"
	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/howood/moggiecollector/infrastructure/requestid"
	"github.com/howood/moggiecollector/interfaces/service/usecase"
	"github.com/labstack/echo/v4"
)

// AccountHandler struct
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
	users, err := usecase.AccountUsecase{}.GetUsers(withinactive)
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
	user, err := usecase.AccountUsecase{}.GetUser(userid)
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
	err = usecase.AccountUsecase{}.CreateUser(form)
	if err != nil {
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
	err = usecase.AccountUsecase{}.UpdateUser(userid, form)
	if err != nil {
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
	err := usecase.AccountUsecase{}.InActiveUser(userid)
	if err != nil {
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
	var token string
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
	user, err := usecase.AccountUsecase{}.AuthUser(form)
	if err != nil {
		return ch.errorResponse(c, http.StatusBadRequest, err)
	}
	if err == nil {
		token, err = ch.createToken(user.UserID, user.Email)
	}
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"token": token}, marshalIndent)
}
