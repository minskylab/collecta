package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"


	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/commons"
	"github.com/minskylab/collecta/api/graph/generated"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/person"
	"github.com/minskylab/collecta/errors"
)

func (r *accountResolver) Owner(ctx context.Context, obj *model.Account) (*model.User, error) {
	ownerResID, err := commons.OwnerOfAccount(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Person.Query().
		Where(person.HasAccountsWith(account.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.PersonToGQL(e), nil
}

// Account returns generated.AccountResolver implementation.
func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

type accountResolver struct{ *Resolver }
