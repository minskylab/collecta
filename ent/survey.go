// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/person"
	"github.com/minskylab/collecta/ent/survey"
)

// Survey is the model entity for the Survey schema.
type Survey struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags []string `json:"tags,omitempty"`
	// LastInteraction holds the value of the "lastInteraction" field.
	LastInteraction time.Time `json:"lastInteraction,omitempty"`
	// DueDate holds the value of the "dueDate" field.
	DueDate time.Time `json:"dueDate,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Done holds the value of the "done" field.
	Done bool `json:"done,omitempty"`
	// IsPublic holds the value of the "isPublic" field.
	IsPublic bool `json:"isPublic,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SurveyQuery when eager-loading is set.
	Edges          SurveyEdges `json:"edges"`
	domain_surveys *uuid.UUID
	person_surveys *uuid.UUID
}

// SurveyEdges holds the relations/edges for other nodes in the graph.
type SurveyEdges struct {
	// Flow holds the value of the flow edge.
	Flow *Flow
	// For holds the value of the for edge.
	For *Person
	// Owner holds the value of the owner edge.
	Owner *Domain
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// FlowOrErr returns the Flow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SurveyEdges) FlowOrErr() (*Flow, error) {
	if e.loadedTypes[0] {
		if e.Flow == nil {
			// The edge flow was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: flow.Label}
		}
		return e.Flow, nil
	}
	return nil, &NotLoadedError{edge: "flow"}
}

// ForOrErr returns the For value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SurveyEdges) ForOrErr() (*Person, error) {
	if e.loadedTypes[1] {
		if e.For == nil {
			// The edge for was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: person.Label}
		}
		return e.For, nil
	}
	return nil, &NotLoadedError{edge: "for"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SurveyEdges) OwnerOrErr() (*Domain, error) {
	if e.loadedTypes[2] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: domain.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Survey) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},      // id
		&[]byte{},         // tags
		&sql.NullTime{},   // lastInteraction
		&sql.NullTime{},   // dueDate
		&sql.NullString{}, // title
		&sql.NullString{}, // description
		&[]byte{},         // metadata
		&sql.NullBool{},   // done
		&sql.NullBool{},   // isPublic
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Survey) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // domain_surveys
		&uuid.UUID{}, // person_surveys
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Survey fields.
func (s *Survey) assignValues(values ...interface{}) error {
	if m, n := len(values), len(survey.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		s.ID = *value
	}
	values = values[1:]

	if value, ok := values[0].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field tags", values[0])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &s.Tags); err != nil {
			return fmt.Errorf("unmarshal field tags: %v", err)
		}
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field lastInteraction", values[1])
	} else if value.Valid {
		s.LastInteraction = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field dueDate", values[2])
	} else if value.Valid {
		s.DueDate = value.Time
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field title", values[3])
	} else if value.Valid {
		s.Title = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[4])
	} else if value.Valid {
		s.Description = value.String
	}

	if value, ok := values[5].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field metadata", values[5])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &s.Metadata); err != nil {
			return fmt.Errorf("unmarshal field metadata: %v", err)
		}
	}
	if value, ok := values[6].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field done", values[6])
	} else if value.Valid {
		s.Done = value.Bool
	}
	if value, ok := values[7].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field isPublic", values[7])
	} else if value.Valid {
		s.IsPublic = value.Bool
	}
	values = values[8:]
	if len(values) == len(survey.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field domain_surveys", values[0])
		} else if value != nil {
			s.domain_surveys = value
		}
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field person_surveys", values[0])
		} else if value != nil {
			s.person_surveys = value
		}
	}
	return nil
}

// QueryFlow queries the flow edge of the Survey.
func (s *Survey) QueryFlow() *FlowQuery {
	return (&SurveyClient{config: s.config}).QueryFlow(s)
}

// QueryFor queries the for edge of the Survey.
func (s *Survey) QueryFor() *PersonQuery {
	return (&SurveyClient{config: s.config}).QueryFor(s)
}

// QueryOwner queries the owner edge of the Survey.
func (s *Survey) QueryOwner() *DomainQuery {
	return (&SurveyClient{config: s.config}).QueryOwner(s)
}

// Update returns a builder for updating this Survey.
// Note that, you need to call Survey.Unwrap() before calling this method, if this Survey
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Survey) Update() *SurveyUpdateOne {
	return (&SurveyClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Survey) Unwrap() *Survey {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Survey is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Survey) String() string {
	var builder strings.Builder
	builder.WriteString("Survey(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", tags=")
	builder.WriteString(fmt.Sprintf("%v", s.Tags))
	builder.WriteString(", lastInteraction=")
	builder.WriteString(s.LastInteraction.Format(time.ANSIC))
	builder.WriteString(", dueDate=")
	builder.WriteString(s.DueDate.Format(time.ANSIC))
	builder.WriteString(", title=")
	builder.WriteString(s.Title)
	builder.WriteString(", description=")
	builder.WriteString(s.Description)
	builder.WriteString(", metadata=")
	builder.WriteString(fmt.Sprintf("%v", s.Metadata))
	builder.WriteString(", done=")
	builder.WriteString(fmt.Sprintf("%v", s.Done))
	builder.WriteString(", isPublic=")
	builder.WriteString(fmt.Sprintf("%v", s.IsPublic))
	builder.WriteByte(')')
	return builder.String()
}

// Surveys is a parsable slice of Survey.
type Surveys []*Survey

func (s Surveys) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
