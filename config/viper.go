package config

import (
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func readConfig() error {
	viper.SetConfigName("collecta")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/collecta/")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Info("config not loaded correctly")
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SetEnvPrefix("collecta")
			viper.AutomaticEnv()
			viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		} else {
			return errors.Wrap(err, "error at read config with viper")
		}
	}

	return nil
}

func ReadConfig() error {
	return readConfig()
}
