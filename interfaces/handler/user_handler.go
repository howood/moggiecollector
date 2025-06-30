package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/howood/moggiecollector/domain/dto"
	"github.com/howood/moggiecollector/domain/entity"
	log "github.com/howood/moggiecollector/infrastructure/logger"
	"github.com/howood/moggiecollector/interfaces/handler/request"
	"github.com/howood/moggiecollector/interfaces/handler/response"
	"github.com/labstack/echo/v4"
)

// UserHandler struct
type UserHandler struct {
	BaseHandler
}

// GetUsers is get all users
func (uh UserHandler) GetUsers(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := uh.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	withinactive := c.QueryParam("withinactive")
	users, err := uh.UcCluster.UserUC.GetUsers(ctx, withinactive)
	if err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, uh.responseUsers(users), marshalIndent)
}

// GetUser is get all users
func (uh UserHandler) GetUser(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := uh.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)

	userid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	user, err := uh.UcCluster.UserUC.GetUser(ctx, userid)
	if err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, uh.responseUser(user), marshalIndent)
}

// CreateUser is get all users
func (uh UserHandler) CreateUser(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := uh.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	var form request.CreateUserForm
	if err := c.Bind(&form); err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	if err := uh.validate(form); err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	user, err := uh.UcCluster.UserUC.CreateUser(ctx, uh.convertToUserDto(form))
	if err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, uh.responseUser(user), marshalIndent)
}

// CreateUser is get all users
func (uh UserHandler) UpdateUser(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := uh.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)

	userid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}

	form := request.CreateUserForm{}
	if err = c.Bind(&form); err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	if err = uh.validate(form); err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	user, err := uh.UcCluster.UserUC.UpdateUser(ctx, userid, uh.convertToUserDto(form))
	if err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, uh.responseUser(user), marshalIndent)
}

// InActiveUser is get all users
func (uh UserHandler) InActiveUser(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := uh.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)

	userid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}

	if err := uh.UcCluster.UserUC.InActiveUser(ctx, userid); err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"message": "success"}, marshalIndent)
}

// GetProfile is get logined user
func (uh UserHandler) GetProfile(c echo.Context) error {
	requesturi := c.Request().URL.RequestURI()
	ctx := uh.initalGenerateContext(c)
	log.Info(ctx, "========= START REQUEST : "+requesturi)
	log.Info(ctx, c.Request().Method)
	log.Info(ctx, c.Request().Header)
	claims := uh.getClaimsFromToken(c)
	user, err := uh.UcCluster.UserUC.GetUserByToken(ctx, claims)
	if err != nil {
		return uh.errorResponse(ctx, c, http.StatusBadRequest, err)
	}
	return c.JSONPretty(http.StatusOK, uh.responseUser(user), marshalIndent)
}

func (uh UserHandler) convertToUserDto(user request.CreateUserForm) *dto.UserDto {
	return &dto.UserDto{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (uh UserHandler) responseUser(user *entity.User) response.UserResponse {
	return response.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}
}

func (uh UserHandler) responseUsers(users []*entity.User) []response.UserResponse {
	resUsers := make([]response.UserResponse, 0)

	for _, user := range users {
		resUsers = append(resUsers, uh.responseUser(user))
	}

	return resUsers
}
