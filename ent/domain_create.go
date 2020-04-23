// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/ent/user"
)

// DomainCreate is the builder for creating a Domain entity.
type DomainCreate struct {
	config
	id             *uuid.UUID
	tags           *[]string
	name           *string
	email          *string
	domain         *string
	collectaDomain *string
	surveys        map[uuid.UUID]struct{}
	users          map[uuid.UUID]struct{}
	admins         map[uuid.UUID]struct{}
}

// SetTags sets the tags field.
func (dc *DomainCreate) SetTags(s []string) *DomainCreate {
	dc.tags = &s
	return dc
}

// SetName sets the name field.
func (dc *DomainCreate) SetName(s string) *DomainCreate {
	dc.name = &s
	return dc
}

// SetEmail sets the email field.
func (dc *DomainCreate) SetEmail(s string) *DomainCreate {
	dc.email = &s
	return dc
}

// SetDomain sets the domain field.
func (dc *DomainCreate) SetDomain(s string) *DomainCreate {
	dc.domain = &s
	return dc
}

// SetCollectaDomain sets the collectaDomain field.
func (dc *DomainCreate) SetCollectaDomain(s string) *DomainCreate {
	dc.collectaDomain = &s
	return dc
}

// SetID sets the id field.
func (dc *DomainCreate) SetID(u uuid.UUID) *DomainCreate {
	dc.id = &u
	return dc
}

// AddSurveyIDs adds the surveys edge to Survey by ids.
func (dc *DomainCreate) AddSurveyIDs(ids ...uuid.UUID) *DomainCreate {
	if dc.surveys == nil {
		dc.surveys = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		dc.surveys[ids[i]] = struct{}{}
	}
	return dc
}

// AddSurveys adds the surveys edges to Survey.
func (dc *DomainCreate) AddSurveys(s ...*Survey) *DomainCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return dc.AddSurveyIDs(ids...)
}

// AddUserIDs adds the users edge to User by ids.
func (dc *DomainCreate) AddUserIDs(ids ...uuid.UUID) *DomainCreate {
	if dc.users == nil {
		dc.users = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		dc.users[ids[i]] = struct{}{}
	}
	return dc
}

// AddUsers adds the users edges to User.
func (dc *DomainCreate) AddUsers(u ...*User) *DomainCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddUserIDs(ids...)
}

// AddAdminIDs adds the admins edge to User by ids.
func (dc *DomainCreate) AddAdminIDs(ids ...uuid.UUID) *DomainCreate {
	if dc.admins == nil {
		dc.admins = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		dc.admins[ids[i]] = struct{}{}
	}
	return dc
}

// AddAdmins adds the admins edges to User.
func (dc *DomainCreate) AddAdmins(u ...*User) *DomainCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddAdminIDs(ids...)
}

// Save creates the Domain in the database.
func (dc *DomainCreate) Save(ctx context.Context) (*Domain, error) {
	if dc.tags == nil {
		return nil, errors.New("ent: missing required field \"tags\"")
	}
	if dc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if err := domain.NameValidator(*dc.name); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
	}
	if dc.email == nil {
		return nil, errors.New("ent: missing required field \"email\"")
	}
	if err := domain.EmailValidator(*dc.email); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"email\": %v", err)
	}
	if dc.domain == nil {
		return nil, errors.New("ent: missing required field \"domain\"")
	}
	if dc.collectaDomain == nil {
		return nil, errors.New("ent: missing required field \"collectaDomain\"")
	}
	return dc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DomainCreate) SaveX(ctx context.Context) *Domain {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DomainCreate) sqlSave(ctx context.Context) (*Domain, error) {
	var (
		d     = &Domain{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: domain.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: domain.FieldID,
			},
		}
	)
	if value := dc.id; value != nil {
		d.ID = *value
		_spec.ID.Value = *value
	}
	if value := dc.tags; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: domain.FieldTags,
		})
		d.Tags = *value
	}
	if value := dc.name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: domain.FieldName,
		})
		d.Name = *value
	}
	if value := dc.email; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: domain.FieldEmail,
		})
		d.Email = *value
	}
	if value := dc.domain; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: domain.FieldDomain,
		})
		d.Domain = *value
	}
	if value := dc.collectaDomain; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: domain.FieldCollectaDomain,
		})
		d.CollectaDomain = *value
	}
	if nodes := dc.surveys; len(nodes) > 0 {
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
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.users; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.UsersTable,
			Columns: domain.UsersPrimaryKey,
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
	if nodes := dc.admins; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   domain.AdminsTable,
			Columns: domain.AdminsPrimaryKey,
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
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}
