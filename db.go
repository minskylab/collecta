package collecta

import (
	"context"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/minskylab/collecta/ent"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

func openDBConnection() {
	drv, err := sql.Open("postgres", "<mysql-dsn>")
	if err != nil {
		panic(errors.Wrap(err, "error at open sql connection"))
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	client := ent.NewClient(ent.Driver(drv))

	user, err := client.User.Create().
		SetID(xid.New()).
		SetName("Bregy Malpartida").
		Save(context.Background())
	if err != nil {
		panic(errors.Wrap(err, "error at try to create new user with mininmal params"))
	}

	spew.Dump(user)
}
