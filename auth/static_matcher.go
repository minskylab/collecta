package auth

import (
	"context"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/markbates/goth"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/pkg/errors"
)

func (collectaAuth *Auth) matchGoogleUserWithCollectaDomain(ctx context.Context, rawUser goth.User) (string, error) {
	spew.Dump(rawUser)

	domainHost, ok := rawUser.RawData["hd"].(string)
	if !ok {
		parts := strings.Split(rawUser.Email, "@")
		if len(parts) != 2 {
			return "", errors.New("invalid domain in raw rawUser data")
		}
		domainHost = strings.TrimSpace(parts[1])
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

	return dom.Callback, nil
}
