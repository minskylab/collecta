package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/minskylab/collecta/errors"
)

func main() {
	entSchema := "./ent/schema"
	gqlgenRoot := "./api"

	generateGQLGenFromEntSchema(entSchema, gqlgenRoot)

	if err := os.Chdir(gqlgenRoot); err != nil {
		panic(errors.Wrap(err, "error at change dir to gqlgen root directory"))
	}

	cfg, err := config.LoadConfig("gqlgen.yml")
	if err != nil {
		panic(errors.Wrap(err, "error at try to load gqlgen config file"))
	}

	builder := func (b *modelgen.ModelBuild) *modelgen.ModelBuild {
		for _, model := range b.Models {
			fmt.Println(model.Name)
		}

		return b
	}

	p := modelgen.Plugin{
		MutateHook: builder,
	}

	err = api.Generate(cfg,
		api.AddPlugin(&p),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

}