package parser

import (
	"github.com/minskylab/collecta"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type inputYaml struct {
	Type     string    `yaml:"type"`
	Default  *string   `yaml:"default"`
	Defaults *[]string `yaml:"defaults"`
	Options  *[]string `yaml:"options"`
}

type questionYaml struct {
	ID        string    `yaml:"id"`
	Title     string    `yaml:"title"`
	Anonymous bool      `yaml:"anonymous"`
	Input     inputYaml `yaml:"input"`
}

type spanYaml struct {
	After  *string `yaml:"after"`
	Before *string `yaml:"before"`
	Text   string  `yaml:"text"`
}
type surveyYaml struct {
	APIVersion string                 `yaml:"apiVersion"`
	Metadata   map[string]interface{} `yaml:"metadata"`
	Span       []spanYaml             `yaml:"span"`
	Questions  []questionYaml         `yaml:"questions"`
}

func fromYAML(yamlData []byte) (*collecta.Survey, error) {
	survey := new(surveyYaml)
	if err := yaml.Unmarshal(yamlData, survey); err != nil {
		return nil, errors.Wrap(err, "error at try to decode yaml data to yamlSurvey")
	}
	return nil, nil
}
