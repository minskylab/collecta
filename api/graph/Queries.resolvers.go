package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/minskylab/collecta/flows"
	"github.com/pkg/errors"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
)

func (r *queryResolver) Domain(ctx context.Context, id string) (*model.Domain, error) {
	e, err := r.DB.Ent.Domain.Get(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, errors.Wrap(err, "error at try to get from ent")
	}

	return &model.Domain{
		ID:             e.ID.String(),
		Tags:           e.Tags,
		Name:           e.Name,
		Email:          e.Email,
		Domain:         e.Domain,
		CollectaDomain: e.CollectaDomain,
	}, nil
}

func (r *queryResolver) Survey(ctx context.Context, id string) (*model.Survey, error) {
	e, err := r.DB.Ent.Survey.Get(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, errors.Wrap(err, "error at try to get from ent")
	}

	return &model.Survey{
		ID:              e.ID.String(),
		Tags:            e.Tags,
		LastInteraction: e.LastInteraction,
		DueDate:         e.DueDate,
		Title:           e.Title,
		Description:     e.Description,
	}, nil
}

func (r *queryResolver) Question(ctx context.Context, id string) (*model.Question, error) {
	e, err := r.DB.Ent.Question.Get(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, errors.Wrap(err, "error at try to get from ent")
	}

	return &model.Question{
		ID:          e.ID.String(),
		Hash:        e.Hash,
		Title:       e.Title,
		Description: e.Description,
		Anonymous:   e.Anonymous,
	}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	e, err := r.DB.Ent.User.Get(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, errors.Wrap(err, "error at try to get from ent")
	}

	return &model.User{
		ID:           e.ID.String(),
		Name:         e.Name,
		Username:     e.Username,
		LastActivity: e.LastActivity,
		Picture:      e.Picture,
	}, nil
}

func (r *queryResolver) LastQuestionOfSurvey(ctx context.Context, id string) (*model.Question, error) {
	surv, err := r.DB.Ent.Survey.Get(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch survey")
	}

	f, err := surv.QueryFlow().Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch flow")
	}

	f, err = flows.PerformFlowStateUpdate(ctx, r.DB.Ent, f)
	if err != nil {
		return nil, errors.Wrap(err, "error at perform flow state update")
	}

	finalFlow, err := r.DB.Ent.Flow.Get(ctx, f.ID)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch final flow")
	}

	currentQuestion, err := r.DB.Ent.Question.Get(ctx, finalFlow.State)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch last question")
	}

	return &model.Question{
		ID:          currentQuestion.ID.String(),
		Hash:        currentQuestion.Hash,
		Title:       currentQuestion.Title,
		Description: currentQuestion.Description,
		Anonymous:   currentQuestion.Anonymous,
		// Metadata:    nil,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
