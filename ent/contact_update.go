// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/ent/contact"
	"github.com/minskylab/collecta/ent/predicate"
	"github.com/minskylab/collecta/ent/user"
	"github.com/rs/xid"
)

// ContactUpdate is the builder for updating Contact entities.
type ContactUpdate struct {
	config
	name         *string
	value        *string
	kind         *contact.Kind
	principal    *bool
	validated    *bool
	fromAccount  *bool
	owner        map[xid.ID]struct{}
	clearedOwner bool
	predicates   []predicate.Contact
}

// Where adds a new predicate for the builder.
func (cu *ContactUpdate) Where(ps ...predicate.Contact) *ContactUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetName sets the name field.
func (cu *ContactUpdate) SetName(s string) *ContactUpdate {
	cu.name = &s
	return cu
}

// SetValue sets the value field.
func (cu *ContactUpdate) SetValue(s string) *ContactUpdate {
	cu.value = &s
	return cu
}

// SetKind sets the kind field.
func (cu *ContactUpdate) SetKind(c contact.Kind) *ContactUpdate {
	cu.kind = &c
	return cu
}

// SetNillableKind sets the kind field if the given value is not nil.
func (cu *ContactUpdate) SetNillableKind(c *contact.Kind) *ContactUpdate {
	if c != nil {
		cu.SetKind(*c)
	}
	return cu
}

// SetPrincipal sets the principal field.
func (cu *ContactUpdate) SetPrincipal(b bool) *ContactUpdate {
	cu.principal = &b
	return cu
}

// SetValidated sets the validated field.
func (cu *ContactUpdate) SetValidated(b bool) *ContactUpdate {
	cu.validated = &b
	return cu
}

// SetFromAccount sets the fromAccount field.
func (cu *ContactUpdate) SetFromAccount(b bool) *ContactUpdate {
	cu.fromAccount = &b
	return cu
}

// SetNillableFromAccount sets the fromAccount field if the given value is not nil.
func (cu *ContactUpdate) SetNillableFromAccount(b *bool) *ContactUpdate {
	if b != nil {
		cu.SetFromAccount(*b)
	}
	return cu
}

// SetOwnerID sets the owner edge to User by id.
func (cu *ContactUpdate) SetOwnerID(id xid.ID) *ContactUpdate {
	if cu.owner == nil {
		cu.owner = make(map[xid.ID]struct{})
	}
	cu.owner[id] = struct{}{}
	return cu
}

