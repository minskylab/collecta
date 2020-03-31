package parser

import (
	"io/ioutil"
	"strings"

	"github.com/minskylab/collecta"
	"github.com/pkg/errors"
)

func ParseFile(filepath string) (*collecta.Survey, error) {
	if strings.HasSuffix(filepath, ".yaml") || strings.HasSuffix(filepath, ".yml") {
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			return nil, errors.Wrap(err, "error at read yaml file")
		}
		return fromYAML(data)
	}
	return nil, nil
}
