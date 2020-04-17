package flows

import (
	"strings"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/question"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type instruction struct {}

func lineToInstruction(line string) string {
	return strings.TrimSpace(line[2:])
}

func evalInstruction(instruction string, currentState string) string {
	if strings.Contains(instruction, "->") {
		chunks := strings.Split(instruction, "->")
		ant := ""
		con := ""
		if len(chunks) < 2 {
			return currentState
		}
		if len(chunks) > 2 {
			ant = strings.TrimSpace(chunks[0])
			con = strings.TrimSpace(chunks[len(chunks)-1])
		} else {
			ant = strings.TrimSpace(chunks[0])
			con = strings.TrimSpace(chunks[1])
		}

		if ant == currentState {
			return con
		}
	}

	return currentState
}

func interpretFlow(ctx context.Context, db *ent.Client, f *ent.Flow) (*ent.Flow, error) {
	lines := strings.Split(f.StateTable, "\n")

	nextState := "not_next_state"

	for _, l := range lines {
		instruction := lineToInstruction(l)
		pseudoNextState := evalInstruction(instruction, f.State.String())
		if pseudoNextState != f.State.String() {
			q, err := f.QueryQuestions().Where(question.ID(uuid.MustParse(pseudoNextState))).Only(ctx)
			if err != nil {
				return nil, errors.Wrap(err, "error at try to extract next state question")
			}
			questionWasAnswered, err := q.QueryAnswers().Exist(ctx)
			if err != nil {
				return nil, errors.Wrap(err , "error at verify if answers exists on question")
			}

			answers, err := q.QueryAnswers().IDs(ctx)
			if err != nil {
				// return nil, errors.Wrap(err, "error at query all answers of question")
				return f, nil
			}

			if questionWasAnswered && len(answers) > 1 {
				nextState = pseudoNextState
				break
			}

			return f, nil
		}
	}

	if nextState != "not_next_state" {
		updatedFlow, err := db.Flow.UpdateOneID(f.ID).
			SetState(uuid.MustParse(nextState)).
			Save(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "error at try to update f")
		}
		return updatedFlow, nil
	} else {
		// surveyDone
		// db.Flow.Query().
	}

	return f, nil

}

func PerformFlowStateUpdate(ctx context.Context, db *ent.Client, f *ent.Flow) (*ent.Flow, error) {
	return interpretFlow(ctx, db, f)
}