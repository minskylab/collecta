package text

import (
	"regexp"

	"github.com/minskylab/collecta"
)

// Text represent an Text Runtime type of any question
type Text struct {
	value        string
	defaultValue string
	maxChars     int
	expected     *regexp.Regexp
}

// New creates a new Text modified with options
func New(mods ...Mod) (*Text, error) {
	text := new(Text)
	text.maxChars = 120 // default
	for _, option := range mods {
		if err := option(text); err != nil {
			return nil, err
		}
	}
	return text, nil
}

// Type returns the type of text, in this case: "text"
func (text *Text) Type() collecta.InputType {
	return collecta.Text
}

func (text *Text) Value() ([]string, error) {
	return []string{text.value}, nil
}
