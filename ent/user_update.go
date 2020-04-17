// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/contact"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/predicate"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	name            *string
	username        *string
	clearusername   bool
	lastActivity    *time.Time
	picture         *string
	clearpicture    bool
	accounts        map[uuid.UUID]struct{}
	contacts        map[uuid.UUID]struct{}
	surveys         map[uuid.UUID]struct{}
	domain          map[uuid.UUID]struct{}
	removedAccounts map[uuid.UUID]struct{}
	removedContacts map[uuid.UUID]struct{}
	removedSurveys  map[uuid.UUID]struct{}
	clearedDomain   bool
	predicates      []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.name = &s
	return uu
}

// SetUsername sets the username field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.username = &s
	return uu
}

// SetNillableUsername sets the username field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// ClearUsername clears the value of username.
func (uu *UserUpdate) ClearUsername() *UserUpdate {
	uu.username = nil
	uu.clearusername = true
	return uu
}

// SetLastActivity sets the lastActivity field.
func (uu *UserUpdate) SetLastActivity(t time.Time) *UserUpdate {
	uu.lastActivity = &t
	return uu
}

// SetNillableLastActivity sets the lastActivity field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastActivity(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetLastActivity(*t)
	}
	return uu
}

// SetPicture sets the picture field.
func (uu *UserUpdate) SetPicture(s string) *UserUpdate {
	uu.picture = &s
	return uu
}

// SetNillablePicture sets the picture field if the given value is not nil.
func (uu *UserUpdate) SetNillablePicture(s *string) *UserUpdate {
	if s != nil {
		uu.SetPicture(*s)
	}
	return uu
}

// ClearPicture clears the value of picture.
func (uu *UserUpdate) ClearPicture() *UserUpdate {
	uu.picture = nil
	uu.clearpicture = true
	return uu
}

// AddAccountIDs adds the accounts edge to Account by ids.
func (uu *UserUpdate) AddAccountIDs(ids ...uuid.UUID) *UserUpdate {
	if uu.accounts == nil {
		uu.accounts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		uu.accounts[ids[i]] = struct{}{}
	}
	return uu
}

// AddAccounts adds the accounts edges to Account.
func (uu *UserUpdate) AddAccounts(a ...*Account) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddAccountIDs(ids...)
}

// AddContactIDs adds the contacts edge to Contact by ids.
func (uu *UserUpdate) AddContactIDs(ids ...uuid.UUID) *UserUpdate {
	if uu.contacts == nil {
		uu.contacts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		uu.contacts[ids[i]] = struct{}{}
	}
	return uu
}

// AddContacts adds the contacts edges to Contact.
func (uu *UserUpdate) AddContacts(c ...*Contact) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddContactIDs(ids...)
}

// AddSurveyIDs adds the surveys edge to Survey by ids.
func (uu *UserUpdate) AddSurveyIDs(ids ...uuid.UUID) *UserUpdate {
	if uu.surveys == nil {
		uu.surveys = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		uu.surveys[ids[i]] = struct{}{}
	}
	return uu
}

// AddSurveys adds the surveys edges to Survey.
func (uu *UserUpdate) AddSurveys(s ...*Survey) *UserUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddSurveyIDs(ids...)
}

// SetDomainID sets the domain edge to Domain by id.
func (uu *UserUpdate) SetDomainID(id uuid.UUID) *UserUpdate {
	if uu.domain == nil {
		uu.domain = make(map[uuid.UUID]struct{})
	}
	uu.domain[id] = struct{}{}
	return uu
}

// SetNillableDomainID sets the domain edge to Domain by id if the given value is not nil.
func (uu *UserUpdate) SetNillableDomainID(id *uuid.UUID) *UserUpdate {
	if id != nil {
		uu = uu.SetDomainID(*id)
	}
	return uu
}

// SetDomain sets the domain edge to Domain.
func (uu *UserUpdate) SetDomain(d *Domain) *UserUpdate {
	return uu.SetDomainID(d.ID)
}

