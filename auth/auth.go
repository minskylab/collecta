package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/minskylab/collecta"
	log "github.com/sirupsen/logrus"
)

type CollectaAuth struct {
	mainRouter *gin.Engine
	db         *collecta.DB
}

func New(engine *gin.Engine, db *collecta.DB) (*CollectaAuth, error) {
	return &CollectaAuth{mainRouter: engine, db: db}, nil
}

func (collectaAuth *CollectaAuth) RegisterCallback(host string) {
	callback := host + "/auth/google/callback"
	if strings.HasSuffix(host, "/") {
		callback = host[:len(host)-1] + "/auth/google/callback"
	}
	log.Info("callback: ", callback)
	collectaAuth.startGoogleAuth(callback)
}
