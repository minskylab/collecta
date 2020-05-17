package auth

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/markbates/goth"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/contact"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func (collectaAuth *Auth) registerNewUserFromGoogle(ctx context.Context, rawUser goth.User) (*ent.Person, error) {
	name := rawUser.Name
	domainHost, ok := rawUser.RawData["hd"].(string)
	if !ok {
		parts := strings.Split(rawUser.Email, "@")
		if len(parts) != 2 {
			return nil, errors.New("invalid domain in raw rawUser data")
		}
		domainHost = strings.TrimSpace(parts[1])
		if name == "" {
			name = parts[0]
		}
	}


	domainExists, err := collectaAuth.db.Ent.Domain.Query().
		Where(domain.Domain(domainHost)).
		Exist(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at fetch domain")
	}

	if !domainExists {
		return nil, errors.New("domain of rawUser not exist")
	}

	userDomain, err := collectaAuth.db.Ent.Domain.Query().Where(domain.Domain(domainHost)).Only(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to get the domain id")
	}

	newUser, err := collectaAuth.db.Ent.Person.Create().
		SetID(uuid.New()).
		SetName(name).
		SetUsername(name).
		SetLastActivity(time.Now()).
		SetPicture(rawUser.AvatarURL).
		SetRoles([]string{"user"}).
		AddDomains(userDomain).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at create new user")
	}

	_, err = collectaAuth.db.Ent.Account.Create().
		SetID(uuid.New()).
		SetType(account.TypeGoogle).
		SetSub(rawUser.Email).
		SetRemoteID(rawUser.UserID).
		SetOwner(newUser).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at create new account")
	}

	_, err = collectaAuth.db.Ent.Contact.Create().
		SetID(uuid.New()).
		SetName("Google Email").
		SetKind(contact.KindEmail).
		SetPrincipal(true).
		SetValidated(false).
		SetValue(rawUser.Email).
		SetFromAccount(true).
		SetOwner(newUser).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at create new contact")
	}

	log.Info("generating demo survey")
	// TODO: refactoring
	// surv, err := drafts.GenerateUTECDemo(ctx, collectaAuth.db, userDomain.ID, newUser.ID)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "error at utec demo generator")
	// }

	// log.WithField("survey", surv).Info("demo generated ")

	return newUser, nil
}
