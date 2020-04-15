package config

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func readConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/collecta/")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Info("config not loaded correctly")
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SetEnvPrefix("collecta")
			viper.AutomaticEnv()
		} else {
			return errors.Wrap(err, "error at read config with viper")
		}
	}

	log.Info(viper.GetString("google.clientID")) // env: COLLECTA_GOOGLE_CLIENTID

	return nil
}

func ReadConfig() error {
	return readConfig()
}
