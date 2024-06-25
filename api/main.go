package api

import (
	"goilerplate-api/handlers"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	e := echo.New()

	e.Use(middleware.Logger())

	// Default cors cfg
	e.Use(middleware.CORS())

	Routes(e)

	e.Logger.Fatal(e.Start(":1337"))
}

func Routes(e *echo.Echo) {

	e.GET("/hello-world", handlers.HelloWorld)
	e.GET("/hello-world-json", handlers.HelloWorldJson)

	e.POST("/login", handlers.Login)

	protected := e.Group("/protected")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(handlers.JWTClaims)
		},
		// Make sure this is the same env var
		SigningKey: []byte("secret"),
	}

	protected.Use(echojwt.WithConfig(config))

	protected.GET("/hello-protected-world", handlers.HelloProtectedWorld)
}
