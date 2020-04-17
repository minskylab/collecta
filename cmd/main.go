package main

import (
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	httpEngine.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:3000", "https://core.collecta.site", "https://utec.collecta.site"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Redirect"},
		AllowAllOrigins:  true,
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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

	_, err = db.Ent.Debug().Domain.Create().
		SetID(uuid.New()).
		SetDomain("utec.collecta.site").
		SetName("UTEC").
		SetTags([]string{"utec", "academic", "university"}).
		SetEmail("contacto@utec.edu.pe").
		SetCollectaDomain("https://utec.collecta.site").
		Save(context.Background())
	if err != nil {
		panic(errors.Cause(err))
	}

	if err = httpEngine.Run(":8080"); err != nil {
		panic(errors.Cause(err))
	}
}
