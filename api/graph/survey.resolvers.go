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
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/ent/person"
)

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

	e, err := r.DB.Ent.Person.Query().
		Where(person.HasSurveysWith(survey.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.PersonToGQL(e), nil
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

// Survey returns generated.SurveyResolver implementation.
func (r *Resolver) Survey() generated.SurveyResolver { return &surveyResolver{r} }

type surveyResolver struct{ *Resolver }
