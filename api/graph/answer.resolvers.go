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
	"github.com/minskylab/collecta/ent/answer"
	"github.com/minskylab/collecta/ent/question"
)

func (r *answerResolver) Question(ctx context.Context, obj *model.Answer) (*model.Question, error) {
	ownerResID, err := commons.OwnerOfAnswer(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Question.Query().
		Where(question.HasAnswersWith(answer.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.QuestionToGQL(e), nil
}

// Answer returns generated.AnswerResolver implementation.
func (r *Resolver) Answer() generated.AnswerResolver { return &answerResolver{r} }

type answerResolver struct{ *Resolver }
