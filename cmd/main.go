package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minskylab/collecta"
	"github.com/minskylab/collecta/api"
	"github.com/minskylab/collecta/auth"
	"github.com/minskylab/collecta/config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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

	utec, err := db.Ent.Debug().Domain.Create().
		SetID(uuid.New()).
		SetDomain("utec.collecta.site").
		SetName("UTEC").
		SetTags([]string{"utec", "academic", "university"}).
		SetEmail("contacto@utec.edu.pe").
		SetCollectaDomain("https://utec.collecta.site").
		Save(context.Background())
	if err != nil {
		log.Error(errors.Cause(err))
	}

	// user, err := db.Ent.Debug().User.Create().
	// 	SetID(uuid.New()).
	// 	SetName("Bregy Malpartida").
	// 	SetUsername("bregy").
	// 	SetDomain(utec).
	// 	Save(context.Background())
	// if err != nil {
	// 	panic(errors.Cause(err))
	// }

	spew.Dump(utec)

	if err = httpEngine.Run(":8080"); err != nil {
		panic(errors.Cause(err))
	}
}
