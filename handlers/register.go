package handlers

import (
	"net/http"
	"os"

	"goilerplate-api/db"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	hashed, _ := HashPassword(password)

	dbName := os.Getenv("DATABASE_NAME")
	authToken := os.Getenv("AUTH_TOKEN")

	dbConfig, err := db.Init(dbName, authToken)
	if err != nil {
		color.Red("cannot initialise database config")
	}

	database := dbConfig.CreateConnection()

	// todo check if this is the actual way to fire insert into database...
	result, err := database.Exec(`INSERT INTO users VALUES (%s, %s)`, username, hashed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "something went wrong",
		})
	}

	id, err := result.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error occured while creating account",
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message":   "account registered",
		"insert_id": id,
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
