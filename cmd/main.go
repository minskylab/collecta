package main

import (
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	httpEngine := gin.New()

	httpEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://core.collecta.site", "https://utec.collecta.site"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Collecta-Token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	collectaDB, err := db.NewDB(context.Background())
	if err != nil {
		panic(errors.Cause(err))
	}

	collectaAuth, err := auth.New(httpEngine, collectaDB)
	if err != nil {
		panic(errors.Cause(err))
	}

	collectaAuth.RegisterCallbacks("https://core.collecta.site")

	collectaAPI, err := api.New(httpEngine, collectaDB, collectaAuth)
	if err != nil {
		panic(errors.Cause(err))
	}

	collectaAPI.RegisterGraphQLHandlers()

	if err = httpEngine.Run(":8080"); err != nil {
		panic(errors.Cause(err))
	}
}
