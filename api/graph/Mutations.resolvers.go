package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"


	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/pkg/errors"
)

func (r *mutationResolver) AnswerQuestion(ctx context.Context, input *model.QuestionResponse) (*model.Survey, error) {
	questionID := input.ID
	q, err := r.DB.Ent.Question.Get(ctx, uuid.MustParse(questionID))
	if err != nil {
		return nil, errors.Wrap(err, "error at try to fetch question")
	}

	// TODO: Validate all

	_, err = r.DB.Ent.Answer.Create().
		SetID(uuid.New()).
		SetQuestion(q).
		SetResponses(input.Answer).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to create new answer")
	}

	f, err := q.QueryFlow().Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err , "error at fetch flow")
	}

	surv, err :=  r.DB.Ent.Survey.Query().Where(survey.HasFlowWith(flow.ID(f.ID))).Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err , "error at fetch survey")
	}

	return & model.Survey{
		ID:              surv.ID.String(),
		Tags:            surv.Tags,
		LastInteraction: surv.LastInteraction,
		DueDate:         surv.DueDate,
		Title:           surv.Title,
		Description:     surv.Description,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
