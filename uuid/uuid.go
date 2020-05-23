package uuid

import (
	"database/sql/driver"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/errors"
)

type UUID [16]byte

var Nil UUID

func New() UUID {
	return UUID(uuid.New())
}

func (uid UUID) Value() (driver.Value, error) {
	return uuid.UUID(uid).String(), nil
}

func (uid UUID) String() string {
	return uuid.UUID(uid).String()
}

func MustParse(s string) UUID {
	return UUID(uuid.MustParse(s))
}
func Parse(s string) (UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return Nil, errors.Wrap(err, "error at try to parse collecta uuid")
	}

	return UUID(id), nil
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (uid *UUID) UnmarshalGQL(v interface{}) error {
	sID, ok := v.(string)
	if !ok {
		return fmt.Errorf("points must be strings")
	}

	genUID, err := uuid.Parse(sID)
	if err != nil {
		return errors.Wrap(err, "invalid uuid, cannot be parsed")
	}
	*uid = UUID(genUID)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (uid UUID) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(uuid.UUID(uid).String()))
}