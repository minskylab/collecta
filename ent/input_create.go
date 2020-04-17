// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/ent/question"
)

// InputCreate is the builder for creating a Input entity.
type InputCreate struct {
	config
	id       *uuid.UUID
	kind     *input.Kind
	multiple *bool
	defaults *[]string
	options  *map[string]string
	question map[uuid.UUID]struct{}
}

// SetKind sets the kind field.
func (ic *InputCreate) SetKind(i input.Kind) *InputCreate {
	ic.kind = &i
	return ic
}

// SetMultiple sets the multiple field.
func (ic *InputCreate) SetMultiple(b bool) *InputCreate {
	ic.multiple = &b
	return ic
}

// SetNillableMultiple sets the multiple field if the given value is not nil.
func (ic *InputCreate) SetNillableMultiple(b *bool) *InputCreate {
	if b != nil {
		ic.SetMultiple(*b)
	}
	return ic
}

// SetDefaults sets the defaults field.
func (ic *InputCreate) SetDefaults(s []string) *InputCreate {
	ic.defaults = &s
	return ic
}

// SetOptions sets the options field.
func (ic *InputCreate) SetOptions(m map[string]string) *InputCreate {
	ic.options = &m
	return ic
}

// SetID sets the id field.
func (ic *InputCreate) SetID(u uuid.UUID) *InputCreate {
	ic.id = &u
	return ic
}

// SetQuestionID sets the question edge to Question by id.
func (ic *InputCreate) SetQuestionID(id uuid.UUID) *InputCreate {
	if ic.question == nil {
		ic.question = make(map[uuid.UUID]struct{})
	}
	ic.question[id] = struct{}{}
	return ic
}

// SetQuestion sets the question edge to Question.
func (ic *InputCreate) SetQuestion(q *Question) *InputCreate {
	return ic.SetQuestionID(q.ID)
}

// Save creates the Input in the database.
func (ic *InputCreate) Save(ctx context.Context) (*Input, error) {
	if ic.kind == nil {
		return nil, errors.New("ent: missing required field \"kind\"")
	}
	if err := input.KindValidator(*ic.kind); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"kind\": %v", err)
	}
	if ic.multiple == nil {
		v := input.DefaultMultiple
		ic.multiple = &v
	}
	if len(ic.question) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"question\"")
	}
	if ic.question == nil {
		return nil, errors.New("ent: missing required edge \"question\"")
	}
	return ic.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InputCreate) SaveX(ctx context.Context) *Input {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ic *InputCreate) sqlSave(ctx context.Context) (*Input, error) {
	var (
		i     = &Input{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: input.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: input.FieldID,
			},
		}
	)
	if value := ic.id; value != nil {
		i.ID = *value
		_spec.ID.Value = *value
	}
	if value := ic.kind; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: input.FieldKind,
		})
		i.Kind = *value
	}
	if value := ic.multiple; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: input.FieldMultiple,
		})
		i.Multiple = *value
	}
	if value := ic.defaults; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: input.FieldDefaults,
		})
		i.Defaults = *value
	}
	if value := ic.options; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: input.FieldOptions,
		})
		i.Options = *value
	}
	if nodes := ic.question; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   input.QuestionTable,
			Columns: []string{input.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return i, nil
}
