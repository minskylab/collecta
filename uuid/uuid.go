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

func (id UUID) Value() (driver.Value, error) {
	return uuid.UUID(id).String(), nil
}

func (id UUID) String() string {
	return uuid.UUID(id).String()
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
func (id *UUID) UnmarshalGQL(v interface{}) error {
	sID, ok := v.(string)
	if !ok {
		return fmt.Errorf("points must be strings")
	}

	genUID, err := uuid.Parse(sID)
	if err != nil {
		return errors.Wrap(err, "invalid uuid, cannot be parsed")
	}

	*id = UUID(genUID)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (id UUID) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(uuid.UUID(id).String()))
}

func (id *UUID) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil

	case string:
		// if an empty UUID comes from a table, we return a null UUID
		if src == "" {
			return nil
		}

		// see Parse for required string format
		u, err := Parse(src)
		if err != nil {
			return fmt.Errorf("Scan: %v", err)
		}

		*id = u

	case []byte:
		// if an empty UUID comes from a table, we return a null UUID
		if len(src) == 0 {
			return nil
		}

		// assumes a simple slice of bytes if 16 bytes
		// otherwise attempts to parse
		if len(src) != 16 {
			return id.Scan(string(src))
		}
		copy((*id)[:], src)

	default:
		return fmt.Errorf("Scan: unable to scan type %T into UUID", src)
	}

	return nil
}
