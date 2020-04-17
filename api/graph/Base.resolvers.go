package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/answer"
	"github.com/minskylab/collecta/ent/contact"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/ent/question"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/ent/user"
)

func (r *accountResolver) Owner(ctx context.Context, obj *model.Account) (*model.User, error) {
	e, err := r.DB.Ent.User.Query().
		Where(user.HasAccountsWith(account.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	m := &model.User{
		ID:           e.ID.String(),
		Name:         e.Name,
		Username:     e.Username,
		LastActivity: e.LastActivity,
		Picture: e.Picture,
	}

	return m, nil
}

func (r *answerResolver) Question(ctx context.Context, obj *model.Answer) (*model.Question, error) {
	e, err := r.DB.Ent.Question.Query().
		Where(question.HasAnswersWith(answer.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	m := &model.Question{
		ID:          e.ID.String(),
		Hash:        e.Hash,
		Title:       e.Title,
		Description: e.Description,
		Anonymous:   e.Anonymous,
		// Metadata:   (e.Metadata),
	}

	return m, nil
}

func (r *contactResolver) Owner(ctx context.Context, obj *model.Contact) (*model.User, error) {
	e, err := r.DB.Ent.User.Query().
		Where(user.HasContactsWith(contact.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	m := &model.User{
		ID:           e.ID.String(),
		Name:         e.Name,
		Username:     e.Username,
		LastActivity: e.LastActivity,
		Picture: e.Picture,
	}

	return m, nil
}

func (r *domainResolver) Surveys(ctx context.Context, obj *model.Domain) ([]*model.Survey, error) {
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
			aa := *a
			arr = append(arr, &model.Survey{
				ID:              aa.ID.String(),
				Tags:            aa.Tags,
				LastInteraction: aa.LastInteraction,
				DueDate:         aa.DueDate,
				Title:           aa.Title,
				Description:     aa.Description,
				// Metadata:        nil,
			})
		}
	}

	return arr, nil
}

func (r *domainResolver) Users(ctx context.Context, obj *model.Domain) ([]*model.User, error) {
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
			aa := *a
			arr = append(arr, &model.User{
				ID:           aa.ID.String(),
				Name:         aa.Name,
				Username:     aa.Username,
				LastActivity: aa.LastActivity,
			})
		}
	}

	return arr, nil
}

func (r *flowResolver) Questions(ctx context.Context, obj *model.Flow) ([]*model.Question, error) {
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
			aa := *a
			arr = append(arr, &model.Question{
				ID:          aa.ID.String(),
				Hash:        aa.Hash,
				Title:       aa.Title,
				Description: aa.Description,
				Anonymous:   aa.Anonymous,
				// Metadata:    nil,

			})
		}
	}

	return arr, nil
}

func (r *inputResolver) Options(ctx context.Context, obj *model.Input) (map[string]interface{}, error) {
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
	e, err := r.DB.Ent.Question.Query().
		Where(question.ID(uuid.MustParse(obj.ID))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	m := &model.Question{
		ID:          e.ID.String(),
		Hash:        e.Hash,
		Title:       e.Title,
		Description: e.Description,
		Anonymous:   e.Anonymous,
		// Metadata:    nil,
	}

	return m, nil
}

func (r *questionResolver) Answers(ctx context.Context, obj *model.Question) ([]*model.Answer, error) {
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
			aa := *a
			arr = append(arr, &model.Answer{
				ID:        aa.ID.String(),
				At:        aa.At,
				Responses: aa.Responses,
				Valid:     aa.Valid,
			})
		}
	}

	return arr, nil
}

