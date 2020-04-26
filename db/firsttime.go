package db

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/config"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/user"
	"github.com/minskylab/collecta/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

func (db *DB) generateFirstAdminUser(ctx context.Context) (*ent.User, error) {
	password := strings.TrimSpace(viper.GetString(config.FirstAdminPassword))
	if password == "" {
		log.Warn("your password wasn't found, collecta set a generic admin password, you may to update that early")
		password = "admin"
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to crypt the admin password")
	}

	newAdmin, err := db.Ent.User.Create().
		SetID(uuid.UUID{}).
		SetName("Admin").
		SetUsername("admin").
		SetLastActivity(time.Now()).
		SetPicture("https://via.placeholder.com/150").
		SetRoles([]string{"admin", "creator"}).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at create new admin user")
	}

	_, err = db.Ent.Account.Create().
		SetID(uuid.New()).
		SetRemoteID("admin").
		SetSecret(string(hash)).
		SetOwner(newAdmin).
		SetSub("admin").
		SetType(account.TypeEmail).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error ar try to create a new admin account")
	}

	return newAdmin, nil
}

// isFirstTimeCollectaInstance determine if collecta instance is a first time execution based on the
// user admin existence
func (db *DB) isFirstTimeCollectaInstance(ctx context.Context) (bool, error) {
	adminUserExist, err := db.Ent.User.Query().
		Where(user.HasAccountsWith(account.And(account.RemoteID("admin"), account.Sub("admin")))).
		Exist(ctx)
	if err != nil {
		return false, errors.Wrap(err, "error at fetch admin user existence")
	}
	return !adminUserExist, nil
}

