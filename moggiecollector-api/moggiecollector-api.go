package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/howood/moggiecollector/application/actor"
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/infrastructure/custommiddleware"
	"github.com/howood/moggiecollector/interfaces/service/handler"
	"github.com/howood/moggiecollector/library/utils"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// DefaultPort is default port of server
var DefaultPort = utils.GetOsEnv("SERVER_PORT", "8080")

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/users", handler.AccountHandler{}.GetUsers)
	e.GET("/users/:id", handler.AccountHandler{}.GetUser)
	e.POST("/users", handler.AccountHandler{}.CreateUser)
	e.PUT("/users/:id", handler.AccountHandler{}.UpdateUser)
	e.DELETE("/users/:id", handler.AccountHandler{}.InActiveUser)

	e.POST("/login", handler.AccountHandler{}.Login)

	jwtconfig := echojwt.Config{
		Skipper: custommiddleware.OptionsMethodSkipper,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entity.JwtClaims)
		},
		SigningKey: []byte(actor.TokenSecret),
		ContextKey: actor.JWTContextKey,
	}
	e.GET("/profile", handler.ClientHandler{}.GetProfile, echojwt.WithConfig(jwtconfig))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", DefaultPort)))

}
