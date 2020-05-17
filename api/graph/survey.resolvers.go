package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/minskylab/collecta/errors"

	"fmt"

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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *surveyResolver) Metadata(ctx context.Context, obj *ent.Survey) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}
