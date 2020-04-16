// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/contact"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/ent/user"
	"github.com/rs/xid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	id           *xid.ID
	name         *string
	username     *string
	lastActivity *time.Time
	accounts     map[xid.ID]struct{}
	contacts     map[xid.ID]struct{}
	surveys      map[xid.ID]struct{}
	domain       map[xid.ID]struct{}
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.name = &s
	return uc
}

// SetUsername sets the username field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.username = &s
	return uc
}

// SetNillableUsername sets the username field if the given value is not nil.
func (uc *UserCreate) SetNillableUsername(s *string) *UserCreate {
	if s != nil {
		uc.SetUsername(*s)
	}
	return uc
}

// SetLastActivity sets the lastActivity field.
func (uc *UserCreate) SetLastActivity(t time.Time) *UserCreate {
	uc.lastActivity = &t
	return uc
}

// SetNillableLastActivity sets the lastActivity field if the given value is not nil.
func (uc *UserCreate) SetNillableLastActivity(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetLastActivity(*t)
	}
	return uc
}

// SetID sets the id field.
func (uc *UserCreate) SetID(x xid.ID) *UserCreate {
	uc.id = &x
	return uc
}

// AddAccountIDs adds the accounts edge to Account by ids.
func (uc *UserCreate) AddAccountIDs(ids ...xid.ID) *UserCreate {
	if uc.accounts == nil {
		uc.accounts = make(map[xid.ID]struct{})
	}
	for i := range ids {
		uc.accounts[ids[i]] = struct{}{}
	}
	return uc
}

// AddAccounts adds the accounts edges to Account.
func (uc *UserCreate) AddAccounts(a ...*Account) *UserCreate {
	ids := make([]xid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uc.AddAccountIDs(ids...)
}

// AddContactIDs adds the contacts edge to Contact by ids.
func (uc *UserCreate) AddContactIDs(ids ...xid.ID) *UserCreate {
	if uc.contacts == nil {
		uc.contacts = make(map[xid.ID]struct{})
	}
	for i := range ids {
		uc.contacts[ids[i]] = struct{}{}
	}
	return uc
}

// AddContacts adds the contacts edges to Contact.
func (uc *UserCreate) AddContacts(c ...*Contact) *UserCreate {
	ids := make([]xid.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddContactIDs(ids...)
}

// AddSurveyIDs adds the surveys edge to Survey by ids.
func (uc *UserCreate) AddSurveyIDs(ids ...xid.ID) *UserCreate {
	if uc.surveys == nil {
		uc.surveys = make(map[xid.ID]struct{})
	}
	for i := range ids {
		uc.surveys[ids[i]] = struct{}{}
	}
	return uc
}

// AddSurveys adds the surveys edges to Survey.
func (uc *UserCreate) AddSurveys(s ...*Survey) *UserCreate {
	ids := make([]xid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uc.AddSurveyIDs(ids...)
}

// SetDomainID sets the domain edge to Domain by id.
func (uc *UserCreate) SetDomainID(id xid.ID) *UserCreate {
	if uc.domain == nil {
		uc.domain = make(map[xid.ID]struct{})
	}
	uc.domain[id] = struct{}{}
	return uc
}

// SetNillableDomainID sets the domain edge to Domain by id if the given value is not nil.
func (uc *UserCreate) SetNillableDomainID(id *xid.ID) *UserCreate {
	if id != nil {
		uc = uc.SetDomainID(*id)
	}
	return uc
}

// SetDomain sets the domain edge to Domain.
func (uc *UserCreate) SetDomain(d *Domain) *UserCreate {
	return uc.SetDomainID(d.ID)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if err := user.NameValidator(*uc.name); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
	}
	if uc.lastActivity == nil {
		v := user.DefaultLastActivity()
		uc.lastActivity = &v
	}
	if len(uc.domain) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"domain\"")
	}
	return uc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		}
	)
	if value := uc.id; value != nil {
		u.ID = *value
		_spec.ID.Value = *value
	}
	if value := uc.name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
		u.Name = *value
	}
	if value := uc.username; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldUsername,
		})
		u.Username = *value
	}
	if value := uc.lastActivity; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldLastActivity,
		})
		u.LastActivity = *value
	}
	if nodes := uc.accounts; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountsTable,
			Columns: []string{user.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.contacts; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ContactsTable,
			Columns: []string{user.ContactsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: contact.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.surveys; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SurveysTable,
			Columns: []string{user.SurveysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: survey.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.domain; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.DomainTable,
			Columns: []string{user.DomainColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: domain.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
