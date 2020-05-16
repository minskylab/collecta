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
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/ent/question"
)

func (r *questionResolver) Answers(ctx context.Context, obj *model.Question) ([]*model.Answer, error) {
	ownerResID, err := commons.OwnerOfQuestion(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Question.Query().
		Where(question.ID(uuid.MustParse(obj.ID))).
		QueryAnswers().
		All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	arr := make([]*model.Answer, 0)
	for _, a := range e {
		if a != nil {
			arr = append(arr, commons.AnswerToGQL(a))
		}
	}

	return arr, nil
}

func (r *questionResolver) Input(ctx context.Context, obj *model.Question) (*model.Input, error) {
	ownerResID, err := commons.OwnerOfQuestion(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Input.Query().
		Where(input.HasQuestionWith(question.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.InputToGQL(e), nil
}

func (r *questionResolver) Flow(ctx context.Context, obj *model.Question) (*model.Flow, error) {
	ownerResID, err := commons.OwnerOfQuestion(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Flow.Query().
		Where(flow.HasQuestionsWith(question.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.FlowToGQL(e), nil
}

// Question returns generated.QuestionResolver implementation.
func (r *Resolver) Question() generated.QuestionResolver { return &questionResolver{r} }

type questionResolver struct{ *Resolver }
