package handlers

import (
	"fmt"
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

type PostHelloWorldDTO struct {
	Message string `json:"message"`
}

func PostHelloWorld(c echo.Context) error {
	json := PostHelloWorldDTO{}

	err := c.Bind(&json)
	if err != nil {
		return err
	}

	fmt.Printf("Requestbody \n%v", json)

	return c.JSON(http.StatusCreated, json)
}

func HelloProtectedWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello protected world!")
}

func PostHelloWorldMultipart(c echo.Context) error {

	msg := c.FormValue("message")

	if len(msg) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]any {
			"error": "400",
			"error_message": "Missing field \"message\" from form data",
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"message": msg,
	})
}
