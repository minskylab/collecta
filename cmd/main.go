package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minskylab/collecta"
	"github.com/minskylab/collecta/auth"
	"github.com/minskylab/collecta/config"
	"github.com/pkg/errors"
)

func main() {
	if err := config.ReadConfig(); err != nil {
		panic(errors.Cause(err))
	}

	httpEngine := gin.New()

	collectaAuth, err := auth.New(httpEngine)
	if err != nil {
		panic(errors.Cause(err))
	}

	collectaAuth.RegisterCallback("https://core.collecta.site")

	db, err := collecta.NewDB(context.Background())
	if err != nil {
		panic(errors.Cause(err))
	}

	utec, err := db.Ent.Debug().Domain.Create().
		SetID(uuid.New()).
		SetDomain("utec.collecta.site").
		SetName("UTEC").
		SetTags([]string{"utec", "academic", "university"}).
		SetEmail("contacto@utec.edu.pe").
		Save(context.Background())
	if err != nil {
		panic(errors.Cause(err))
	}

	user, err := db.Ent.Debug().User.Create().
		SetID(uuid.New()).
		SetName("Bregy Malpartida").
		SetUsername("bregy").
		SetDomain(utec).
		Save(context.Background())
	if err != nil {
		panic(errors.Cause(err))
	}

	spew.Dump(user)

	if err = httpEngine.Run(":8080"); err != nil {
		panic(errors.Cause(err))
	}
}
