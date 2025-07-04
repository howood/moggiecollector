package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/howood/moggiecollector/application/actor"
	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/di/svcluster"
	"github.com/howood/moggiecollector/di/uccluster"
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/infrastructure/custommiddleware"
	"github.com/howood/moggiecollector/interfaces/handler"
	"github.com/howood/moggiecollector/library/utils"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	defaultPort := utils.GetOsEnv("SERVER_PORT", "8080")

	dataStore := dbcluster.NewDatastore()
	svcluster := svcluster.NewServiceCluster(dataStore)
	uccluster := uccluster.NewUsecaseCluster(dataStore, svcluster)
	baseHandler := handler.BaseHandler{UcCluster: uccluster, SvCluster: svcluster}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	v1AdminAPI := e.Group("/admin/v1", custommiddleware.RequestLog(svcluster))

	v1AdminAPI.GET("/users", handler.UserHandler{BaseHandler: baseHandler}.GetUsers)
	v1AdminAPI.GET("/users/:id", handler.UserHandler{BaseHandler: baseHandler}.GetUser)
	v1AdminAPI.POST("/users", handler.UserHandler{BaseHandler: baseHandler}.CreateUser)
	v1AdminAPI.PUT("/users/:id", handler.UserHandler{BaseHandler: baseHandler}.UpdateUser)
	v1AdminAPI.DELETE("/users/:id", handler.UserHandler{BaseHandler: baseHandler}.InActiveUser)
	v1AdminAPI.GET("/users/:id/mfa", handler.UserMfaHandler{BaseHandler: baseHandler}.InitialUserAuthenticator)
	v1AdminAPI.PUT("/users/:id/mfa", handler.UserMfaHandler{BaseHandler: baseHandler}.UpsertUserAuthenticator)

	jwtconfig := echojwt.Config{
		Skipper: custommiddleware.OptionsMethodSkipper,
		NewClaimsFunc: func(_ echo.Context) jwt.Claims {
			return new(entity.JwtClaims)
		},
		SigningKey: []byte(actor.TokenSecret),
		ContextKey: actor.JWTContextKey,
	}
	v1API := e.Group("/api/v1", custommiddleware.RequestLog(svcluster))
	v1API.POST("/login", handler.AuthHandler{BaseHandler: baseHandler}.Login)
	v1API.POST("/login/verify_authenticator", handler.AuthHandler{BaseHandler: baseHandler}.VerifyAuthenticator)
	v1API.GET("/profile", handler.UserHandler{}.GetProfile, echojwt.WithConfig(jwtconfig))

	e.Logger.Fatal(e.Start(":" + defaultPort))
}
