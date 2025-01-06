package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/howood/moggiecollector/application/actor"
	"github.com/howood/moggiecollector/di/dbcluster"
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
	baseHandler := handler.BaseHandler{UcCluster: uccluster}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/users", handler.AccountHandler{BaseHandler: baseHandler}.GetUsers)
	e.GET("/users/:id", handler.AccountHandler{BaseHandler: baseHandler}.GetUser)
	e.POST("/users", handler.AccountHandler{BaseHandler: baseHandler}.CreateUser)
	e.PUT("/users/:id", handler.AccountHandler{BaseHandler: baseHandler}.UpdateUser)
	e.DELETE("/users/:id", handler.AccountHandler{BaseHandler: baseHandler}.InActiveUser)

	e.POST("/login", handler.AccountHandler{BaseHandler: baseHandler}.Login)

	jwtconfig := echojwt.Config{
		Skipper: custommiddleware.OptionsMethodSkipper,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entity.JwtClaims)
		},
		SigningKey: []byte(actor.TokenSecret),
		ContextKey: actor.JWTContextKey,
	}
	e.GET("/profile", handler.ClientHandler{}.GetProfile, echojwt.WithConfig(jwtconfig))

	e.Logger.Fatal(e.Start(":" + defaultPort))

}
