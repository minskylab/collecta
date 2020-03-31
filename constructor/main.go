package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/minskylab/collecta/parser"
	"github.com/pkg/errors"
)

func main() {
	filepath := "../drafts/survey.yaml"
	survey, err := parser.ParseFile(filepath)
	if err != nil {
		panic(errors.Cause(err))
	}
	spew.Dump(survey)
}
