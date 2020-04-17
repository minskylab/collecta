package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/minskylab/collecta"
	"github.com/minskylab/collecta/api/graph"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/auth"
)

// RegisterGraphQLHandlers register the playground and graphql query endpoint to my main gin engine
func RegisterGraphQLHandlers(r *gin.Engine, db *collecta.DB, auth *auth.CollectaAuth) {
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{
				DB: db,
				Auth: auth,
			}},
		),
	)

	query := func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}

	play := func(c *gin.Context) {
		playground.Handler("Colecta playground", "/graphql").ServeHTTP(c.Writer, c.Request)
	}

	r.GET("/", play)
	r.POST("/graphql", query)
}
