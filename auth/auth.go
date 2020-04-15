package auth

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/spf13/viper"
)

type CollectaAuth struct {
	mainRouter *gin.Engine
}

func New(engine *gin.Engine) (*CollectaAuth, error) {
	return &CollectaAuth{mainRouter: engine}, nil
}

func (collectaAuth *CollectaAuth) start(callbackURL string) {
	clientID := viper.GetString("google.clientID")
	secretKey := viper.GetString("google.secret")

	goth.UseProviders(
		google.New(clientID, secretKey, callbackURL),
	)

	r := collectaAuth.mainRouter.Group("/auth")

	r.GET("/google", func(c *gin.Context) {
		modQuery := c.Request.URL.Query()
		modQuery.Add("provider", "google")
		c.Request.URL.RawQuery = modQuery.Encode()
		if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
			log.Info(gothUser)
			spew.Fdump(c.Writer, gothUser)
		} else {
			log.WithError(err).Info("trying new google auth")
			gothic.BeginAuthHandler(c.Writer, c.Request)
		}
	})

	r.GET("/google/callback", func(c *gin.Context) {
		modQuery := c.Request.URL.Query()
		modQuery.Add("provider", "google")
		c.Request.URL.RawQuery = modQuery.Encode()
		user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			log.WithError(err).Info("fail at complete google user auth")
			return
		}
		log.Info(user)
		spew.Fdump(c.Writer, user)
	})

}
func (collectaAuth *CollectaAuth) Start(host string) {
	callback := host + "/auth/google/callback"
	if strings.HasSuffix(host, "/") {
		callback = host[:len(host)-1] + "/auth/google/callback"
	}
	log.Info("callback: ", callback)
	collectaAuth.start(callback)
}
