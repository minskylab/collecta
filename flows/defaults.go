package flows

import (
	"strings"

	"github.com/google/uuid"
)

func defaultSequentialProgramFromQuestions(questions []string) string {
	strArray := "[" + strings.Join(questions, ",") + "]"
	script := "questions = immutable(" + strArray + ")\n" +
			  "res := state\n" +
		      "for i, v in questions {\n" +
			  "  if v == state { res = questions[i+1] }\n" +
			  "}"
	return script
}

func DefaultSequentialProgram(questions []uuid.UUID) string {
	parsedQuestions := make([]string, 0)
	for _, q := range questions {
		parsedQuestions = append(parsedQuestions, q.String())
	}
	return defaultSequentialProgramFromQuestions(parsedQuestions)
}