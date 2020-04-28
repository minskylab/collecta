package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/commons"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/auth"
	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/answer"
	"github.com/minskylab/collecta/ent/contact"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/ent/question"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/ent/user"
	"github.com/minskylab/collecta/errors"
)

func (r *accountResolver) Owner(ctx context.Context, obj *model.Account) (*model.User, error) {
	ownerResID, err := commons.OwnerOfAccount(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.User.Query().
		Where(user.HasAccountsWith(account.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.UserToGQL(e), nil
}

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

func (r *contactResolver) Owner(ctx context.Context, obj *model.Contact) (*model.User, error) {
	ownerResID, err := commons.OwnerOfContact(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.User.Query().
		Where(user.HasContactsWith(contact.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.UserToGQL(e), nil
}

func (r *domainResolver) Surveys(ctx context.Context, obj *model.Domain) ([]*model.Survey, error) {
	ownerResID, err := commons.OwnerOfDomain(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID...); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Domain.Query().
		Where(domain.ID(uuid.MustParse(obj.ID))).
		QuerySurveys().
		All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	arr := make([]*model.Survey, 0)
	for _, a := range e {
		if a != nil {
			arr = append(arr, commons.SurveyToGQL(a))
		}
	}

	return arr, nil
}

func (r *domainResolver) Users(ctx context.Context, obj *model.Domain) ([]*model.User, error) {
	ownerResID, err := commons.OwnerOfDomain(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID...); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Domain.Query().
		Where(domain.ID(uuid.MustParse(obj.ID))).
		QueryUsers().
		All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	arr := make([]*model.User, 0)
	for _, a := range e {
		if a != nil {
			arr = append(arr, commons.UserToGQL(a))
		}
	}

	return arr, nil
}

func (r *domainResolver) Admins(ctx context.Context, obj *model.Domain) ([]*model.User, error) {
	ownerResID, err := commons.OwnerOfDomain(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID...); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Domain.Query().
		Where(domain.ID(uuid.MustParse(obj.ID))).
		QueryAdmins().
		All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	arr := make([]*model.User, 0)
	for _, a := range e {
		if a != nil {
			arr = append(arr, commons.UserToGQL(a))
		}
	}

	return arr, nil
}

func (r *flowResolver) Questions(ctx context.Context, obj *model.Flow) ([]*model.Question, error) {
	ownerResID, err := commons.OwnerOfFlow(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Flow.Query().
		Where(flow.ID(uuid.MustParse(obj.ID))).
		QueryQuestions().
		All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	arr := make([]*model.Question, 0)
	for _, a := range e {
		if a != nil {
			arr = append(arr, commons.QuestionToGQL(a))
		}
	}

	return arr, nil
}

func (r *inputResolver) Options(ctx context.Context, obj *model.Input) (map[string]interface{}, error) {
	ownerResID, err := commons.OwnerOfInput(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Input.Get(ctx, uuid.MustParse(obj.ID))
	if err != nil {
		return nil, errors.Wrap(err, "error at ent get")
	}

	m := map[string]interface{}{}
	for k, v := range e.Options {
		m[k] = v
	}

	return m, nil
}

func (r *inputResolver) Question(ctx context.Context, obj *model.Input) (*model.Question, error) {
	ownerResID, err := commons.OwnerOfInput(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Question.Query().
		Where(question.HasInputWith(input.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.QuestionToGQL(e), nil
}

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

func (r *surveyResolver) Flow(ctx context.Context, obj *model.Survey) (*model.Flow, error) {
	ownerResID, err := commons.OwnerOfSurvey(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Survey.Query().
		Where(survey.ID(uuid.MustParse(obj.ID))).
		QueryFlow().
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.FlowToGQL(e), nil
}

func (r *surveyResolver) For(ctx context.Context, obj *model.Survey) (*model.User, error) {
	ownerResID, err := commons.OwnerOfSurvey(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.User.Query().
		Where(user.HasSurveysWith(survey.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.UserToGQL(e), nil
}

func (r *surveyResolver) Owner(ctx context.Context, obj *model.Survey) (*model.Domain, error) {
	ownerResID, err := commons.OwnerOfSurvey(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Domain.Query().
		Where(domain.HasSurveysWith(survey.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.DomainToGQL(e), nil
}

func (r *userResolver) Accounts(ctx context.Context, obj *model.User) (*model.Account, error) {
	ownerResID, err := commons.OwnerOfUser(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Account.Query().
		Where(account.HasOwnerWith(user.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.AccountToGQL(e), nil
}

func (r *userResolver) Contacts(ctx context.Context, obj *model.User) (*model.Contact, error) {
	ownerResID, err := commons.OwnerOfUser(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Contact.Query().
		Where(contact.HasOwnerWith(user.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.ContactToGQL(e), nil
}

func (r *userResolver) Surveys(ctx context.Context, obj *model.User) ([]*model.Survey, error) {
	ownerResID, err := commons.OwnerOfUser(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	u, err := r.DB.Ent.User.Get(ctx, uuid.MustParse(obj.ID))
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch user by user id")
	}

	e, err := r.DB.Ent.User.QuerySurveys(u).All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	arr := make([]*model.Survey, 0)
	for _, a := range e {
		if a != nil {
			arr = append(arr, commons.SurveyToGQL(a))
		}
	}

	return arr, nil
}

func (r *userResolver) Domains(ctx context.Context, obj *model.User) ([]*model.Domain, error) {
	ownerResID, err := commons.OwnerOfUser(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.User.Query().
		Where(user.ID(uuid.MustParse(obj.ID))).
		QueryDomains().
		All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	arr := make([]*model.Domain, 0)
	for _, a := range e {
		if a != nil {
			arr = append(arr, commons.DomainToGQL(a))
		}
	}

	return arr, nil
}

func (r *userResolver) AdminOf(ctx context.Context, obj *model.User) ([]*model.Domain, error) {
	ownerResID, err := commons.OwnerOfUser(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.User.Query().
		Where(user.ID(uuid.MustParse(obj.ID))).
		QueryAdminOf().
		All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	arr := make([]*model.Domain, 0)
	for _, a := range e {
		if a != nil {
			arr = append(arr, commons.DomainToGQL(a))
		}
	}

	return arr, nil
}

// Account returns generated.AccountResolver implementation.
func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

// Answer returns generated.AnswerResolver implementation.
func (r *Resolver) Answer() generated.AnswerResolver { return &answerResolver{r} }

// Contact returns generated.ContactResolver implementation.
func (r *Resolver) Contact() generated.ContactResolver { return &contactResolver{r} }

// Domain returns generated.DomainResolver implementation.
func (r *Resolver) Domain() generated.DomainResolver { return &domainResolver{r} }

// Flow returns generated.FlowResolver implementation.
func (r *Resolver) Flow() generated.FlowResolver { return &flowResolver{r} }

// Input returns generated.InputResolver implementation.
func (r *Resolver) Input() generated.InputResolver { return &inputResolver{r} }

// Question returns generated.QuestionResolver implementation.
func (r *Resolver) Question() generated.QuestionResolver { return &questionResolver{r} }

// Survey returns generated.SurveyResolver implementation.
func (r *Resolver) Survey() generated.SurveyResolver { return &surveyResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type accountResolver struct{ *Resolver }
type answerResolver struct{ *Resolver }
type contactResolver struct{ *Resolver }
type domainResolver struct{ *Resolver }
type flowResolver struct{ *Resolver }
type inputResolver struct{ *Resolver }
type questionResolver struct{ *Resolver }
type surveyResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func validateAuthorization(ctx context.Context, auth *auth.Auth, ownerResources ...uuid.UUID) error {
	userRequester := auth.UserOfContext(ctx)
	if userRequester == nil {
		return errors.New("unauthorized, please include a valid token in your header")
	}

	isOwner := false

	if strings.Contains(strings.Join(userRequester.Roles, " "), "admin") { // is admin
		return nil
	}

	for _, ownerResource := range ownerResources {
		if ownerResource == userRequester.ID {
			isOwner = true
			break
		}
	}

	if !isOwner {
		return errors.New("resource unavailable for you")
	}

	return nil
}
