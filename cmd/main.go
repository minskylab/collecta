package main

import (
	"context"

	"github.com/minskylab/collecta/api"
	"github.com/minskylab/collecta/auth"
	"github.com/minskylab/collecta/config"
	"github.com/minskylab/collecta/db"
	"github.com/pkg/errors"
)

func main() {
	if err := config.ReadConfig(); err != nil {
		panic(errors.Cause(err))
	}

	httpEngine, err := getHTTPEngine()
	if err != nil {
		panic(errors.Cause(err))
	}

	collectaDB, err := db.NewDB(context.Background())
	if err != nil {
		panic(errors.Cause(err))
	}

	collectaAuth, err := auth.New(httpEngine, collectaDB)
	if err != nil {
		panic(errors.Cause(err))
	}

	host := getHost()
	collectaAuth.RegisterCallbacks(host)

	collectaAPI, err := api.New(httpEngine, collectaDB, collectaAuth)
	if err != nil {
		panic(errors.Cause(err))
	}

	collectaAPI.RegisterGraphQLHandlers(true)

	port := getPort()
	if err = httpEngine.Run(port); err != nil {
		panic(errors.Cause(err))
	}
}
