/*
Bootstrap the database here.
We want to basically fire off all the required queries to create a database.
We need to make a check still to see if we are already have done any bootstrapping
If so, we need to figure out a migration strategy to migrate new tables/columns etc
*/

package bootstrap

import (
	"goilerplate-api/db"
	"os"

	"github.com/fatih/color"
)

func Bootstrap() {

	dbName := os.Getenv("DATABASE_NAME")
	authToken := os.Getenv("AUTH_TOKEN")

	dbConfig, err := db.Init(dbName, authToken)
	if err != nil {
		color.Red("cannot initialise database config")
	}

	database := dbConfig.CreateConnection()

	database.Exec(CREATE_USER_TABLE)

	// close the db connection
	defer database.Close()
}
