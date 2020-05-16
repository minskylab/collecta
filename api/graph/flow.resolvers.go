package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/minskylab/collecta/errors"

	"fmt"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/commons"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/ent/flow"
)

func (r *flowResolver) Survey(ctx context.Context, obj *model.Flow) (*model.Survey, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *flowResolver) Questions(ctx context.Context, obj *model.Flow) ([]*model.Question, error) {
	ownerResID, err := commons.OwnerOfFlow(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Flow.Query().
		Where(flow.ID(uuid.MustParse(obj.ID))).
		QueryQuestions().
		All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	arr := make([]*model.Question, 0)
	for _, a := range e {
		if a != nil {
			arr = append(arr, commons.QuestionToGQL(a))
		}
	}

	return arr, nil
}

// Flow returns generated.FlowResolver implementation.
func (r *Resolver) Flow() generated.FlowResolver { return &flowResolver{r} }

type flowResolver struct{ *Resolver }
