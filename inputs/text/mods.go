package text

import (
	"regexp"

	"github.com/pkg/errors"
)

type Mod func(text *Text) error

func WithDefaultValue(defVal string) Mod {
	return func(text *Text) error {
		text.defaultValue = defVal
		return nil
	}
}

func WithMaxChars(maxChars int) Mod {
	return func(text *Text) error {
		if maxChars < 0 {
			return errors.New("invalid chars length, it must be positive")
		}
		text.maxChars = maxChars
		return nil
	}
}

func WithExpected(regExpression string) Mod {
	return func(text *Text) error {
		re, err := regexp.Compile(regExpression)
		if err != nil {
			return err
		}
		text.expected = re
		return nil
	}
}
