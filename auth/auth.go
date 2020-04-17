package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/markbates/goth"
)

type CollectaAuth struct {
	mainRouter      *gin.Engine
	callbackMatcher func(user goth.User) (string, error)
}

func New(engine *gin.Engine) (*CollectaAuth, error) {
	return &CollectaAuth{mainRouter: engine}, nil
}

func (collectaAuth *CollectaAuth) RegisterCallback(host string) {
	callback := host + "/auth/google/callback"
	if strings.HasSuffix(host, "/") {
		callback = host[:len(host)-1] + "/auth/google/callback"
	}
	log.Info("callback: ", callback)
	collectaAuth.startGoogleAuth(callback)
}