// RemoveAccountIDs removes the accounts edge to Account by ids.
func (uu *UserUpdate) RemoveAccountIDs(ids ...uuid.UUID) *UserUpdate {
	if uu.removedAccounts == nil {
		uu.removedAccounts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		uu.removedAccounts[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveAccounts removes accounts edges to Account.
func (uu *UserUpdate) RemoveAccounts(a ...*Account) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveAccountIDs(ids...)
}

// RemoveContactIDs removes the contacts edge to Contact by ids.
func (uu *UserUpdate) RemoveContactIDs(ids ...uuid.UUID) *UserUpdate {
	if uu.removedContacts == nil {
		uu.removedContacts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		uu.removedContacts[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveContacts removes contacts edges to Contact.
func (uu *UserUpdate) RemoveContacts(c ...*Contact) *UserUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveContactIDs(ids...)
}

// RemoveSurveyIDs removes the surveys edge to Survey by ids.
func (uu *UserUpdate) RemoveSurveyIDs(ids ...uuid.UUID) *UserUpdate {
	if uu.removedSurveys == nil {
		uu.removedSurveys = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		uu.removedSurveys[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveSurveys removes surveys edges to Survey.
func (uu *UserUpdate) RemoveSurveys(s ...*Survey) *UserUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveSurveyIDs(ids...)
}

// ClearDomain clears the domain edge to Domain.
func (uu *UserUpdate) ClearDomain() *UserUpdate {
	uu.clearedDomain = true
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if uu.name != nil {
		if err := user.NameValidator(*uu.name); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(uu.domain) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"domain\"")
	}
	return uu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := uu.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
	}
	if value := uu.username; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldUsername,
		})
	}
	if uu.clearusername {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldUsername,
		})
	}
	if value := uu.lastActivity; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldLastActivity,
		})
	}
	if value := uu.picture; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldPicture,
		})
	}
	if uu.clearpicture {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPicture,
		})
	}
	if nodes := uu.removedAccounts; len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.accounts; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uu.removedContacts; len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.contacts; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uu.removedSurveys; len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.surveys; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.clearedDomain {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.domain; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	id              uuid.UUID
	name            *string
	username        *string
	clearusername   bool
	lastActivity    *time.Time
	picture         *string
	clearpicture    bool
	accounts        map[uuid.UUID]struct{}
	contacts        map[uuid.UUID]struct{}
	surveys         map[uuid.UUID]struct{}
	domain          map[uuid.UUID]struct{}
	removedAccounts map[uuid.UUID]struct{}
	removedContacts map[uuid.UUID]struct{}
	removedSurveys  map[uuid.UUID]struct{}
	clearedDomain   bool
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// SetUsername sets the username field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.username = &s
	return uuo
}

// SetNillableUsername sets the username field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// ClearUsername clears the value of username.
func (uuo *UserUpdateOne) ClearUsername() *UserUpdateOne {
	uuo.username = nil
	uuo.clearusername = true
	return uuo
}

// SetLastActivity sets the lastActivity field.
func (uuo *UserUpdateOne) SetLastActivity(t time.Time) *UserUpdateOne {
	uuo.lastActivity = &t
	return uuo
}

// SetNillableLastActivity sets the lastActivity field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastActivity(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetLastActivity(*t)
	}
	return uuo
}

// SetPicture sets the picture field.
func (uuo *UserUpdateOne) SetPicture(s string) *UserUpdateOne {
	uuo.picture = &s
	return uuo
}

// SetNillablePicture sets the picture field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePicture(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPicture(*s)
	}
	return uuo
}

// ClearPicture clears the value of picture.
func (uuo *UserUpdateOne) ClearPicture() *UserUpdateOne {
	uuo.picture = nil
	uuo.clearpicture = true
	return uuo
}

// AddAccountIDs adds the accounts edge to Account by ids.
func (uuo *UserUpdateOne) AddAccountIDs(ids ...uuid.UUID) *UserUpdateOne {
	if uuo.accounts == nil {
		uuo.accounts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		uuo.accounts[ids[i]] = struct{}{}
	}
	return uuo
}

