// starting on writing some shit for a auditter to track actions done in the database
// we need to make sure we can easily do this by making an abstraction of an Audit structure
// and then making this making sure we fire the Auditor (whatever this is in the future, for now its TBD)
// and move the spawning of each auditor type to a different go channel
// ideally we keep the auditor away from the api itself but not abstracted to a different app as we want
// instant intergration, just no thread blocking.

package util

type Audit struct {
	Id     string `json: id`
	Action string `json: action`
}

func (a *Audit) LogAction() (*Audit, error) {
	a.Id = "some id"
	a.Action = "some action"

	return a, nil
}