func (r *questionResolver) Input(ctx context.Context, obj *model.Question) (*model.Input, error) {
	e, err := r.DB.Ent.Input.Query().
		Where(input.HasQuestionWith(question.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	defs := make([]*string, 0)
	for _, s := range e.Defaults {
		defs = append(defs, &s)
	}

	opts := map[string]interface{}{}

	for k, v := range e.Options {
		opts[k] = v
	}
	m := &model.Input{
		ID:       e.ID.String(),
		Kind:     e.Kind.String(),
		Multiple: e.Multiple,
		Defaults: defs,
		Options:  opts,
	}

	return m, nil
}

func (r *questionResolver) Flow(ctx context.Context, obj *model.Question) (*model.Flow, error) {
	e, err := r.DB.Ent.Flow.Query().
		Where(flow.HasQuestionsWith(question.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	m := &model.Flow{
		ID:         e.ID.String(),
		State:      e.State.String(),
		StateTable: e.StateTable,
		Inputs:     e.Inputs,
	}

	return m, nil
}

func (r *surveyResolver) Flow(ctx context.Context, obj *model.Survey) (*model.Flow, error) {
	e, err := r.DB.Ent.Survey.Query().
		Where(survey.ID(uuid.MustParse(obj.ID))).
		QueryFlow().
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	m := &model.Flow{
		ID:         e.ID.String(),
		State:      e.State.String(),
		StateTable: e.StateTable,
		Inputs:     e.Inputs,
	}

	return m, nil
}

func (r *surveyResolver) For(ctx context.Context, obj *model.Survey) (*model.User, error) {
	e, err := r.DB.Ent.User.Query().
		Where(user.HasSurveysWith(survey.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	m := &model.User{
		ID:           e.ID.String(),
		Name:         e.Name,
		Username:     e.Username,
		LastActivity: e.LastActivity,
		Picture: e.Picture,
	}

	return m, nil
}

func (r *surveyResolver) Owner(ctx context.Context, obj *model.Survey) (*model.Domain, error) {
	e, err := r.DB.Ent.Domain.Query().
		Where(domain.HasSurveysWith(survey.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	m := &model.Domain{
		ID:             e.ID.String(),
		Tags:           e.Tags,
		Name:           e.Name,
		Email:          e.Email,
		Domain:         e.Domain,
		CollectaDomain: e.CollectaDomain,
	}

	return m, nil
}

func (r *userResolver) Accounts(ctx context.Context, obj *model.User) (*model.Account, error) {
	e, err := r.DB.Ent.Account.Query().
		Where(account.HasOwnerWith(user.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	m := &model.Account{
		ID:       e.ID.String(),
		Type:     e.Type.String(),
		Sub:      e.Sub,
		RemoteID: e.RemoteID,
	}

	return m, nil
}

func (r *userResolver) Contacts(ctx context.Context, obj *model.User) (*model.Contact, error) {
	e, err := r.DB.Ent.Contact.Query().
		Where(contact.HasOwnerWith(user.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	m := &model.Contact{
		ID:          e.ID.String(),
		Name:        &e.Name,
		Value:       e.Value,
		Kind:        e.Kind.String(),
		Principal:   e.Principal,
		Validated:   e.Validated,
		FromAccount: e.FromAccount,
	}

	return m, nil
}

func (r *userResolver) Surveys(ctx context.Context, obj *model.User) ([]*model.Survey, error) {
	log.Info("at survey query")
	e, err := r.DB.Ent.User.Query().
		Where(user.ID(uuid.MustParse(obj.ID))).
		QuerySurveys().
		All(ctx)

	log.Info(spew.Sdump(e))

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	arr := make([]*model.Survey, 0)
	for _, a := range e {
		if a != nil {
			aa := *a
			arr = append(arr, &model.Survey{
				ID:              aa.ID.String(),
				Tags:            aa.Tags,
				LastInteraction: aa.LastInteraction,
				DueDate:         aa.DueDate,
				Title:           aa.Title,
				Description:     aa.Description,
			})
		}
	}

	log.Info(spew.Sdump(arr))
	
	return arr, nil
}

func (r *userResolver) Domain(ctx context.Context, obj *model.User) (*model.Domain, error) {
	e, err := r.DB.Ent.Domain.Query().
		Where(domain.HasUsersWith(user.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	m := &model.Domain{
		ID:     e.ID.String(),
		Tags:   e.Tags,
		Name:   e.Name,
		Email:  e.Email,
		Domain: e.Domain,
		CollectaDomain: e.CollectaDomain,
	}

	return m, nil
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
