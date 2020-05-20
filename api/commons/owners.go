package commons

import (
	"github.com/minskylab/collecta/db"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/answer"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/ent/person"
	"github.com/minskylab/collecta/ent/question"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/errors"
	"github.com/minskylab/collecta/uuid"
	"golang.org/x/net/context"
)

func OwnerOfDomain(ctx context.Context, db *db.DB, obj *ent.Domain) ([]uuid.UUID, error) {
	d, err := db.Ent.Domain.Get(ctx, obj.ID)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch domain")
	}

	admins, err := d.QueryAdmins().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch your users")
	}

	ids := make([]uuid.UUID, 0)
	for _, u := range admins {
		ids = append(ids, u.ID)
	}

	return ids, nil
}

func OwnerOfPerson(ctx context.Context, db *db.DB, obj *ent.Person) (uuid.UUID, error) {
	return obj.ID, nil
}

func OwnerOfAccount(ctx context.Context, db *db.DB, obj *ent.Account) (uuid.UUID, error) {
	a, err := db.Ent.Account.Get(ctx, obj.ID)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch domain")
	}

	u, err := a.QueryOwner().Only(ctx)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch your user")
	}
	return u.ID, nil
}

func OwnerOfContact(ctx context.Context, db *db.DB, obj *ent.Contact) (uuid.UUID, error) {
	c, err := db.Ent.Contact.Get(ctx, obj.ID)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch domain")
	}

	u, err := c.QueryOwner().Only(ctx)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch your user")
	}
	return u.ID, nil
}

func OwnerOfFlow(ctx context.Context, db *db.DB, obj *ent.Flow) (uuid.UUID, error) {
	u, err := db.Ent.Survey.Query().Where(survey.HasFlowWith(flow.ID(obj.ID))).QueryFor().Only(ctx)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch your user")
	}
	return u.ID, nil
}

func OwnerOfQuestion(ctx context.Context, db *db.DB, obj *ent.Question) (uuid.UUID, error) {
	u, err := db.Ent.Survey.Query().
		Where(survey.HasFlowWith(flow.HasQuestionsWith(question.ID(obj.ID)))).
		QueryFor().Only(ctx)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch your user")
	}
	return u.ID, nil
}

func OwnerOfAnswer(ctx context.Context, db *db.DB, obj *ent.Answer) (uuid.UUID, error) {
	u, err := db.Ent.Survey.Query().
		Where(survey.HasFlowWith(flow.HasQuestionsWith(question.HasAnswersWith(answer.ID(obj.ID))))).
		QueryFor().Only(ctx)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch your user")
	}
	return u.ID, nil
}

func OwnerOfInput(ctx context.Context, db *db.DB, obj *ent.Input) (uuid.UUID, error) {
	u, err := db.Ent.Survey.Query().
		Where(
			survey.HasFlowWith(flow.HasQuestionsWith(question.HasInputWith(input.ID(obj.ID))))).
		QueryFor().Only(ctx)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch your user")
	}
	return u.ID, nil
}

func OwnerOfSurvey(ctx context.Context, db *db.DB, obj *ent.Survey) (uuid.UUID, error) {
	u, err := db.Ent.Person.Query().Where(person.HasSurveysWith(survey.ID(obj.ID))).Only(ctx)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch your user")
	}
	return u.ID, nil
}