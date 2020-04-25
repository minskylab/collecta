package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/collecta/answers"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/ent/question"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/errors"
	"github.com/minskylab/collecta/flows"
)

func (r *mutationResolver) AnswerQuestion(ctx context.Context, token string, questionID string, answer []string) (*model.Survey, error) {
	userRequester, err := r.Auth.VerifyJWTToken(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "invalid token, probably user not registered")
	}

	qID, err := uuid.Parse(questionID)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to parse the domain id")
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
		return nil, errors.Wrap(err, "error at try to search domain related to question")
	}

	if !isOwnerOfQuestionSurveyDomain {
		isQuestionOwner, err := userRequester.QuerySurveys().QueryFlow().QueryQuestions().Where(question.ID(qID)).Exist(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "question cannot be fetch")
		}
		if !isQuestionOwner {
			return nil, errors.New("resource isn't accessible for you")
		}
	}

	q, err := r.DB.Ent.Question.Get(ctx, qID)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to fetch question")
	}

	in, err := q.QueryInput().Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at query input from your question")
	}

	var answerIsOk bool
	switch in.Kind {
	case input.KindSatisfaction:
		answerIsOk, err = answers.AnswerIsSatisfaction(answer, in.Multiple)
		if err != nil {
			return nil, errors.Wrap(err, "error at validate your answer")
		}

	case input.KindOptions:
		answerIsOk, err = answers.AnswerIsOption(answer, in.Options, in.Multiple)
		if err != nil {
			return nil, errors.Wrap(err, "error at validate your answer")
		}

	case input.KindText:
		answerIsOk, err = answers.AnswerIsText(answer, in.Multiple)
		if err != nil {
			return nil, errors.Wrap(err, "error at validate your answer")
		}

	case input.KindBoolean:
		answerIsOk, err = answers.AnswerIsBoolean(answer, in.Multiple)
		if err != nil {
			return nil, errors.Wrap(err, "error at validate your answer")
		}

	default:
		return nil, errors.New("invalid input kind, that's so rare")
	}

	if !answerIsOk {
		return nil, errors.New("invalid answer, please choose a correct one")
	}

	f, err := q.QueryFlow().Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch flow")
	}

	if q.ID != f.State {
		return nil, errors.New("you only can answered a current question, as state, in a flow")
	}

	_, err = r.DB.Ent.Answer.Create().
		SetID(uuid.New()).
		SetQuestion(q).
		SetResponses(answer).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to create new answer")
	}

	surv, err := r.DB.Ent.Survey.Query().Where(survey.HasFlowWith(flow.ID(f.ID))).Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch survey")
	}

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
