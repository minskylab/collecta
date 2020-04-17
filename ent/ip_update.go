// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/ent/ip"
	"github.com/minskylab/collecta/ent/predicate"
)

// IPUpdate is the builder for updating IP entities.
type IPUpdate struct {
	config
	predicates []predicate.IP
}

// Where adds a new predicate for the builder.
func (iu *IPUpdate) Where(ps ...predicate.IP) *IPUpdate {
	iu.predicates = append(iu.predicates, ps...)
	return iu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (iu *IPUpdate) Save(ctx context.Context) (int, error) {
	return iu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IPUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IPUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IPUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *IPUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ip.Table,
			Columns: ip.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ip.FieldID,
			},
		},
	}
	if ps := iu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ip.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// IPUpdateOne is the builder for updating a single IP entity.
type IPUpdateOne struct {
	config
	id int
}

// Save executes the query and returns the updated entity.
func (iuo *IPUpdateOne) Save(ctx context.Context) (*IP, error) {
	return iuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IPUpdateOne) SaveX(ctx context.Context) *IP {
	i, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return i
}

// Exec executes the query on the entity.
func (iuo *IPUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IPUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *IPUpdateOne) sqlSave(ctx context.Context) (i *IP, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ip.Table,
			Columns: ip.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  iuo.id,
				Type:   field.TypeInt,
				Column: ip.FieldID,
			},
		},
	}
	i = &IP{config: iuo.config}
	_spec.Assign = i.assignValues
	_spec.ScanValues = i.scanValues()
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ip.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return i, nil
}