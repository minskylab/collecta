package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/minskylab/collecta/errors"

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
