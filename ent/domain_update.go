// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/person"
	"github.com/minskylab/collecta/ent/predicate"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/uuid"
)

// DomainUpdate is the builder for updating Domain entities.
type DomainUpdate struct {
	config
	hooks      []Hook
	mutation   *DomainMutation
	predicates []predicate.Domain
}

// Where adds a new predicate for the builder.
func (du *DomainUpdate) Where(ps ...predicate.Domain) *DomainUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetName sets the name field.
func (du *DomainUpdate) SetName(s string) *DomainUpdate {
	du.mutation.SetName(s)
	return du
}

// SetEmail sets the email field.
func (du *DomainUpdate) SetEmail(s string) *DomainUpdate {
	du.mutation.SetEmail(s)
	return du
}

// SetDomain sets the domain field.
func (du *DomainUpdate) SetDomain(s string) *DomainUpdate {
	du.mutation.SetDomain(s)
	return du
}

// SetCallback sets the callback field.
func (du *DomainUpdate) SetCallback(s string) *DomainUpdate {
	du.mutation.SetCallback(s)
	return du
}

// SetTags sets the tags field.
func (du *DomainUpdate) SetTags(s []string) *DomainUpdate {
	du.mutation.SetTags(s)
	return du
}

// AddSurveyIDs adds the surveys edge to Survey by ids.
func (du *DomainUpdate) AddSurveyIDs(ids ...uuid.UUID) *DomainUpdate {
	du.mutation.AddSurveyIDs(ids...)
	return du
}

// AddSurveys adds the surveys edges to Survey.
func (du *DomainUpdate) AddSurveys(s ...*Survey) *DomainUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.AddSurveyIDs(ids...)
}

// AddUserIDs adds the users edge to Person by ids.
func (du *DomainUpdate) AddUserIDs(ids ...uuid.UUID) *DomainUpdate {
	du.mutation.AddUserIDs(ids...)
	return du
}

// AddUsers adds the users edges to Person.
func (du *DomainUpdate) AddUsers(p ...*Person) *DomainUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.AddUserIDs(ids...)
}

// AddAdminIDs adds the admins edge to Person by ids.
func (du *DomainUpdate) AddAdminIDs(ids ...uuid.UUID) *DomainUpdate {
	du.mutation.AddAdminIDs(ids...)
	return du
}

// AddAdmins adds the admins edges to Person.
func (du *DomainUpdate) AddAdmins(p ...*Person) *DomainUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.AddAdminIDs(ids...)
}

// RemoveSurveyIDs removes the surveys edge to Survey by ids.
func (du *DomainUpdate) RemoveSurveyIDs(ids ...uuid.UUID) *DomainUpdate {
	du.mutation.RemoveSurveyIDs(ids...)
	return du
}

// RemoveSurveys removes surveys edges to Survey.
func (du *DomainUpdate) RemoveSurveys(s ...*Survey) *DomainUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.RemoveSurveyIDs(ids...)
}

// RemoveUserIDs removes the users edge to Person by ids.
func (du *DomainUpdate) RemoveUserIDs(ids ...uuid.UUID) *DomainUpdate {
	du.mutation.RemoveUserIDs(ids...)
	return du
}

// RemoveUsers removes users edges to Person.
func (du *DomainUpdate) RemoveUsers(p ...*Person) *DomainUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.RemoveUserIDs(ids...)
}

// RemoveAdminIDs removes the admins edge to Person by ids.
func (du *DomainUpdate) RemoveAdminIDs(ids ...uuid.UUID) *DomainUpdate {
	du.mutation.RemoveAdminIDs(ids...)
	return du
}

