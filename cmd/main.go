package main

import (
	"github.com/gin-gonic/gin"
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

	collectaAuth.Start("https://core.collecta.site")

	if err = httpEngine.Run(":8081"); err != nil {
		panic(err)
	}
}
