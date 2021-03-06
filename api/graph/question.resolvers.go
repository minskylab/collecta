package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/minskylab/collecta/errors"

	"github.com/minskylab/collecta/api/commons"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/ent/question"
)

func (r *questionResolver) Answers(ctx context.Context, obj *ent.Question) ([]*ent.Answer, error) {
	ownerResID, err := commons.OwnerOfQuestion(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Question.Query().
		Where(question.ID(obj.ID)).
		QueryAnswers().
		All(ctx)
}

func (r *questionResolver) Input(ctx context.Context, obj *ent.Question) (*ent.Input, error) {
	ownerResID, err := commons.OwnerOfQuestion(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Input.Query().
		Where(input.HasQuestionWith(question.ID(obj.ID))).
		Only(ctx)
}

func (r *questionResolver) Flow(ctx context.Context, obj *ent.Question) (*ent.Flow, error) {
	ownerResID, err := commons.OwnerOfQuestion(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Flow.Query().
		Where(flow.HasQuestionsWith(question.ID(obj.ID))).
		Only(ctx)
}

// Question returns generated.QuestionResolver implementation.
func (r *Resolver) Question() generated.QuestionResolver { return &questionResolver{r} }

type questionResolver struct{ *Resolver }
