package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/minskylab/collecta/errors"

	"fmt"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/commons"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/person"
	"github.com/minskylab/collecta/ent/question"
	"github.com/minskylab/collecta/ent/survey"
)

func (r *queryResolver) Domain(ctx context.Context, id string) (*model.Domain, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
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

	return commons.DomainToGQL(e), nil
}

func (r *queryResolver) Survey(ctx context.Context, id string) (*model.Survey, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
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

	return commons.SurveyToGQL(e), nil
}

func (r *queryResolver) Question(ctx context.Context, id string) (*model.Question, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
	}

	questionID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to parse the domain id")
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
		return nil, errors.Wrap(err, "error at try to search domain related to question")
	}

	if !isOwnerOfQuestionSurveyDomain {
		isQuestionOwner, err := userRequester.QuerySurveys().QueryFlow().QueryQuestions().Where(question.ID(questionID)).Exist(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "question cannot be fetch")
		}
		if !isQuestionOwner {
			return nil, errors.New("resource isn't accessible for you")
		}
	}

	e, err := r.DB.Ent.Question.Get(ctx, questionID)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to get from ent")
	}

	return commons.QuestionToGQL(e), nil
}

func (r *queryResolver) Person(ctx context.Context, id string) (*model.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Profile(ctx context.Context) (*model.Person, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
	}

	return commons.PersonToGQL(userRequester), nil
}

func (r *queryResolver) IsFirstQuestion(ctx context.Context, questionID string) (bool, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return false, errors.New("unauthorized, please include a valid token in your header")
	}

	qID, err := uuid.Parse(questionID)
	if err != nil {
		return false, errors.Wrap(err, "error at try to parse the domain id")
	}

	isOwnerOfQuestionSurveyDomain, err := userRequester.
		QueryAdminOf().
		Where(
			domain.HasSurveysWith(
				survey.HasFlowWith(
					flow.HasQuestionsWith(question.ID(qID)),
				),
			),
		).Exist(ctx)
	if err != nil {
		return false, errors.Wrap(err, "error at try to search domain related to question")
	}

	if !isOwnerOfQuestionSurveyDomain {
		isQuestionOwner, err := userRequester.QuerySurveys().QueryFlow().QueryQuestions().Where(question.ID(qID)).Exist(ctx)
		if err != nil {
			return false, errors.Wrap(err, "question cannot be fetch")
		}
		if !isQuestionOwner {
			return false, errors.New("resource isn't accessible for you")
		}
	}

	q, err := r.DB.Ent.Question.Get(ctx, qID)
	if err != nil {
		return false, errors.Wrap(err, "error at try to get from ent")
	}

	f, err := q.QueryFlow().Only(ctx)
	if err != nil {
		return false, errors.Wrap(err, "error at fetch the question flow")
	}

	return qID == f.InitialState, nil
}

func (r *queryResolver) IsFinalQuestion(ctx context.Context, questionID string) (bool, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return false, errors.New("unauthorized, please include a valid token in your header")
	}

	qID, err := uuid.Parse(questionID)
	if err != nil {
		return false, errors.Wrap(err, "error at try to parse the domain id")
	}

	isOwnerOfQuestionSurveyDomain, err := userRequester.
		QueryAdminOf().
		Where(
			domain.HasSurveysWith(
				survey.HasFlowWith(
					flow.HasQuestionsWith(question.ID(qID)),
				),
			),
		).Exist(ctx)
	if err != nil {
		return false, errors.Wrap(err, "error at try to search domain related to question")
	}

	if !isOwnerOfQuestionSurveyDomain {
		isQuestionOwner, err := userRequester.QuerySurveys().QueryFlow().QueryQuestions().Where(question.ID(qID)).Exist(ctx)
		if err != nil {
			return false, errors.Wrap(err, "question cannot be fetch")
		}
		if !isQuestionOwner {
			return false, errors.New("resource isn't accessible for you")
		}
	}

	q, err := r.DB.Ent.Question.Get(ctx, qID)
	if err != nil {
		return false, errors.Wrap(err, "error at try to get from ent")
	}

	f, err := q.QueryFlow().Only(ctx)
	if err != nil {
		return false, errors.Wrap(err, "error at fetch the question flow")
	}

	return qID == f.TerminationState, nil
}

func (r *queryResolver) SurveyPercent(ctx context.Context, surveyID string) (float64, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return 0.0, errors.New("unauthorized, please include a valid token in your header")
	}

	sID, err := uuid.Parse(surveyID)
	if err != nil {
		return 0.0, errors.Wrap(err, "error at try to parse the domain id")
	}

	isSurveyOwner, err := userRequester.QuerySurveys().Where(survey.ID(sID)).Exist(ctx)
	if err != nil {
		return 0.0, errors.Wrap(err, "error at try to search surveys")
	}

	if !isSurveyOwner {
		isOwnerOfSurveyDomain, err := userRequester.QueryAdminOf().Where(domain.HasSurveysWith(survey.ID(sID))).Exist(ctx)
		if err != nil {
			return 0.0, errors.Wrap(err, "error at try to search domain related to survey")
		}

		if !isOwnerOfSurveyDomain {
			return 0.0, errors.New("resource isn't accessible for you")
		}

	}

	surv, err := r.DB.Ent.Survey.Get(ctx, sID)
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

func (r *queryResolver) LastQuestionOfSurvey(ctx context.Context, surveyID string) (*model.LastSurveyState, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
	}

	sID, err := uuid.Parse(surveyID)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to parse the domain id")
	}

	isSurveyOwner, err := userRequester.QuerySurveys().Where(survey.ID(sID)).Exist(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to search surveys")
	}

	if !isSurveyOwner {
		isOwnerOfSurveyDomain, err := userRequester.QueryAdminOf().Where(domain.HasSurveysWith(survey.ID(sID))).Exist(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "error at try to search domain related to survey")
		}

		if !isOwnerOfSurveyDomain {
			return nil, errors.New("resource isn't accessible for you")
		}

	}

	surv, err := r.DB.Ent.Survey.Get(ctx, sID)
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
		LastQuestion: commons.QuestionToGQL(currentQuestion),
		Percent:      percent,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) User(ctx context.Context, id string) (*model.Person, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
	}

	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to parse the domain id")
	}

	if userRequester.ID != userID {
		requesterIsOwnerOfUser, err := userRequester.QueryAdminOf().Where(domain.HasUsersWith(person.ID(userID))).Exist(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "error at request resource to verify credentials")
		}
		if !requesterIsOwnerOfUser {
			return nil, errors.New("you don't allowed to consume this resource")
		}
	}

	e, err := r.DB.Ent.Person.Get(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to get from ent")
	}

	return commons.PersonToGQL(e), nil
}
