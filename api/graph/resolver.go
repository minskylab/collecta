package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/minskylab/collecta"
	"github.com/minskylab/collecta/auth"
)

type Resolver struct {
	Auth *auth.CollectaAuth
	DB *collecta.DB
}
