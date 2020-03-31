package parser

import (
	"github.com/minskylab/collecta"
	"github.com/minskylab/collecta/questions"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type inputYaml struct {
	Type              string    `yaml:"type"`
	Default           *string   `yaml:"default"`
	Defaults          *[]string `yaml:"defaults"`
	Options           *[]string `yaml:"options"`
	Expected          *string   `yaml:"expected"`
	MaxChars          *int      `yaml:"maxChars"`
	MultipleSelection *bool     `yaml:"multiple"`
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
	Title      string                 `yaml:"title"`
	Spans      []spanYaml             `yaml:"span"`
	Questions  []questionYaml         `yaml:"questions"`
}

func fromYAML(yamlData []byte) (*collecta.Survey, error) {
	survey := new(surveyYaml)
	if err := yaml.Unmarshal(yamlData, survey); err != nil {
		return nil, errors.Wrap(err, "error at try to decode yaml data to yamlSurvey")
	}

	collectaSurvey := new(collecta.Survey)

	collectaSurvey.Metadata = survey.Metadata
	collectaSurvey.Title = survey.Title

	collectaSurvey.Spans = make([]collecta.Span, 0)
	for _, span := range survey.Spans {
		collectaSurvey.Spans = append(collectaSurvey.Spans, collecta.Span{
			After:  span.After,
			Before: span.Before,
			Text:   span.Text,
		})
	}

	collectaSurvey.Questions = make([]collecta.Question, 0)
	for i, question := range survey.Questions {
		q, err := questions.NewQuestionFromFactory(questions.QuestionFactory{
			Order:             i,
			Title:             question.Title,
			Anonymous:         question.Anonymous,
			Type:              collecta.InputType(question.Input.Type),
			Default:           question.Input.Default,
			Defaults:          question.Input.Defaults,
			Options:           question.Input.Options,
			Expected:          question.Input.Expected,
			MaxChars:          question.Input.MaxChars,
			MultipleSelection: question.Input.MultipleSelection,
		})
		if err != nil {
			return nil, errors.Wrap(err, "error at generate new question with collecta factory question")
		}

		collectaSurvey.Questions = append(collectaSurvey.Questions, *q)
	}

	return collectaSurvey, nil
}
