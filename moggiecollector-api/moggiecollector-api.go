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
	uccluster := uccluster.NewUsecaseCluster(dataStore)
	sccluster := svcluster.NewServiceCluster(dataStore)
	baseHandler := handler.BaseHandler{UcCluster: uccluster, SvCluster: sccluster}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	v1AdminAPI := e.Group("/admin/v1", custommiddleware.RequestLog(sccluster))

	v1AdminAPI.GET("/users", handler.AccountHandler{BaseHandler: baseHandler}.GetUsers)
	v1AdminAPI.GET("/users/:id", handler.AccountHandler{BaseHandler: baseHandler}.GetUser)
	v1AdminAPI.POST("/users", handler.AccountHandler{BaseHandler: baseHandler}.CreateUser)
	v1AdminAPI.PUT("/users/:id", handler.AccountHandler{BaseHandler: baseHandler}.UpdateUser)
	v1AdminAPI.DELETE("/users/:id", handler.AccountHandler{BaseHandler: baseHandler}.InActiveUser)

	jwtconfig := echojwt.Config{
		Skipper: custommiddleware.OptionsMethodSkipper,
		NewClaimsFunc: func(_ echo.Context) jwt.Claims {
			return new(entity.JwtClaims)
		},
		SigningKey: []byte(actor.TokenSecret),
		ContextKey: actor.JWTContextKey,
	}
	v1API := e.Group("/api/v1", custommiddleware.RequestLog(sccluster))
	v1API.POST("/login", handler.AccountHandler{BaseHandler: baseHandler}.Login)
	v1API.GET("/profile", handler.ClientHandler{}.GetProfile, echojwt.WithConfig(jwtconfig))

	e.Logger.Fatal(e.Start(":" + defaultPort))
}
