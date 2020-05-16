package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/minskylab/collecta/errors"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/commons"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/ent/question"
)

func (r *inputResolver) Question(ctx context.Context, obj *model.Input) (*model.Question, error) {
	ownerResID, err := commons.OwnerOfInput(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Question.Query().
		Where(question.HasInputWith(input.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.QuestionToGQL(e), nil
}

// Input returns generated.InputResolver implementation.
func (r *Resolver) Input() generated.InputResolver { return &inputResolver{r} }

type inputResolver struct{ *Resolver }
