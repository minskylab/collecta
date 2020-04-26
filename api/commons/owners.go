package commons

import (
	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/db"
	"github.com/minskylab/collecta/ent/answer"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/ent/question"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/errors"
	"golang.org/x/net/context"
)

func OwnerOfDomain(ctx context.Context, db *db.DB, obj *model.Domain) ([]uuid.UUID, error) {
	d, err := db.Ent.Domain.Get(ctx, uuid.MustParse(obj.ID))
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch domain")
	}

	return d.QueryAdmins().IDs(ctx)
}

func OwnerOfUser(ctx context.Context, db *db.DB, obj *model.User) (uuid.UUID, error) {
	return uuid.Parse(obj.ID)
}

func OwnerOfAccount(ctx context.Context, db *db.DB, obj *model.Account) (uuid.UUID, error) {
	a, err := db.Ent.Account.Get(ctx, uuid.MustParse(obj.ID))
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch domain")
	}

	return a.QueryOwner().OnlyID(ctx)
}

func OwnerOfContact(ctx context.Context, db *db.DB, obj *model.Contact) (uuid.UUID, error) {
	c, err := db.Ent.Contact.Get(ctx, uuid.MustParse(obj.ID))
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "error at fetch domain")
	}

	return c.QueryOwner().OnlyID(ctx)
}

func OwnerOfFlow(ctx context.Context, db *db.DB, obj *model.Flow) (uuid.UUID, error) {
	return db.Ent.Survey.Query().Where(survey.HasFlowWith(flow.ID(uuid.MustParse(obj.ID)))).QueryFor().OnlyID(ctx)
}

func OwnerOfQuestion(ctx context.Context, db *db.DB, obj *model.Question) (uuid.UUID, error) {
	return db.Ent.Survey.Query().
		Where(survey.HasFlowWith(flow.HasQuestionsWith(question.ID(uuid.MustParse(obj.ID))))).
		QueryFor().OnlyID(ctx)
}

func OwnerOfAnswer(ctx context.Context, db *db.DB, obj *model.Answer) (uuid.UUID, error) {
	return db.Ent.Survey.Query().
		Where(survey.HasFlowWith(flow.HasQuestionsWith(question.HasAnswersWith(answer.ID(uuid.MustParse(obj.ID)))))).
		QueryFor().OnlyID(ctx)
}

func OwnerOfInput(ctx context.Context, db *db.DB, obj *model.Input) (uuid.UUID, error) {
	return db.Ent.Survey.Query().
		Where(
			survey.HasFlowWith(flow.HasQuestionsWith(question.HasInputWith(input.ID(uuid.MustParse(obj.ID)))))).
		QueryFor().OnlyID(ctx)
}

func OwnerOfSurvey(ctx context.Context, db *db.DB, obj *model.Survey) (uuid.UUID, error) {
	return db.Ent.Survey.Query().Where(survey.ID(uuid.MustParse(obj.ID))).QueryFor().OnlyID(ctx)
}