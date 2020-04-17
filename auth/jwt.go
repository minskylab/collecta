package auth

import (
	"context"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/markbates/goth"
	"github.com/minskylab/collecta/drafts"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/contact"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/user"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (collectaAuth *CollectaAuth) ingressWithGoogle(ctx context.Context, rawUser goth.User) (string, error) {
	spew.Dump(rawUser)
	googleUserExist, err := collectaAuth.db.Ent.User.Query().
		Where(user.HasAccountsWith(
			account.And(
				account.Sub(rawUser.Email),
				account.RemoteID(rawUser.UserID),
			)),
		).
		Exist(ctx)

	if err != nil {
		return "", errors.Wrap(err, "error at try to verify the user existance")
	}

	var googleUser *ent.User

	if !googleUserExist {
		googleUser, err = collectaAuth.registerNewUserFromGoogle(ctx, rawUser)
		if err != nil {
			return "", errors.Wrap(err, "error at try to register new google user")
		}
	} else {
		googleUser, err = collectaAuth.db.Ent.User.Query().
			Where(user.HasAccountsWith(
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

func (collectaAuth *CollectaAuth) registerNewUserFromGoogle(ctx context.Context, rawUser goth.User) (*ent.User, error) {
	domainHost, ok := rawUser.RawData["hd"].(string)
	if !ok {
		return nil, errors.New("invalid domain in raw rawUser data")
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

	newUser, err := collectaAuth.db.Ent.User.Create().
		SetID(uuid.New()).
		SetName(rawUser.Name).
		SetDomain(userDomain).
		SetUsername(rawUser.Name).
		SetLastActivity(time.Now()).
		SetPicture(rawUser.AvatarURL).
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
	surv, err := drafts.GenerateUTECDemo(ctx, collectaAuth.db, userDomain.ID, newUser.ID)
	if err != nil {
		return nil, errors.Wrap(err, "error at utec demo generator")
	}

	log.WithField("survey", surv).Info("demo generated ")

	return newUser, nil
}

func (collectaAuth *CollectaAuth) createJWTToken(u *ent.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  "user",
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		// Id:        u.ID.String(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "collecta",
		Subject:   u.ID.String(),
	})

	t, err := token.SignedString([]byte("collecta.temporal.stupid.secret.key"))
	if err != nil {
		return "", errors.Wrap(err, "error caused by jwt SignedString method")
	}

	return t, nil
}


func (collectaAuth *CollectaAuth) verifyJWTToken(ctx context.Context, tokenString string) (*ent.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New( "unexpected signing method")
		}
		return []byte("collecta.temporal.stupid.secret.key"), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "error caused by jwt SignedString method")
	}

	spew.Dump(token)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uID, ok := claims["sub"].(string)
		if ok {
			return collectaAuth.db.Ent.User.Get(ctx, uuid.MustParse(uID))
		}
	}

	return nil, errors.New("invalid token claims")
}

func (collectaAuth *CollectaAuth) VerifyJWTToken(ctx context.Context, token string) (*ent.User, error) {
	return collectaAuth.verifyJWTToken(ctx, token)
}