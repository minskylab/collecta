package auth

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/markbates/goth"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/person"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func (collectaAuth *Auth) ingressWithGoogle(ctx context.Context, rawUser goth.User) (string, error) {
	spew.Dump(rawUser)
	googleUserExist, err := collectaAuth.db.Ent.Person.Query().
		Where(person.HasAccountsWith(
			account.And(
				account.Sub(rawUser.Email),
				account.RemoteID(rawUser.UserID),
			)),
		).
		Exist(ctx)

	if err != nil {
		return "", errors.Wrap(err, "error at try to verify the user existance")
	}

	var googleUser *ent.Person

	if !googleUserExist {
		googleUser, err = collectaAuth.registerNewUserFromGoogle(ctx, rawUser)
		if err != nil {
			return "", errors.Wrap(err, "error at try to register new google user")
		}
	} else {
		googleUser, err = collectaAuth.db.Ent.Person.Query().
			Where(person.HasAccountsWith(
				account.And(
					account.Sub(rawUser.Email),
					account.RemoteID(rawUser.UserID),
				)),
			).
			Only(ctx)
		if err != nil {
			return "", errors.Wrap(err, "error at try to fetch user ent")
		}
	}

	return collectaAuth.createJWTToken(googleUser)
}
