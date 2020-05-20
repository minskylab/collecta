package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/minskylab/collecta/auth"
	"github.com/minskylab/collecta/errors"
	"github.com/minskylab/collecta/uuid"
)

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func validateAuthorization(ctx context.Context, auth *auth.Auth, ownerResources ...uuid.UUID) error {
	userRequester := auth.UserOfContext(ctx)
	if userRequester == nil {
		return errors.New("unauthorized, please include a valid token in your header")
	}

	isOwner := false

	if strings.Contains(strings.Join(userRequester.Roles, " "), "admin") { // is admin
		return nil
	}

	for _, ownerResource := range ownerResources {
		if ownerResource == userRequester.ID {
			isOwner = true
			break
		}
	}

	if !isOwner {
		return errors.New("resource unavailable for you")
	}

	return nil
}
