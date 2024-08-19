// starting on writing some shit for a auditter to track actions done in the database
// we need to make sure we can easily do this by making an abstraction of an Audit structure
// and then making this making sure we fire the Auditor (whatever this is in the future, for now its TBD)
// and move the spawning of each auditor type to a different go channel
// ideally we keep the auditor away from the api itself but not abstracted to a different app as we want
// instant intergration, just no thread blocking.

package util

import (
	"errors"
	"fmt"
	"goilerplate-api/db"
	"os"

	"github.com/google/uuid"
)

type Audit struct {
	Id     string `json: id`
	Action string `json: action`
	Table  string `json: table`
}

func (a *Audit) CreateAction(table, action string) (*Audit, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return &Audit{}, errors.New("some unknown error occured...")
	}

	a.Id = id.String()
	a.Action = action
	a.Table = table

	return a, nil
}

func (a *Audit) ReadAction(id, table string) (*Audit, error) {
	// some db logic to fetch the id and assign it here. also error check it
	// TODO: NYI
	const action = "Some action name"

	return &Audit{
		Id:     id,
		Action: action,
		Table:  table,
	}, nil
}

func (a *Audit) DeleteAction() {
	dbName := os.Getenv("DATABASE_NAME")
	authToken := os.Getenv("AUTH_TOKEN")

	// delete from database, check if error then return error.
	db, err := db.Init(dbName, authToken)
	if err != nil {
		fmt.Println(err.Error())
	}

	db.CreateConnection()

	// TODO: create the rest of the stuff here
}
