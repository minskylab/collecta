package collecta

import (
	"github.com/minskylab/collecta/api"
	"github.com/minskylab/collecta/auth"
	"github.com/minskylab/collecta/db"
)

// Collecta wraps all the thing related to data manipulation on collecta, that is for
// create a robust API based on Golang
type Collecta struct {
	db  *db.DB
	api *api.API

	auth *auth.Auth
}
