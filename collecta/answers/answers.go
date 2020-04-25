package answers

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// AnswerKind try to give a name to the answer
type AnswerKind string

const Satisfaction AnswerKind = "satisfaction"
const Option AnswerKind = "option"
const Text AnswerKind = "text"
const Boolean AnswerKind = "boolean"

func validate(input []string, kind AnswerKind, options ...map[string]string) (bool, bool, error) {
	if len(input) == 0 {
		return false, false, errors.New("invalid answers length, please add one answer at lest")
	}

	for _, in := range input {
		in = strings.TrimSpace(in)
		switch kind {
		case Satisfaction:
			if _, err := strconv.ParseFloat(in, 64); err != nil {
				return false, false, errors.Wrap(err, "invalid response, that should be a decimal number <0,1> (e.g. 0.4)")
			}
		case Text:
			if in == "" {
				return false, false, errors.New("void answer")
			}
		case Option:
			if len(options) == 0 {
				return false, false, errors.New("please add the alteratives as a map[string]string")
			}
			keys := make([]string, 0)
			values := make([]string, 0)
			for k, v := range options[0] {
				keys = append(keys, strings.TrimSpace(k))
				values = append(values, strings.TrimSpace(v))
			}

			ok1, ok2 := true, true
			if !strings.Contains(strings.Join(keys, " "), strings.TrimSpace(in)) {
				ok1 = false
			}
			if !strings.Contains(strings.Join(values, " "), strings.TrimSpace(in)) {
				ok2 = false
			}

			if !ok1 && !ok2 {
				return false, false, errors.New("option not found on the available values")
			}
		case Boolean:
			if _, err := strconv.ParseBool(in); err != nil {
				if !strings.Contains("yes no si no", in) {
					return false, false, errors.New("invalid boolean answer")
				}
			}
		default:
			return false, false, errors.New("invalid kind of answer, are valid: satisfaction|text|option|boolean")
		}
	}

	m := false
	if len(input) > 0 {
		m = true
	}

	return true, m, nil
}

func answerIsOfKind(kind AnswerKind, input []string, acceptMultiple bool, options ...map[string]string) (bool, error) {
	valid, multiple, err := validate(input, kind, options...)
	if err != nil {
		return false, errors.Wrap(err, "error at validate your input")
	}

	if multiple && !acceptMultiple {
		return false, nil
	}

	return valid, nil
}

func AnswerIsSatisfaction(input []string, acceptMultiple bool) (bool, error) {
	return answerIsOfKind(Satisfaction, input, acceptMultiple)
}

func AnswerIsOption(input []string, options map[string]string, acceptMultiple bool) (bool, error) {
	return answerIsOfKind(Option, input, acceptMultiple, options)
}

func AnswerIsBoolean(input []string, acceptMultiple bool) (bool, error) {
	return answerIsOfKind(Boolean, input, acceptMultiple)
}

func AnswerIsText(input []string, acceptMultiple bool) (bool, error) {
	return answerIsOfKind(Text, input, acceptMultiple)
}
