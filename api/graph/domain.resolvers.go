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
	"github.com/minskylab/collecta/ent/domain"
)

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
			arr = append(arr, commons.PersonToGQL(a))
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
			arr = append(arr, commons.PersonToGQL(a))
		}
	}

	return arr, nil
}

// Domain returns generated.DomainResolver implementation.
func (r *Resolver) Domain() generated.DomainResolver { return &domainResolver{r} }

type domainResolver struct{ *Resolver }