// SetOwner sets the owner edge to User.
func (cu *ContactUpdate) SetOwner(u *User) *ContactUpdate {
	return cu.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (cu *ContactUpdate) ClearOwner() *ContactUpdate {
	cu.clearedOwner = true
	return cu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *ContactUpdate) Save(ctx context.Context) (int, error) {
	if cu.value != nil {
		if err := contact.ValueValidator(*cu.value); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"value\": %v", err)
		}
	}
	if cu.kind != nil {
		if err := contact.KindValidator(*cu.kind); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"kind\": %v", err)
		}
	}
	if len(cu.owner) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	if cu.clearedOwner && cu.owner == nil {
		return 0, errors.New("ent: clearing a unique edge \"owner\"")
	}
	return cu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ContactUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ContactUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ContactUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ContactUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contact.Table,
			Columns: contact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: contact.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := cu.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: contact.FieldName,
		})
	}
	if value := cu.value; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: contact.FieldValue,
		})
	}
	if value := cu.kind; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: contact.FieldKind,
		})
	}
	if value := cu.principal; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: contact.FieldPrincipal,
		})
	}
	if value := cu.validated; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: contact.FieldValidated,
		})
	}
	if value := cu.fromAccount; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: contact.FieldFromAccount,
		})
	}
	if cu.clearedOwner {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.OwnerTable,
			Columns: []string{contact.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.OwnerTable,
			Columns: []string{contact.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contact.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ContactUpdateOne is the builder for updating a single Contact entity.
type ContactUpdateOne struct {
	config
	id           xid.ID
	name         *string
	value        *string
	kind         *contact.Kind
	principal    *bool
	validated    *bool
	fromAccount  *bool
	owner        map[xid.ID]struct{}
	clearedOwner bool
}

// SetName sets the name field.
func (cuo *ContactUpdateOne) SetName(s string) *ContactUpdateOne {
	cuo.name = &s
	return cuo
}

// SetValue sets the value field.
func (cuo *ContactUpdateOne) SetValue(s string) *ContactUpdateOne {
	cuo.value = &s
	return cuo
}

// SetKind sets the kind field.
func (cuo *ContactUpdateOne) SetKind(c contact.Kind) *ContactUpdateOne {
	cuo.kind = &c
	return cuo
}

// SetNillableKind sets the kind field if the given value is not nil.
func (cuo *ContactUpdateOne) SetNillableKind(c *contact.Kind) *ContactUpdateOne {
	if c != nil {
		cuo.SetKind(*c)
	}
	return cuo
}

// SetPrincipal sets the principal field.
func (cuo *ContactUpdateOne) SetPrincipal(b bool) *ContactUpdateOne {
	cuo.principal = &b
	return cuo
}

// SetValidated sets the validated field.
func (cuo *ContactUpdateOne) SetValidated(b bool) *ContactUpdateOne {
	cuo.validated = &b
	return cuo
}

// SetFromAccount sets the fromAccount field.
func (cuo *ContactUpdateOne) SetFromAccount(b bool) *ContactUpdateOne {
	cuo.fromAccount = &b
	return cuo
}

// SetNillableFromAccount sets the fromAccount field if the given value is not nil.
func (cuo *ContactUpdateOne) SetNillableFromAccount(b *bool) *ContactUpdateOne {
	if b != nil {
		cuo.SetFromAccount(*b)
	}
	return cuo
}

// SetOwnerID sets the owner edge to User by id.
func (cuo *ContactUpdateOne) SetOwnerID(id xid.ID) *ContactUpdateOne {
	if cuo.owner == nil {
		cuo.owner = make(map[xid.ID]struct{})
	}
	cuo.owner[id] = struct{}{}
	return cuo
}

// SetOwner sets the owner edge to User.
func (cuo *ContactUpdateOne) SetOwner(u *User) *ContactUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (cuo *ContactUpdateOne) ClearOwner() *ContactUpdateOne {
	cuo.clearedOwner = true
	return cuo
}

// Save executes the query and returns the updated entity.
func (cuo *ContactUpdateOne) Save(ctx context.Context) (*Contact, error) {
	if cuo.value != nil {
		if err := contact.ValueValidator(*cuo.value); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"value\": %v", err)
		}
	}
	if cuo.kind != nil {
		if err := contact.KindValidator(*cuo.kind); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"kind\": %v", err)
		}
	}
	if len(cuo.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	if cuo.clearedOwner && cuo.owner == nil {
		return nil, errors.New("ent: clearing a unique edge \"owner\"")
	}
	return cuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ContactUpdateOne) SaveX(ctx context.Context) *Contact {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *ContactUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ContactUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ContactUpdateOne) sqlSave(ctx context.Context) (c *Contact, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contact.Table,
			Columns: contact.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  cuo.id,
				Type:   field.TypeUUID,
				Column: contact.FieldID,
			},
		},
	}
	if value := cuo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: contact.FieldName,
		})
	}
	if value := cuo.value; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: contact.FieldValue,
		})
	}
	if value := cuo.kind; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: contact.FieldKind,
		})
	}
	if value := cuo.principal; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: contact.FieldPrincipal,
		})
	}
	if value := cuo.validated; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: contact.FieldValidated,
		})
	}
	if value := cuo.fromAccount; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: contact.FieldFromAccount,
		})
	}
	if cuo.clearedOwner {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.OwnerTable,
			Columns: []string{contact.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contact.OwnerTable,
			Columns: []string{contact.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	c = &Contact{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contact.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
