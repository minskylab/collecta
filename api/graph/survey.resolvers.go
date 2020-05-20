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
	"github.com/minskylab/collecta/ent/person"
	"github.com/minskylab/collecta/ent/survey"
)

func (r *surveyResolver) Flow(ctx context.Context, obj *ent.Survey) (*ent.Flow, error) {
	ownerResID, err := commons.OwnerOfSurvey(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Survey.Query().
		Where(survey.ID(obj.ID)).
		QueryFlow().
		Only(ctx)
}

func (r *surveyResolver) For(ctx context.Context, obj *ent.Survey) (*ent.Person, error) {
	ownerResID, err := commons.OwnerOfSurvey(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Person.Query().
		Where(person.HasSurveysWith(survey.ID(obj.ID))).
		Only(ctx)
}

func (r *surveyResolver) Owner(ctx context.Context, obj *ent.Survey) (*ent.Domain, error) {
	ownerResID, err := commons.OwnerOfSurvey(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	return r.DB.Ent.Domain.Query().
		Where(domain.HasSurveysWith(survey.ID(obj.ID))).
		Only(ctx)
}

// Survey returns generated.SurveyResolver implementation.
func (r *Resolver) Survey() generated.SurveyResolver { return &surveyResolver{r} }

type surveyResolver struct{ *Resolver }