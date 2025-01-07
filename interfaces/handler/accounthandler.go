package handler

import (
	"net/http"
	"strconv"

	"github.com/howood/moggiecollector/domain/dto"
	"github.com/howood/moggiecollector/domain/entity"
	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/howood/moggiecollector/interfaces/handler/request"
	"github.com/howood/moggiecollector/interfaces/handler/response"
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
	return c.JSONPretty(http.StatusOK, ch.responseUsers(users), marshalIndent)
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
	return c.JSONPretty(http.StatusOK, ch.responseUser(user), marshalIndent)
}

// CreateUser is get all users
func (ch AccountHandler) CreateUser(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := ch.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	form := request.CreateUserForm{}
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
	err = ch.UcCluster.AccountUC.CreateUser(ctx, ch.convertToUserDto(form))
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
	form := request.CreateUserForm{}
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
	err = ch.UcCluster.AccountUC.UpdateUser(ctx, userid, ch.convertToUserDto(form))
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
	form := request.LoginUserForm{}
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
	user, err := ch.UcCluster.AccountUC.AuthUser(ctx, ch.convertToLoginrDto(form))
	if err != nil {
		return ch.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	if err == nil {
		token = ch.createToken(ctx, user.UserID, user.Email)
	}
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"token": token}, marshalIndent)
}

func (ch AccountHandler) convertToUserDto(user request.CreateUserForm) *dto.UserDto {
	return &dto.UserDto{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (ch AccountHandler) convertToLoginrDto(user request.LoginUserForm) *dto.LoginDto {
	return &dto.LoginDto{
		Email:    user.Email,
		Password: user.Password,
	}
}

func (ch AccountHandler) responseUser(user *entity.User) response.UserResponse {
	return response.UserResponse{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}
}

func (ch AccountHandler) responseUsers(users []*entity.User) []response.UserResponse {
	resUsers := make([]response.UserResponse, 0)

	for _, user := range users {
		resUsers = append(resUsers, ch.responseUser(user))
	}

	return resUsers
}
