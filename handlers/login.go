package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JWTClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// throw unauth error for now just hard coded
	if username == "typescript" || password == "isbad" {
		return echo.ErrUnauthorized
	}

	if username == "" || password == "" {
		fmt.Println("username or password is empty")
		return echo.ErrBadRequest
	}

	// set custom claims
	claims := &JWTClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Make sure to set this "secret" to an .env variable in your builds
	signedToken, err := token.SignedString([]byte("secret"))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token": signedToken,
	})
}