// RemoveAdmins removes admins edges to Person.
func (du *DomainUpdate) RemoveAdmins(p ...*Person) *DomainUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.RemoveAdminIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DomainUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := du.mutation.Name(); ok {
		if err := domain.NameValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if v, ok := du.mutation.Email(); ok {
		if err := domain.EmailValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"email\": %v", err)
		}
	}

	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DomainMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DomainUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DomainUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DomainUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DomainUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   domain.Table,
			Columns: domain.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: domain.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: domain.FieldName,
		})
	}
	if value, ok := du.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: domain.FieldEmail,
		})
	}
	if value, ok := du.mutation.Domain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: domain.FieldDomain,
		})
	}
	if value, ok := du.mutation.Callback(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: domain.FieldCallback,
		})
	}
	if value, ok := du.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: domain.FieldTags,
		})
	}
	if nodes := du.mutation.RemovedSurveysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.SurveysTable,
			Columns: []string{domain.SurveysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: survey.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.SurveysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.SurveysTable,
			Columns: []string{domain.SurveysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: survey.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := du.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.UsersTable,
			Columns: domain.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.UsersTable,
			Columns: domain.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := du.mutation.RemovedAdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.AdminsTable,
			Columns: domain.AdminsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.AdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.AdminsTable,
			Columns: domain.AdminsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{domain.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DomainUpdateOne is the builder for updating a single Domain entity.
type DomainUpdateOne struct {
	config
	hooks    []Hook
	mutation *DomainMutation
}

// SetName sets the name field.
func (duo *DomainUpdateOne) SetName(s string) *DomainUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetEmail sets the email field.
func (duo *DomainUpdateOne) SetEmail(s string) *DomainUpdateOne {
	duo.mutation.SetEmail(s)
	return duo
}

// SetDomain sets the domain field.
func (duo *DomainUpdateOne) SetDomain(s string) *DomainUpdateOne {
	duo.mutation.SetDomain(s)
	return duo
}

// SetCallback sets the callback field.
func (duo *DomainUpdateOne) SetCallback(s string) *DomainUpdateOne {
	duo.mutation.SetCallback(s)
	return duo
}

// SetTags sets the tags field.
func (duo *DomainUpdateOne) SetTags(s []string) *DomainUpdateOne {
	duo.mutation.SetTags(s)
	return duo
}

// AddSurveyIDs adds the surveys edge to Survey by ids.
func (duo *DomainUpdateOne) AddSurveyIDs(ids ...uuid.UUID) *DomainUpdateOne {
	duo.mutation.AddSurveyIDs(ids...)
	return duo
}

// AddSurveys adds the surveys edges to Survey.
func (duo *DomainUpdateOne) AddSurveys(s ...*Survey) *DomainUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.AddSurveyIDs(ids...)
}

// AddUserIDs adds the users edge to Person by ids.
func (duo *DomainUpdateOne) AddUserIDs(ids ...uuid.UUID) *DomainUpdateOne {
	duo.mutation.AddUserIDs(ids...)
	return duo
}

// AddUsers adds the users edges to Person.
func (duo *DomainUpdateOne) AddUsers(p ...*Person) *DomainUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.AddUserIDs(ids...)
}

// AddAdminIDs adds the admins edge to Person by ids.
func (duo *DomainUpdateOne) AddAdminIDs(ids ...uuid.UUID) *DomainUpdateOne {
	duo.mutation.AddAdminIDs(ids...)
	return duo
}

// AddAdmins adds the admins edges to Person.
func (duo *DomainUpdateOne) AddAdmins(p ...*Person) *DomainUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.AddAdminIDs(ids...)
}

// RemoveSurveyIDs removes the surveys edge to Survey by ids.
func (duo *DomainUpdateOne) RemoveSurveyIDs(ids ...uuid.UUID) *DomainUpdateOne {
	duo.mutation.RemoveSurveyIDs(ids...)
	return duo
}

// RemoveSurveys removes surveys edges to Survey.
func (duo *DomainUpdateOne) RemoveSurveys(s ...*Survey) *DomainUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.RemoveSurveyIDs(ids...)
}

// RemoveUserIDs removes the users edge to Person by ids.
func (duo *DomainUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *DomainUpdateOne {
	duo.mutation.RemoveUserIDs(ids...)
	return duo
}

// RemoveUsers removes users edges to Person.
func (duo *DomainUpdateOne) RemoveUsers(p ...*Person) *DomainUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.RemoveUserIDs(ids...)
}

// RemoveAdminIDs removes the admins edge to Person by ids.
func (duo *DomainUpdateOne) RemoveAdminIDs(ids ...uuid.UUID) *DomainUpdateOne {
	duo.mutation.RemoveAdminIDs(ids...)
	return duo
}

// RemoveAdmins removes admins edges to Person.
func (duo *DomainUpdateOne) RemoveAdmins(p ...*Person) *DomainUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.RemoveAdminIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DomainUpdateOne) Save(ctx context.Context) (*Domain, error) {
	if v, ok := duo.mutation.Name(); ok {
		if err := domain.NameValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if v, ok := duo.mutation.Email(); ok {
		if err := domain.EmailValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"email\": %v", err)
		}
	}

	var (
		err  error
		node *Domain
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DomainMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DomainUpdateOne) SaveX(ctx context.Context) *Domain {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DomainUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DomainUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DomainUpdateOne) sqlSave(ctx context.Context) (d *Domain, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   domain.Table,
			Columns: domain.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: domain.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Domain.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: domain.FieldName,
		})
	}
	if value, ok := duo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: domain.FieldEmail,
		})
	}
	if value, ok := duo.mutation.Domain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: domain.FieldDomain,
		})
	}
	if value, ok := duo.mutation.Callback(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: domain.FieldCallback,
		})
	}
	if value, ok := duo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: domain.FieldTags,
		})
	}
	if nodes := duo.mutation.RemovedSurveysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.SurveysTable,
			Columns: []string{domain.SurveysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: survey.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.SurveysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   domain.SurveysTable,
			Columns: []string{domain.SurveysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: survey.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := duo.mutation.RemovedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.UsersTable,
			Columns: domain.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.UsersTable,
			Columns: domain.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := duo.mutation.RemovedAdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.AdminsTable,
			Columns: domain.AdminsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.AdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.AdminsTable,
			Columns: domain.AdminsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	d = &Domain{config: duo.config}
	_spec.Assign = d.assignValues
	_spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{domain.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
