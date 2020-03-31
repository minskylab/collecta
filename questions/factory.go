package questions

import (
	"strconv"

	"github.com/minskylab/collecta"
	"github.com/minskylab/collecta/inputs/options"
	"github.com/minskylab/collecta/inputs/satisfaction"
	"github.com/minskylab/collecta/inputs/text"
	"github.com/minskylab/collecta/inputs/yesno"
	"github.com/pkg/errors"
)

type QuestionFactory struct {
	Order             int
	Title             string
	Anonymous         bool
	Type              collecta.InputType
	Default           *string
	Defaults          *[]string
	Options           *[]string
	Expected          *string
	MaxChars          *int
	MultipleSelection *bool
}

func NewQuestionFromFactory(factory QuestionFactory) (*collecta.Question, error) {
	var input collecta.Input
	var err error

	switch factory.Type {
	case collecta.Text:
		mods := make([]text.Mod, 0)

		if factory.Defaults != nil {
			if len(*factory.Defaults) > 0 {
				def := (*factory.Defaults)[0]
				mods = append(mods, text.WithDefaultValue(def))
			}
		}

		if factory.Default != nil {
			mods = append(mods, text.WithDefaultValue(*factory.Default))
		}

		if factory.Expected != nil {
			mods = append(mods, text.WithExpected(*factory.Expected))
		}

		if factory.MaxChars != nil {
			mods = append(mods, text.WithMaxChars(*factory.MaxChars))
		}

		input, err = text.New(mods...)
		if err != nil {
			return nil, errors.Wrap(err, "error at create new input")
		}
	case collecta.Satisfaction:
		mods := make([]satisfaction.Mod, 0)
		if factory.Defaults != nil {
			if len(*factory.Defaults) > 0 {
				def := (*factory.Defaults)[0]
				defaultValue, err := strconv.ParseFloat(def, 64)
				if err != nil {
					return nil, errors.Wrap(err, "default float value parse failed")
				}
				mods = append(mods, satisfaction.WithDefaultValue(defaultValue))
			}
		}

		if factory.Default != nil {
			defaultValue, err := strconv.ParseFloat(*factory.Default, 64)
			if err != nil {
				return nil, errors.Wrap(err, "default float value parse failed")
			}
			mods = append(mods, satisfaction.WithDefaultValue(defaultValue))
		}

		if factory.Options != nil {
			opts := map[int]satisfaction.FeelingOption{}
			for i, opt := range *factory.Options {
				opts[i] = satisfaction.FeelingOption(opt)
			}
			satisfaction.WithCustomFeelingMap(opts)
		}

		input, err = satisfaction.New(mods...)
		if err != nil {
			return nil, errors.Wrap(err, "error at create new satisfaction input")
		}

	case collecta.Options:
		mods := make([]options.Mod, 0)

		if factory.Defaults != nil {
			defaults := make([]int, 0)
			for _, def := range *factory.Defaults {
				d, err := strconv.Atoi(def)
				if err != nil {
					return nil, errors.Wrap(err, "error at parse string to int for default index")
				}
				defaults = append(defaults, d)
			}
			mods = append(mods, options.WithDefaults(defaults))
		}

		if factory.Default != nil {
			d, err := strconv.Atoi(*factory.Default)
			if err != nil {
				return nil, errors.Wrap(err, "error at parse string to int for default index")
			}
			mods = append(mods, options.WithDefaults([]int{d}))
		}

		if factory.MultipleSelection != nil {
			mods = append(mods, options.WithMultipleSelection(*factory.MultipleSelection))
		}

		if factory.Options == nil {
			return nil, errors.New("options is null, but for create new option question you need declare that")
		}

		input, err = options.New(*factory.Options, mods...)
		if err != nil {
			return nil, errors.Wrap(err, "error at create new options input")
		}

	case collecta.YesNo:
		mods := make([]yesno.Mod, 0)
		if factory.Defaults != nil {
			if len(*factory.Defaults) > 0 {
				def := (*factory.Defaults)[0]
				defaultValue, err := strconv.ParseBool(def)
				if err != nil {
					return nil, errors.Wrap(err, "error at parse boolean default")
				}
				mods = append(mods, yesno.WithDefaultValue(defaultValue))
			}
		}

		if factory.Default != nil {
			defaultValue, err := strconv.ParseBool(*factory.Default)
			if err != nil {
				return nil, errors.Wrap(err, "error at parse boolean default")
			}
			mods = append(mods, yesno.WithDefaultValue(defaultValue))
		}

		input, err = yesno.New(mods...)
		if err != nil {
			return nil, errors.New("error at create new yesNo input type")
		}

	default:
		return nil, errors.New("unsupported question type, select one of: text, satisfaction, options, yesNo")
	}

	q := &collecta.Question{
		Order:     factory.Order,
		Title:     factory.Title,
		Anonymous: factory.Anonymous,
		Input:     input,
	}

	q.CalculateID()

	return q, nil
}
