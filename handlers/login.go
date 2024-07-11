package handlers

import (
	"goilerplate-api/util"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
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
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "username or password can not be empty",
		})
	}

	hashed, _ := util.HashPassword(password)

	if util.CheckPasswordHash(password, hashed) == false {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "username or password is incorrect",
		})
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
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		color.Red("Could not sign JWT token")
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token": signedToken,
	})
}
