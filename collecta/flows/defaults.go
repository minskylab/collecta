package flows

import (
	"strings"

	"github.com/minskylab/collecta/uuid"
)

const defaultSequentialProgram = `
// Forward function calculates the next state of the system
forward := func(state, input) {
    res := state
    for i, v in questions {
      if v == state { 
          if i < len(questions) - 1 {
            res = questions[i + 1] 
          } 
      }
    }
    return res
}

// Backward function calculates the next state of the system
backward := func(state, input) {
    res := state
    for i, v in questions {
      if v == state { 
          if i > 0 {
            res = questions[i - 1] 
          } 
      }
    }
    return res
}

next := forward(state, input)
last := backward(state, input)
`
func defaultSequentialProgramFromQuestions(questions []string) string {
	finalQuestions := make([]string, 0)
	for _, q := range questions {
		finalQuestions = append(finalQuestions, "\""+q+"\"")
	}

	strArray := "([\n " + strings.Join(finalQuestions, ",\n ") + "\n])\n"
	start := "questions := immutable" + strArray

	script := start + defaultSequentialProgram

	return script
}

func DefaultSequentialProgram(questions []uuid.UUID) string {
	parsedQuestions := make([]string, 0)
	for _, q := range questions {
		parsedQuestions = append(parsedQuestions, q.String())
	}
	return defaultSequentialProgramFromQuestions(parsedQuestions)
}
