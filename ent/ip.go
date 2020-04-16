// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/minskylab/collecta/ent/ip"
)

// IP is the model entity for the IP schema.
type IP struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*IP) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the IP fields.
func (i *IP) assignValues(values ...interface{}) error {
	if m, n := len(values), len(ip.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	i.ID = int(value.Int64)
	values = values[1:]
	return nil
}

// Update returns a builder for updating this IP.
// Note that, you need to call IP.Unwrap() before calling this method, if this IP
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *IP) Update() *IPUpdateOne {
	return (&IPClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (i *IP) Unwrap() *IP {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: IP is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *IP) String() string {
	var builder strings.Builder
	builder.WriteString("IP(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteByte(')')
	return builder.String()
}

// IPs is a parsable slice of IP.
type IPs []*IP

func (i IPs) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
