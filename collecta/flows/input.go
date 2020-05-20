package flows

import (
	"context"

	"github.com/minskylab/collecta/db"
	"github.com/minskylab/collecta/errors"
	"github.com/minskylab/collecta/uuid"
)

// NextState run the flow script program and calculates the next question id
func NextState(ctx context.Context, db *db.DB, surveyID uuid.UUID) (uuid.UUID, error) {
	surv, err := db.Ent.Survey.Get(ctx, surveyID)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch survey")
	}

	f, err := surv.QueryFlow().Only(ctx)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch query from survey")
	}

	res, err := evalProgram(ctx, f.StateTable, input{state: f.State.String(), externalInputs: f.Inputs})
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at eval internal program")
	}

	return uuid.Parse(res.next)
}

// LastState run the flow script program and calculates the last question id
func LastState(ctx context.Context, db *db.DB, surveyID uuid.UUID) (uuid.UUID, error) {
	surv, err := db.Ent.Survey.Get(ctx, surveyID)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch survey")
	}

	f, err := surv.QueryFlow().Only(ctx)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch query from survey")
	}

	res, err := evalProgram(ctx, f.StateTable, input{state: f.State.String(), externalInputs: f.Inputs})
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at eval internal program")
	}

	return uuid.Parse(res.last)
}
