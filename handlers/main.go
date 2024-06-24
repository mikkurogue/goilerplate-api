package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

func HelloWorldJson(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Hello, world",
		"date":    time.Now(),
	})
}
