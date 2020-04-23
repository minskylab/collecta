package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/survey"

	"github.com/pkg/errors"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
)

func (r *queryResolver) Domain(ctx context.Context, token string, id string) (*model.Domain, error) {
	userRequester, err := r.Auth.VerifyJWTToken(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "invalid token, probably user not registered")
	}

	domainID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to parse the domain id")
	}

	isAdminOfCurrentDomain, err := userRequester.QueryAdminOf().Where(domain.ID(domainID)).Exist(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "user is not admin of the current domain")
	}

	if !isAdminOfCurrentDomain {
		return nil, errors.New("access not allowed for your token")
	}

	e, err := r.DB.Ent.Domain.Get(ctx, domainID)
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

func (r *queryResolver) Survey(ctx context.Context, token string, id string) (*model.Survey, error) {
	userRequester, err := r.Auth.VerifyJWTToken(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "invalid token, probably user not registered")
	}

	surveyID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to parse the domain id")
	}

	isSurveyOwner, err := userRequester.QuerySurveys().Where(survey.ID(surveyID)).Exist(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to search surveys")
	}

	if !isSurveyOwner {
		isOwnerOfSurveyDomain, err := userRequester.QueryAdminOf().Where(domain.HasSurveysWith(survey.ID(surveyID))).Exist(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "error at try to search domain related to survey")
		}

		if !isOwnerOfSurveyDomain {
			return nil, errors.New("resource isn't accessible for you")
		}

	}

	e, err := r.DB.Ent.Survey.Get(ctx, surveyID)
	if err != nil {
		return nil, errors.Wrap(err, "error at try resource  to get from ent")
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

func (r *queryResolver) Question(ctx context.Context, token string, id string) (*model.Question, error) {
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

func (r *queryResolver) User(ctx context.Context, token string, id string) (*model.User, error) {
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

func (r *queryResolver) UserByToken(ctx context.Context, token string) (*model.User, error) {
	e, err := r.Auth.VerifyJWTToken(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "error at verify jwt")
	}

	return &model.User{
		ID:           e.ID.String(),
		Name:         e.Name,
		Username:     e.Username,
		LastActivity: e.LastActivity,
		Picture:      e.Picture,
	}, nil
}

func (r *queryResolver) IsFirstQuestion(ctx context.Context, token string, questionID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IsFinalQuestion(ctx context.Context, token string, questionID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) LastQuestionOfSurvey(ctx context.Context, token string, questionID string) (*model.Question, error) {
	surv, err := r.DB.Ent.Survey.Get(ctx, uuid.MustParse(questionID))
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch survey")
	}

	f, err := surv.QueryFlow().Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch flow")
	}

	currentQuestion, err := r.DB.Ent.Question.Get(ctx, f.State)
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
