package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	//"os"
	//"time"

	//	"goilerplate-api/db"
	"goilerplate-api/db"
	"goilerplate-api/structs"
	"goilerplate-api/util"

	//	"github.com/fatih/color"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if len(username) == 0 || len(password) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "username or password empty",
		})
	}

	hashed, _ := util.HashPassword(password)

	client := db.InitMongo()
	client.StartSession()
	err := client.Connect(context.TODO())
	if err != nil {
		color.Red("cannot connect to database")
	}

	coll := client.Database("db").Collection("users")

	id, err := uuid.NewV7()
	if err != nil {
		// this should never happen, but if we get here, i want it logged
		// we then should try again anyway because the chances of it happening again is next to 0
		color.Red("user id already exists")
	}

	doc := structs.User{UserId: id.String(), Username: username, Password: hashed, CreatedOn: time.Now().String()}

	res, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Println(err.Error())
		color.Red("cannot insert user into database")
	}

	//	databaseName := os.Getenv("DATABASE_NAME")
	//	authToken := os.Getenv("AUTH_TOKEN")

	//	dbConfig, err := db.Init(databaseName, authToken)
	//	if err != nil {
	//		log.Fatal(err)
	//		color.Red("cannot initialise database config")
	//	}

	//	database := dbConfig.CreateConnection()

	//	result, err := database.Exec("INSERT INTO users VALUES (?, ?, ?, ?, ?)",
	//		uuid.New().String(),
	//		username,
	//		hashed,
	//		time.Now().Format(time.DateTime),
	//		time.Now().Format(time.DateTime))
	//	if err != nil {
	//		log.Fatalln(err)
	//		return c.JSON(http.StatusInternalServerError, echo.Map{
	//			"message": "something went wrong",
	//		})
	//	}

	//	id, err := result.LastInsertId()
	//	if err != nil {
	//		return c.JSON(http.StatusBadRequest, echo.Map{
	//			"message": "error occured while creating account",
	//		})
	//	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message":   "account registered",
		"insert_id": res.InsertedID,
	})
}
