package auth

import (
	"context"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/minskylab/collecta/ent"
	"github.com/pkg/errors"
)



func (collectaAuth *Auth) createJWTToken(u *ent.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  "user",
		ExpiresAt: time.Now().Add(collectaAuth.jwtDuration).Unix(),
		// Id:        u.ID.String(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "collecta",
		Subject:   u.ID.String(),
	})

	t, err := token.SignedString(collectaAuth.jwtSecret)
	if err != nil {
		return "", errors.Wrap(err, "error caused by jwt SignedString method")
	}

	return t, nil
}


func (collectaAuth *Auth) verifyJWTToken(ctx context.Context, tokenString string) (*ent.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New( "unexpected signing method")
		}
		return collectaAuth.jwtSecret, nil
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

func (collectaAuth *Auth) VerifyJWTToken(ctx context.Context, token string) (*ent.User, error) {
	return collectaAuth.verifyJWTToken(ctx, token)
}