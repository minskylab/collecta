package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/minskylab/collecta/config"
	"github.com/spf13/viper"
)

func getHTTPEngine() (*gin.Engine, error) {
	engine := gin.New()

	origins := viper.GetStringSlice(config.CorsOrigins)
	if len(origins) == 0 {
		origins = []string{"http://localhost:3000", "https://core.collecta.site", "https://utec.collecta.site"}
	}

	methods := viper.GetStringSlice(config.CorsMethods)
	if len(methods) == 0 {
		methods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	}

	headers := viper.GetStringSlice(config.CorsHeaders)
	if len(headers) == 0 {
		headers = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	}

	maxAge := viper.GetDuration(config.CorsMaxAge)
	if maxAge == 0 {
		maxAge = 12 * time.Hour
	}


	engine.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     methods,
		AllowHeaders:     headers,
		AllowCredentials: true,
		MaxAge:           maxAge,
	}))

	return engine, nil
}
