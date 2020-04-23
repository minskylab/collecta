package auth

import (
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/minskylab/collecta/config"
	"github.com/minskylab/collecta/db"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Auth struct {
	mainRouter *gin.Engine
	db         *db.DB
	jwtSecret []byte
	jwtDuration time.Duration
}

func New(engine *gin.Engine, db *db.DB) (*Auth, error) {
	secret := viper.GetString(config.AuthJWTSecret)
	if secret != "" {
		return nil, errors.New("invalid secret for jwt, please set a valid one")
	}

	exp := viper.GetString(config.AuthJWTExpiration)
	dur, err := time.ParseDuration(exp)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to parse the duration, set correctly in the format (e.g. '24h')")
	}
	a := Auth{
		mainRouter:  engine,
		db:          db,
		jwtSecret:   []byte(secret),
		jwtDuration: dur,
	}

	return &a, nil
}

func (collectaAuth *Auth) RegisterCallbacks(host string) {
	callback := host + "/auth/google/callback"
	if strings.HasSuffix(host, "/") {
		callback = host[:len(host)-1] + "/auth/google/callback"
	}

	log.Info("google callback: ", callback)
	collectaAuth.startGoogleAuth(callback)
}
