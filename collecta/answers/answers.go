package answers

import (
	"strconv"
	"strings"

	"github.com/minskylab/collecta/ent/input"
	"github.com/pkg/errors"
)

func AnswerIsKind(kind input.Kind, in []string, options ...map[string]string) (bool, error) {
	if len(in) == 0 {
		return false, errors.New("invalid answers length, please add one answer at lest")
	}

	switch kind {
	case input.KindSatisfaction:
		for _, i := range in {
			if _, err := strconv.ParseFloat(strings.TrimSpace(i), 64); err != nil {
				return false, errors.Wrap(err, "invalid response, that should be a decimal number <0,1> (e.g. 0.4)")
			}
		}

	case input.KindOptions:
		if len(options) == 0 {
			return false, errors.New("please add the alternatives as a map[string]string")
		}
		keys := make([]string, 0)
		values := make([]string, 0)
		for k, v := range options[0] {
			keys = append(keys, strings.TrimSpace(k))
			values = append(values, strings.TrimSpace(v))
		}

		for _, i := range in {
			ok1, ok2 := true, true
			if !strings.Contains(strings.Join(keys, " "), strings.TrimSpace(i)) {
				ok1 = false
			}
			if !strings.Contains(strings.Join(values, " "), strings.TrimSpace(i)) {
				ok2 = false
			}

			if !ok1 && !ok2 {
				return false, errors.New("option not found on the available values")
			}
		}

	case input.KindText:
		for _, i := range in {
			if i == "" {
				return false, errors.New("void answer")
			}
		}

	case input.KindBoolean:
		for _, i := range in {
			if _, err := strconv.ParseBool(i); err != nil {
				if !strings.Contains("yes no si no", i) {
					return false, errors.New("invalid boolean answer")
				}
			}
		}

	default:
		return false, errors.New("invalid input kind, that's so rare")
	}

	return true, nil
}
