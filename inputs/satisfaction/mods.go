package satisfaction

import "github.com/pkg/errors"

type Mod func(satisfaction *Satisfaction) error

func WithDefaultValue(defVal float64) Mod {
	return func(satisfaction *Satisfaction) error {
		satisfaction.defaultValue = defVal
		return nil
	}
}

func WithCustomFeelingOptions(options []FeelingOption) Mod {
	return func(satisfaction *Satisfaction) error {
		if len(options) != 3 && len(options) != 5 {
			return errors.New("invalid number of options, it must be 3 or 5")
		}
		satisfaction.totalOptions = len(options)
		opts := map[int]FeelingOption{}
		for i, opt := range options {
			opts[i] = opt
		}
		satisfaction.options = opts
		return nil
	}
}

func WithCustomFeelingMap(options map[int]FeelingOption) Mod {
	return func(satisfaction *Satisfaction) error {
		if len(options) != 3 && len(options) != 5 {
			return errors.New("invalid number of options, it must be 3 or 5")
		}
		satisfaction.totalOptions = len(options)
		satisfaction.options = options
		return nil
	}
}
