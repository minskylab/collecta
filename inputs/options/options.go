package options

import (
	"github.com/minskylab/collecta"
	"github.com/pkg/errors"
)

// Options represent a Option Runtime type of any question
type Options struct {
	value          []int
	defaultValue   []int
	multipleSelect bool
	options        []Option
}

// New creates a new Option modified with options
func New(options []string, mods ...Mod) (*Options, error) {
	optionsInput := new(Options)
	optionsInput.options = make([]Option, 0)
	for _, opt := range options {
		optionsInput.options = append(optionsInput.options, newOptionFromValue(opt))
	}
	return optionsInput, nil
}

// Type returns the type of text, in this case: "text"
func (options *Options) Type() collecta.InputType {
	return collecta.Options
}

func (options *Options) Value() ([]string, error) {
	values := make([]string, 0)
	for _, val := range options.value {
		if val < 0 || val > len(options.options)-1 {
			return nil, errors.New("invalid value, index out of options size")
		}
		values = append(values, options.options[val].value)
	}
	return values, nil
}
