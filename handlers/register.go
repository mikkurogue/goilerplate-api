package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"goilerplate-api/db"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {

	err := godotenv.Load()

	username := c.FormValue("username")
	password := c.FormValue("password")

	if len(username) == 0 || len(password) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "username or password empty",
		})
	}

	hashed, _ := HashPassword(password)

	databaseName := os.Getenv("DATABASE_NAME")
	authToken := os.Getenv("AUTH_TOKEN")

	dbConfig, err := db.Init(databaseName, authToken)
	if err != nil {
		log.Fatal(err)
		color.Red("cannot initialise database config")
	}

	database := dbConfig.CreateConnection()

	result, err := database.Exec("INSERT INTO users VALUES (?, ?, ?, ?, ?)",
		uuid.New().String(),
		username,
		hashed,
		time.Now().Format(time.DateTime),
		time.Now().Format(time.DateTime))
	if err != nil {
		log.Fatalln(err)
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
