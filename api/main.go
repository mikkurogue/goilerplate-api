package api

import (
	"goilerplate-api/handlers"

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
}
