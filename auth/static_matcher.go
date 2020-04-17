package auth

import (
	"context"

	"github.com/markbates/goth"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/pkg/errors"
)

func (collectaAuth *CollectaAuth) matchGoogleUserWithCollectaDomain(ctx context.Context, rawUser goth.User) (string, error) {
	domainHost, ok := rawUser.RawData["hd"].(string)
	if !ok {
		return "", errors.New("invalid domain in raw rawUser data")
	}

	domainExists, err := collectaAuth.db.Ent.Domain.Query().
		Where(domain.Domain(domainHost)).
		Exist(ctx)
	if err != nil {
		return "", errors.Wrap(err, "error at fetch domain")
	}

	if !domainExists {
		return "", errors.New("domain of rawUser not exist")
	}

	dom, err := collectaAuth.db.Ent.Domain.Query().
		Where(domain.Domain(domainHost)).
		Only(ctx)

	if err != nil {
		return "", errors.Wrap(err, "error at fetch domain")
	}

	return dom.CollectaDomain, nil
}
