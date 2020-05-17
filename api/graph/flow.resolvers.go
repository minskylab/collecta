package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/minskylab/collecta/errors"

	"fmt"

	"github.com/minskylab/collecta/api/commons"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/flow"
)

func (r *flowResolver) Survey(ctx context.Context, obj *ent.Flow) (*ent.Survey, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *flowResolver) Questions(ctx context.Context, obj *ent.Flow) ([]*ent.Question, error) {
	ownerResID, err := commons.OwnerOfFlow(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Flow.Query().
		Where(flow.ID(obj.ID)).
		QueryQuestions().
		All(ctx)
}

// Flow returns generated.FlowResolver implementation.
func (r *Resolver) Flow() generated.FlowResolver { return &flowResolver{r} }

type flowResolver struct{ *Resolver }
