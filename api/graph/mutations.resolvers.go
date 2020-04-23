package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/flows"
)

func (r *mutationResolver) AnswerQuestion(ctx context.Context, token string, questionID string, answer []string) (*model.Survey, error) {
	q, err := r.DB.Ent.Question.Get(ctx, uuid.MustParse(questionID))
	if err != nil {
		return nil, errors.Wrap(err, "error at try to fetch question")
	}

	// TODO: Validate all

	_, err = r.DB.Ent.Answer.Create().
		SetID(uuid.New()).
		SetQuestion(q).
		SetResponses(answer).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to create new answer")
	}

	f, err := q.QueryFlow().Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch flow")
	}

	surv, err := r.DB.Ent.Survey.Query().Where(survey.HasFlowWith(flow.ID(f.ID))).Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch survey")
	}

	// TODO: That is completely incorrect, please Bregy solve it fast!
	nexState, err := flows.NextState(ctx, r.DB, surv.ID)
	if err != nil {
		return nil, errors.Wrap(err, "error at calculate the next state")
	}

	if _, err = r.DB.Ent.Flow.UpdateOneID(f.ID).SetState(nexState).Save(ctx); err != nil {
		return nil, errors.Wrap(err, "error at update flow with the next state")
	}

	surv, err = r.DB.Ent.Survey.UpdateOneID(surv.ID).SetLastInteraction(time.Now()).Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at update last interaction of the survey")
	}

	return &model.Survey{
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
