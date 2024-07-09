package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

type DbConfig struct {
	DatabaseName string
	AuthToken    string
}

func Init(databaseName, authToken string) (DbConfig, error) {
	color.Yellow("initialising database connection configuration...")

	if len(databaseName) == 0 || len(authToken) == 0 {
		return DbConfig{}, errors.New("no database name or auth token provided")
	}

	color.Green("successfully configured the connection!")
	return DbConfig{
		DatabaseName: databaseName,
		AuthToken:    authToken,
	}, nil
}

func (database DbConfig) CreateConnection() *sql.DB {
	color.Yellow("starting connection...")

	// Change this URL to be your database provider.
	// this example uses Turso, because its a quick setup
	// but this should work with most sql database providers
	url := fmt.Sprintf("libsql://%s.turso.io=authToken=%s",
		database.DatabaseName,
		database.AuthToken)

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}

	color.Green("successfully connected to database: " + database.DatabaseName)

	// check when to close database connection
	// defer db.Close()
	return db
}
