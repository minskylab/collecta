package main

import (
	"os"
	"strings"

	"github.com/minskylab/collecta/config"
	"github.com/spf13/viper"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return port
}


func getHost() string {
	host := viper.GetString(config.DomainHost)
	if host == "" {
		host = "https://core.collecta.site"
	}

	return host
}