package flows

import (
	"github.com/google/uuid"
	"github.com/minskylab/collecta"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func NextState(ctx context.Context, db *collecta.DB, surveyID uuid.UUID) (uuid.UUID, error) {
	surv, err := db.Ent.Survey.Get(ctx, surveyID)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch survey")
	}

	f, err := surv.QueryFlow().Only(ctx)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch query from survey")
	}

	nextState, err := evalProgram(ctx, f.StateTable, input{state: f.State.String(), externalInputs: f.Inputs})
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at eval internal program")
	}

	return uuid.Parse(nextState)
}