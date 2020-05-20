package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/minskylab/collecta/errors"

	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/question"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/uuid"
)

func (r *queryResolver) Domain(ctx context.Context, id uuid.UUID) (*ent.Domain, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
	}

	isAdminOfCurrentDomain, err := userRequester.QueryAdminOf().Where(domain.ID(id)).Exist(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "user is not admin of the current domain")
	}

	if !isAdminOfCurrentDomain {
		return nil, errors.New("access not allowed for your token")
	}

	return r.DB.Ent.Domain.Get(ctx, id)
}

func (r *queryResolver) Survey(ctx context.Context, id uuid.UUID) (*ent.Survey, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
	}

	isSurveyOwner, err := userRequester.QuerySurveys().Where(survey.ID(id)).Exist(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to search surveys")
	}

	if !isSurveyOwner {
		isOwnerOfSurveyDomain, err := userRequester.QueryAdminOf().Where(domain.HasSurveysWith(survey.ID(id))).Exist(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "error at try to search domain related to survey")
		}

		if !isOwnerOfSurveyDomain {
			return nil, errors.New("resource isn't accessible for you")
		}

	}

	return r.DB.Ent.Survey.Get(ctx, id)
}

func (r *queryResolver) Question(ctx context.Context, id uuid.UUID) (*ent.Question, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
	}

	isOwnerOfQuestionSurveyDomain, err := userRequester.
		QueryAdminOf().
		Where(
			domain.HasSurveysWith(
				survey.HasFlowWith(
					flow.HasQuestionsWith(question.ID(id)),
				),
			),
		).Exist(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to search domain related to question")
	}

	if !isOwnerOfQuestionSurveyDomain {
		isQuestionOwner, err := userRequester.QuerySurveys().QueryFlow().QueryQuestions().Where(question.ID(id)).Exist(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "question cannot be fetch")
		}
		if !isQuestionOwner {
			return nil, errors.New("resource isn't accessible for you")
		}
	}

	return r.DB.Ent.Question.Get(ctx, id)
}

func (r *queryResolver) Person(ctx context.Context, id uuid.UUID) (*ent.Person, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
	}

	if userRequester.ID != id {
		if strings.Contains(strings.Join(userRequester.Roles, " "), "admin") { // is admin
			return nil, errors.New("forbidden resource")
		}
	}

	return r.DB.Ent.Person.Get(ctx, id)
}

func (r *queryResolver) Profile(ctx context.Context) (*ent.Person, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
	}

	return userRequester, nil
}

func (r *queryResolver) IsFirstQuestion(ctx context.Context, questionID uuid.UUID) (bool, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return false, errors.New("unauthorized, please include a valid token in your header")
	}

	isOwnerOfQuestionSurveyDomain, err := userRequester.
		QueryAdminOf().
		Where(
			domain.HasSurveysWith(
				survey.HasFlowWith(
					flow.HasQuestionsWith(question.ID(questionID)),
				),
			),
		).Exist(ctx)
	if err != nil {
		return false, errors.Wrap(err, "error at try to search domain related to question")
	}

	if !isOwnerOfQuestionSurveyDomain {
		isQuestionOwner, err := userRequester.QuerySurveys().QueryFlow().QueryQuestions().Where(question.ID(questionID)).Exist(ctx)
		if err != nil {
			return false, errors.Wrap(err, "question cannot be fetch")
		}
		if !isQuestionOwner {
			return false, errors.New("resource isn't accessible for you")
		}
	}

	q, err := r.DB.Ent.Question.Get(ctx, questionID)
	if err != nil {
		return false, errors.Wrap(err, "error at try to get from ent")
	}

	f, err := q.QueryFlow().Only(ctx)
	if err != nil {
		return false, errors.Wrap(err, "error at fetch the question flow")
	}

	return questionID == f.InitialState, nil
}

func (r *queryResolver) IsFinalQuestion(ctx context.Context, questionID uuid.UUID) (bool, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return false, errors.New("unauthorized, please include a valid token in your header")
	}

	isOwnerOfQuestionSurveyDomain, err := userRequester.
		QueryAdminOf().
		Where(
			domain.HasSurveysWith(
				survey.HasFlowWith(
					flow.HasQuestionsWith(question.ID(questionID)),
				),
			),
		).Exist(ctx)
	if err != nil {
		return false, errors.Wrap(err, "error at try to search domain related to question")
	}

	if !isOwnerOfQuestionSurveyDomain {
		isQuestionOwner, err := userRequester.QuerySurveys().QueryFlow().QueryQuestions().Where(question.ID(questionID)).Exist(ctx)
		if err != nil {
			return false, errors.Wrap(err, "question cannot be fetch")
		}
		if !isQuestionOwner {
			return false, errors.New("resource isn't accessible for you")
		}
	}

	q, err := r.DB.Ent.Question.Get(ctx, questionID)
	if err != nil {
		return false, errors.Wrap(err, "error at try to get from ent")
	}

	f, err := q.QueryFlow().Only(ctx)
	if err != nil {
		return false, errors.Wrap(err, "error at fetch the question flow")
	}

	return questionID == f.TerminationState, nil
}

func (r *queryResolver) SurveyPercent(ctx context.Context, surveyID uuid.UUID) (float64, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return 0.0, errors.New("unauthorized, please include a valid token in your header")
	}

	isSurveyOwner, err := userRequester.QuerySurveys().Where(survey.ID(surveyID)).Exist(ctx)
	if err != nil {
		return 0.0, errors.Wrap(err, "error at try to search surveys")
	}

	if !isSurveyOwner {
		isOwnerOfSurveyDomain, err := userRequester.QueryAdminOf().Where(domain.HasSurveysWith(survey.ID(surveyID))).Exist(ctx)
		if err != nil {
			return 0.0, errors.Wrap(err, "error at try to search domain related to survey")
		}

		if !isOwnerOfSurveyDomain {
			return 0.0, errors.New("resource isn't accessible for you")
		}

	}

	surv, err := r.DB.Ent.Survey.Get(ctx, surveyID)
	if err != nil {
		return 0.0, errors.Wrap(err, "error at fetch survey")
	}

	answeredQuestions, err := surv.QueryFlow().QueryQuestions().Where(question.HasAnswers()).All(ctx)
	if err != nil {
		return 0.0, errors.Wrap(err, "error at fetch answered questions")
	}

	totalQuestions, err := surv.QueryFlow().QueryQuestions().All(ctx)
	if err != nil {
		return 0.0, errors.Wrap(err, "error at fetch total questions of survey")
	}

	percent := float64(len(answeredQuestions)) / float64(len(totalQuestions))

	return percent, nil
}

func (r *queryResolver) LastQuestionOfSurvey(ctx context.Context, surveyID uuid.UUID) (*model.LastSurveyState, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
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

	surv, err := r.DB.Ent.Survey.Get(ctx, surveyID)
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

	answeredQuestions, err := surv.QueryFlow().QueryQuestions().Where(question.HasAnswers()).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch answered questions")
	}

	totalQuestions, err := surv.QueryFlow().QueryQuestions().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch total questions of survey")
	}

	percent := float64(len(answeredQuestions)) / float64(len(totalQuestions))

	return &model.LastSurveyState{
		LastQuestion: currentQuestion,
		Percent:      percent,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
