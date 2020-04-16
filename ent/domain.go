// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/rs/xid"
)

// Domain is the model entity for the Domain schema.
type Domain struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags []string `json:"tags,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Domain holds the value of the "domain" field.
	Domain string `json:"domain,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DomainQuery when eager-loading is set.
	Edges DomainEdges `json:"edges"`
}

// DomainEdges holds the relations/edges for other nodes in the graph.
type DomainEdges struct {
	// Surveys holds the value of the surveys edge.
	Surveys []*Survey
	// Users holds the value of the users edge.
	Users []*User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SurveysOrErr returns the Surveys value or an error if the edge
// was not loaded in eager-loading.
func (e DomainEdges) SurveysOrErr() ([]*Survey, error) {
	if e.loadedTypes[0] {
		return e.Surveys, nil
	}
	return nil, &NotLoadedError{edge: "surveys"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e DomainEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Domain) scanValues() []interface{} {
	return []interface{}{
		&xid.ID{},         // id
		&[]byte{},         // tags
		&sql.NullString{}, // name
		&sql.NullString{}, // email
		&sql.NullString{}, // domain
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Domain fields.
func (d *Domain) assignValues(values ...interface{}) error {
	if m, n := len(values), len(domain.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*xid.ID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		d.ID = *value
	}
	values = values[1:]

	if value, ok := values[0].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field tags", values[0])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &d.Tags); err != nil {
			return fmt.Errorf("unmarshal field tags: %v", err)
		}
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		d.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[2])
	} else if value.Valid {
		d.Email = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field domain", values[3])
	} else if value.Valid {
		d.Domain = value.String
	}
	return nil
}

// QuerySurveys queries the surveys edge of the Domain.
func (d *Domain) QuerySurveys() *SurveyQuery {
	return (&DomainClient{config: d.config}).QuerySurveys(d)
}

// QueryUsers queries the users edge of the Domain.
func (d *Domain) QueryUsers() *UserQuery {
	return (&DomainClient{config: d.config}).QueryUsers(d)
}

// Update returns a builder for updating this Domain.
// Note that, you need to call Domain.Unwrap() before calling this method, if this Domain
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Domain) Update() *DomainUpdateOne {
	return (&DomainClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Domain) Unwrap() *Domain {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Domain is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Domain) String() string {
	var builder strings.Builder
	builder.WriteString("Domain(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", tags=")
	builder.WriteString(fmt.Sprintf("%v", d.Tags))
	builder.WriteString(", name=")
	builder.WriteString(d.Name)
	builder.WriteString(", email=")
	builder.WriteString(d.Email)
	builder.WriteString(", domain=")
	builder.WriteString(d.Domain)
	builder.WriteByte(')')
	return builder.String()
}

// Domains is a parsable slice of Domain.
type Domains []*Domain

func (d Domains) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
