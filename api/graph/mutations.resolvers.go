package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/minskylab/collecta/errors"

	"github.com/microcosm-cc/bluemonday"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/collecta/answers"
	"github.com/minskylab/collecta/collecta/flows"
	"github.com/minskylab/collecta/collecta/surveys"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/person"
	"github.com/minskylab/collecta/ent/question"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) AnswerQuestion(ctx context.Context, questionID uuid.UUID, answer []string) (*ent.Survey, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
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

	q, err := r.DB.Ent.Question.Get(ctx, questionID)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to fetch question")
	}

	in, err := q.QueryInput().Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at query input from your question")
	}

	policy := bluemonday.UGCPolicy()
	policy.AllowElements("h1").AllowElements("h2").AllowElements("h3")
	policy.AllowAttrs("href").OnElements("a")
	policy.AllowElements("p")
	// TODO: Sanitize Policy should be into core struct (Collecta)

	for i, v := range answer {
		answer[i] = policy.Sanitize(v)
	}

	opts := map[string]string{}
	for k, v := range in.Options {
		opts[k], _ = v.(string)
	}

	answerIsOk, err := answers.AnswerIsKind(in.Kind, answer, opts) // TODO, optimize: only pass a *Input

	if !answerIsOk {
		return nil, errors.New("invalid answer, please choose a correct one")
	}

	f, err := q.QueryFlow().Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch flow")
	}

	if q.ID != f.State {
		return nil, errors.New("you can only answer the current question in the flow")
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

	return r.DB.Ent.Survey.UpdateOneID(surv.ID).SetLastInteraction(time.Now()).Save(ctx)
}

func (r *mutationResolver) BackwardSurvey(ctx context.Context, surveyID uuid.UUID) (*ent.Survey, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
	}

	isOwnerOfSurveyDomain, err := userRequester.
		QueryAdminOf().
		Where(
			domain.HasSurveysWith(
				survey.ID(surveyID),
			),
		).Exist(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to search domain related to question")
	}

	if !isOwnerOfSurveyDomain {
		ownerOfSurvey, err := userRequester.QuerySurveys().Where(survey.ID(surveyID)).Exist(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "error at try to search survey from owns")
		}

		if !ownerOfSurvey {
			return nil, errors.Wrap(err, "operation not allowed for you")
		}
	}

	lastState, err := flows.LastState(ctx, r.DB, surveyID)
	if err != nil {
		return nil, errors.Wrap(err, "error at calculate the last state")
	}

	currentFlow, err := r.DB.Ent.Flow.Query().Where(flow.HasSurveyWith(survey.ID(surveyID))).Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch the flow of your survey")
	}

	currentFlow, err = r.DB.Ent.Flow.UpdateOneID(currentFlow.ID).
		SetState(lastState).
		SetPastState(currentFlow.State).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at update your survey flow")
	}

	// TODO: Update now or create hook to update the survey last interaction, that is important

	return currentFlow.QuerySurvey().Only(ctx)
}

func (r *mutationResolver) LoginByPassword(ctx context.Context, username string, password string) (*model.LoginResponse, error) {
	loginAccount, err := r.DB.Ent.Account.Query().
		Where(account.And(
			account.TypeEQ(account.TypeEmail),                             // email type
			account.Or(account.RemoteID(username), account.Sub(username)), // by username == sub or username == remoteID
		)).
		Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to get your login account")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(loginAccount.Secret), []byte(password)); err != nil {
		return nil, errors.Wrap(err, "invalid password its'nt correct")
	}

	loginUser, err := loginAccount.QueryOwner().Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch user related to your account")
	}

	jwtToken, err := r.Auth.GenerateTokenByUser(loginUser)
	if err != nil {
		return nil, errors.Wrap(err, "error at create a new jwt token")
	}

	return &model.LoginResponse{Token: jwtToken}, nil
}

func (r *mutationResolver) UpdatePassword(ctx context.Context, oldPassword string, newPassword string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateNewDomain(ctx context.Context, draft model.DomainCreator) (*ent.Domain, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
	}

	if !strings.Contains(strings.Join(userRequester.Roles, " "), "admin") {
		return nil, errors.New("invalid user, you need to be an admin to create new domain")
	}

	// continue if user is an admin

	return r.DB.Ent.Domain.Create().
		SetID(uuid.New()).
		SetName(draft.Name).
		SetEmail(draft.Email).
		SetDomain(draft.Domain).
		SetCallback(draft.Callback).
		SetTags(draft.Tags).
		AddAdmins(userRequester).
		Save(ctx)
}

func (r *mutationResolver) GenerateSurveys(ctx context.Context, domainSelector model.SurveyDomain, draft model.SurveyGenerator) (*model.SuveyGenerationResult, error) {
	userRequester := r.Auth.UserOfContext(ctx)
	if userRequester == nil {
		return nil, errors.New("unauthorized, please include a valid token in your header")
	}

	var targetDomain *ent.Domain
	var err error
	if strings.Contains(strings.Join(userRequester.Roles, " "), "admin") { // if is super admin
		if domainSelector.ByID != nil {
			targetDomain, err = r.DB.Ent.Domain.Get(ctx, *domainSelector.ByID)
			if err != nil {
				return nil, errors.Wrap(err, "error at fetch domain")
			}
		} else if domainSelector.ByDomainName != nil {
			targetDomain, err = r.DB.Ent.Domain.Query().
				Where(domain.Name(*domainSelector.ByDomainName)).
				Only(ctx)
			if err != nil {
				return nil, errors.Wrap(err, "error at fetch domain, probably you aren't an admin for this domain ")
			}
		} else {
			return nil, errors.New("invalid domain selector, please specify one of between 'by domain' or 'by id'")
		}
	} else { // or if admin of the related domain
		if domainSelector.ByID != nil { // by id
			targetDomain, err = r.DB.Ent.Domain.Query().
				Where(domain.And(domain.ID(*domainSelector.ByID), domain.HasAdminsWith(person.ID(userRequester.ID)))).
				Only(ctx)
			if err != nil {
				return nil, errors.Wrap(err, "error at fetch domain, probably you aren't an admin for this domain ")
			}

		} else if domainSelector.ByDomainName != nil { // by domain name
			targetDomain, err = r.DB.Ent.Domain.Query().
				Where(domain.And(
					domain.Name(*domainSelector.ByDomainName),
					domain.HasAdminsWith(person.ID(userRequester.ID)),
				)).
				Only(ctx)
			if err != nil {
				return nil, errors.Wrap(err, "error at fetch domain, probably you aren't an admin for this domain ")
			}
		} else { // invalid
			return nil, errors.New("invalid domain selector, please specify one of between 'by domain' or 'by id'")
		}

		if targetDomain == nil {
			return nil, errors.New("invalid user, you need to be an admin to create new domain")
		}
	}

	generatedSurveys, err := surveys.GenerateSurveys(ctx, r.DB, targetDomain.ID, draft)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to generate your surveys")
	}

	return &model.SuveyGenerationResult{
		How:     len(generatedSurveys),
		Surveys: generatedSurveys,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
