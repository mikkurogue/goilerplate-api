package handlers

import (
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

	// first check if nothing is empty
	if username == "" || password == "" {
		return echo.ErrBadRequest
	}

	hashed, _ := HashPassword(password)

	if CheckPasswordHash(password, hashed) == false {
		// todo probably expand this, but not sure when or how or with what
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
