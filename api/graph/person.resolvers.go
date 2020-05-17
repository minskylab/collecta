package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/minskylab/collecta/errors"

	"github.com/minskylab/collecta/api/commons"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/contact"
	"github.com/minskylab/collecta/ent/person"
)

func (r *personResolver) Accounts(ctx context.Context, obj *ent.Person) ([]*ent.Account, error) {
	ownerResID, err := commons.OwnerOfPerson(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Account.Query().
		Where(account.HasOwnerWith(person.ID(obj.ID))).
		All(ctx)
}

func (r *personResolver) Contacts(ctx context.Context, obj *ent.Person) ([]*ent.Contact, error) {
	ownerResID, err := commons.OwnerOfPerson(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Contact.Query().
		Where(contact.HasOwnerWith(person.ID(obj.ID))).
		All(ctx)
}

func (r *personResolver) Surveys(ctx context.Context, obj *ent.Person) ([]*ent.Survey, error) {
	ownerResID, err := commons.OwnerOfPerson(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	u, err := r.DB.Ent.Person.Get(ctx, obj.ID)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch user by user id")
	}

	return r.DB.Ent.Person.QuerySurveys(u).All(ctx)
}

func (r *personResolver) Domains(ctx context.Context, obj *ent.Person) ([]*ent.Domain, error) {
	ownerResID, err := commons.OwnerOfPerson(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Person.Query().
		Where(person.ID(obj.ID)).
		QueryDomains().
		All(ctx)
}

func (r *personResolver) AdminOf(ctx context.Context, obj *ent.Person) ([]*ent.Domain, error) {
	ownerResID, err := commons.OwnerOfPerson(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Person.Query().
		Where(person.ID(obj.ID)).
		QueryAdminOf().
		All(ctx)
}

// Person returns generated.PersonResolver implementation.
func (r *Resolver) Person() generated.PersonResolver { return &personResolver{r} }

type personResolver struct{ *Resolver }
