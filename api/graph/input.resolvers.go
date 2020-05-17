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
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/ent/question"
)

func (r *inputResolver) Question(ctx context.Context, obj *ent.Input) (*ent.Question, error) {
	ownerResID, err := commons.OwnerOfInput(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Question.Query().
		Where(question.HasInputWith(input.ID(obj.ID))).
		Only(ctx)
}

// Input returns generated.InputResolver implementation.
func (r *Resolver) Input() generated.InputResolver { return &inputResolver{r} }

type inputResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *inputResolver) Options(ctx context.Context, obj *ent.Input) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}
