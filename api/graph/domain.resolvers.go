package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/minskylab/collecta/errors"

	uuid1 "github.com/google/uuid"
	"github.com/minskylab/collecta/api/commons"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/domain"
)

func (r *domainResolver) Surveys(ctx context.Context, obj *ent.Domain) ([]*ent.Survey, error) {
	ownerResID, err := commons.OwnerOfDomain(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID...); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Domain.Query().
		Where(domain.ID(obj.ID)).
		QuerySurveys().
		All(ctx)
}

func (r *domainResolver) Users(ctx context.Context, obj *ent.Domain) ([]*ent.Person, error) {
	ownerResID, err := commons.OwnerOfDomain(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID...); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Domain.Query().
		Where(domain.ID(obj.ID)).
		QueryUsers().
		All(ctx)
}

func (r *domainResolver) Admins(ctx context.Context, obj *ent.Domain) ([]*ent.Person, error) {
	ownerResID, err := commons.OwnerOfDomain(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID...); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Domain.Query().
		Where(domain.ID(obj.ID)).
		QueryAdmins().
		All(ctx)
}

// Domain returns generated.DomainResolver implementation.
func (r *Resolver) Domain() generated.DomainResolver { return &domainResolver{r} }

type domainResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *domainResolver) ID(ctx context.Context, obj *ent.Domain) (uuid1.UUID, error) {
	panic(fmt.Errorf("not implemented"))
}
