// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/user"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	id       *uuid.UUID
	_type    *account.Type
	sub      *string
	remoteID *string
	secret   *string
	owner    map[uuid.UUID]struct{}
}

// SetType sets the type field.
func (ac *AccountCreate) SetType(a account.Type) *AccountCreate {
	ac._type = &a
	return ac
}

// SetSub sets the sub field.
func (ac *AccountCreate) SetSub(s string) *AccountCreate {
	ac.sub = &s
	return ac
}

// SetRemoteID sets the remoteID field.
func (ac *AccountCreate) SetRemoteID(s string) *AccountCreate {
	ac.remoteID = &s
	return ac
}

// SetSecret sets the secret field.
func (ac *AccountCreate) SetSecret(s string) *AccountCreate {
	ac.secret = &s
	return ac
}

// SetNillableSecret sets the secret field if the given value is not nil.
func (ac *AccountCreate) SetNillableSecret(s *string) *AccountCreate {
	if s != nil {
		ac.SetSecret(*s)
	}
	return ac
}

// SetID sets the id field.
func (ac *AccountCreate) SetID(u uuid.UUID) *AccountCreate {
	ac.id = &u
	return ac
}

// SetOwnerID sets the owner edge to User by id.
func (ac *AccountCreate) SetOwnerID(id uuid.UUID) *AccountCreate {
	if ac.owner == nil {
		ac.owner = make(map[uuid.UUID]struct{})
	}
	ac.owner[id] = struct{}{}
	return ac
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (ac *AccountCreate) SetNillableOwnerID(id *uuid.UUID) *AccountCreate {
	if id != nil {
		ac = ac.SetOwnerID(*id)
	}
	return ac
}

// SetOwner sets the owner edge to User.
func (ac *AccountCreate) SetOwner(u *User) *AccountCreate {
	return ac.SetOwnerID(u.ID)
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	if ac._type == nil {
		return nil, errors.New("ent: missing required field \"type\"")
	}
	if err := account.TypeValidator(*ac._type); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"type\": %v", err)
	}
	if ac.sub == nil {
		return nil, errors.New("ent: missing required field \"sub\"")
	}
	if err := account.SubValidator(*ac.sub); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"sub\": %v", err)
	}
	if ac.remoteID == nil {
		return nil, errors.New("ent: missing required field \"remoteID\"")
	}
	if len(ac.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return ac.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	var (
		a     = &Account{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: account.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		}
	)
	if value := ac.id; value != nil {
		a.ID = *value
		_spec.ID.Value = *value
	}
	if value := ac._type; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: account.FieldType,
		})
		a.Type = *value
	}
	if value := ac.sub; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: account.FieldSub,
		})
		a.Sub = *value
	}
	if value := ac.remoteID; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: account.FieldRemoteID,
		})
		a.RemoteID = *value
	}
	if value := ac.secret; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: account.FieldSecret,
		})
		a.Secret = *value
	}
	if nodes := ac.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.OwnerTable,
			Columns: []string{account.OwnerColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}