// AddAccounts adds the accounts edges to Account.
func (uuo *UserUpdateOne) AddAccounts(a ...*Account) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddAccountIDs(ids...)
}

// AddContactIDs adds the contacts edge to Contact by ids.
func (uuo *UserUpdateOne) AddContactIDs(ids ...uuid.UUID) *UserUpdateOne {
	if uuo.contacts == nil {
		uuo.contacts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		uuo.contacts[ids[i]] = struct{}{}
	}
	return uuo
}

// AddContacts adds the contacts edges to Contact.
func (uuo *UserUpdateOne) AddContacts(c ...*Contact) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddContactIDs(ids...)
}

// AddSurveyIDs adds the surveys edge to Survey by ids.
func (uuo *UserUpdateOne) AddSurveyIDs(ids ...uuid.UUID) *UserUpdateOne {
	if uuo.surveys == nil {
		uuo.surveys = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		uuo.surveys[ids[i]] = struct{}{}
	}
	return uuo
}

// AddSurveys adds the surveys edges to Survey.
func (uuo *UserUpdateOne) AddSurveys(s ...*Survey) *UserUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddSurveyIDs(ids...)
}

// SetDomainID sets the domain edge to Domain by id.
func (uuo *UserUpdateOne) SetDomainID(id uuid.UUID) *UserUpdateOne {
	if uuo.domain == nil {
		uuo.domain = make(map[uuid.UUID]struct{})
	}
	uuo.domain[id] = struct{}{}
	return uuo
}

// SetNillableDomainID sets the domain edge to Domain by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDomainID(id *uuid.UUID) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetDomainID(*id)
	}
	return uuo
}

// SetDomain sets the domain edge to Domain.
func (uuo *UserUpdateOne) SetDomain(d *Domain) *UserUpdateOne {
	return uuo.SetDomainID(d.ID)
}

// RemoveAccountIDs removes the accounts edge to Account by ids.
func (uuo *UserUpdateOne) RemoveAccountIDs(ids ...uuid.UUID) *UserUpdateOne {
	if uuo.removedAccounts == nil {
		uuo.removedAccounts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		uuo.removedAccounts[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveAccounts removes accounts edges to Account.
func (uuo *UserUpdateOne) RemoveAccounts(a ...*Account) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveAccountIDs(ids...)
}

// RemoveContactIDs removes the contacts edge to Contact by ids.
func (uuo *UserUpdateOne) RemoveContactIDs(ids ...uuid.UUID) *UserUpdateOne {
	if uuo.removedContacts == nil {
		uuo.removedContacts = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		uuo.removedContacts[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveContacts removes contacts edges to Contact.
func (uuo *UserUpdateOne) RemoveContacts(c ...*Contact) *UserUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveContactIDs(ids...)
}

// RemoveSurveyIDs removes the surveys edge to Survey by ids.
func (uuo *UserUpdateOne) RemoveSurveyIDs(ids ...uuid.UUID) *UserUpdateOne {
	if uuo.removedSurveys == nil {
		uuo.removedSurveys = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		uuo.removedSurveys[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveSurveys removes surveys edges to Survey.
func (uuo *UserUpdateOne) RemoveSurveys(s ...*Survey) *UserUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveSurveyIDs(ids...)
}

// ClearDomain clears the domain edge to Domain.
func (uuo *UserUpdateOne) ClearDomain() *UserUpdateOne {
	uuo.clearedDomain = true
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if uuo.name != nil {
		if err := user.NameValidator(*uuo.name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(uuo.domain) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"domain\"")
	}
	return uuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  uuo.id,
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
	if value := uuo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
	}
	if value := uuo.username; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldUsername,
		})
	}
	if uuo.clearusername {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldUsername,
		})
	}
	if value := uuo.lastActivity; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldLastActivity,
		})
	}
	if value := uuo.picture; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldPicture,
		})
	}
	if uuo.clearpicture {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPicture,
		})
	}
	if nodes := uuo.removedAccounts; len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.accounts; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uuo.removedContacts; len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.contacts; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uuo.removedSurveys; len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.surveys; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.clearedDomain {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.domain; len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}
