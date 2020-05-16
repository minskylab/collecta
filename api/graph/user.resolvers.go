package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/minskylab/collecta/errors"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/commons"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/contact"
	"github.com/minskylab/collecta/ent/person"
)

func (r *userResolver) Accounts(ctx context.Context, obj *model.User) ([]*model.Account, error) {
	ownerResID, err := commons.OwnerOfUser(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Account.Query().
		Where(account.HasOwnerWith(person.ID(uuid.MustParse(obj.ID)))).
		All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	arr := make([]*model.Account, 0)
	for _, a := range e {
		if a != nil {
			arr = append(arr, commons.AccountToGQL(a))
		}
	}

	return arr, nil
}

func (r *userResolver) Contacts(ctx context.Context, obj *model.User) ([]*model.Contact, error) {
	ownerResID, err := commons.OwnerOfUser(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Contact.Query().
		Where(contact.HasOwnerWith(person.ID(uuid.MustParse(obj.ID)))).
		All(ctx)

	arr := make([]*model.Contact, 0)
	for _, a := range e {
		if a != nil {
			arr = append(arr, commons.ContactToGQL(a))
		}
	}

	return arr, nil
}

func (r *userResolver) Surveys(ctx context.Context, obj *model.User) ([]*model.Survey, error) {
	ownerResID, err := commons.OwnerOfUser(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	u, err := r.DB.Ent.Person.Get(ctx, uuid.MustParse(obj.ID))
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch user by user id")
	}

	e, err := r.DB.Ent.Person.QuerySurveys(u).All(ctx)

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

	e, err := r.DB.Ent.Person.Query().
		Where(person.ID(uuid.MustParse(obj.ID))).
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

	e, err := r.DB.Ent.Person.Query().
		Where(person.ID(uuid.MustParse(obj.ID))).
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

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
