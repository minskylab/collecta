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
	"github.com/minskylab/collecta/ent/contact"
	"github.com/minskylab/collecta/ent/person"
)

func (r *contactResolver) Owner(ctx context.Context, obj *model.Contact) (*model.User, error) {
	ownerResID, err := commons.OwnerOfContact(ctx, r.DB, obj)
	if err != nil {
		return nil, errors.Wrap(err, "error at extract owner of resource")
	}

	if err := validateAuthorization(ctx, r.Auth, ownerResID); err != nil {
		return nil, errors.Wrap(err, "error at validate your credentials")
	}

	e, err := r.DB.Ent.Person.Query().
		Where(person.HasContactsWith(contact.ID(uuid.MustParse(obj.ID)))).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at ent query")
	}

	return commons.PersonToGQL(e), nil
}

// Contact returns generated.ContactResolver implementation.
func (r *Resolver) Contact() generated.ContactResolver { return &contactResolver{r} }

type contactResolver struct{ *Resolver }
