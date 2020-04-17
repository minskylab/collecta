package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/minskylab/collecta"
	"github.com/minskylab/collecta/api"
	"github.com/minskylab/collecta/auth"
	"github.com/minskylab/collecta/config"
	"github.com/pkg/errors"
)

func main() {
	if err := config.ReadConfig(); err != nil {
		panic(errors.Cause(err))
	}

	httpEngine := gin.New()


	db, err := collecta.NewDB(context.Background())
	if err != nil {
		panic(errors.Cause(err))
	}

	collectaAuth, err := auth.New(httpEngine, db)
	if err != nil {
		panic(errors.Cause(err))
	}

	collectaAuth.RegisterCallback("https://core.collecta.site")

	api.RegisterGraphQLHandlers(httpEngine, db)

	if err = httpEngine.Run(":8080"); err != nil {
		panic(errors.Cause(err))
	}
}
