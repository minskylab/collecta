package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/minskylab/collecta/api/graph"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/auth"
	"github.com/minskylab/collecta/db"
	log "github.com/sirupsen/logrus"
)

type API struct {
	engine *gin.Engine
	db     *db.DB
	auth   *auth.Auth
}

func New(r *gin.Engine, db *db.DB, auth *auth.Auth) (*API, error) {
	return &API{
		engine: r,
		db:     db,
		auth:   auth,
	}, nil
}

// RegisterGraphQLHandlers register the playground and graphql query endpoint to my main gin engine
func (api *API) RegisterGraphQLHandlers(withPlayground bool) {
	config := generated.Config{Resolvers: &graph.Resolver{
		DB:   api.db,
		Auth: api.auth,
	}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	srv.Use(extension.FixedComplexityLimit(500))

	query := func(c *gin.Context) {
		log.Warn("graphql request intercept | content-type: ", c.ContentType())
		srv.ServeHTTP(c.Writer, c.Request)
	}

	if withPlayground {
		play := func(c *gin.Context) {
			playground.Handler("GraphQL Playground | Collecta", "/graphql").ServeHTTP(c.Writer, c.Request)
		}
		api.engine.GET("/", play)
	}

	api.engine.Use(api.auth.Middleware())
	api.engine.POST("/graphql", query)
}
