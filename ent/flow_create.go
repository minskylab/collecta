// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/question"
)

// FlowCreate is the builder for creating a Flow entity.
type FlowCreate struct {
	config
	id         *uuid.UUID
	state      *uuid.UUID
	stateTable *string
	inputs     *[]string
	questions  map[uuid.UUID]struct{}
}

// SetState sets the state field.
func (fc *FlowCreate) SetState(u uuid.UUID) *FlowCreate {
	fc.state = &u
	return fc
}

// SetStateTable sets the stateTable field.
func (fc *FlowCreate) SetStateTable(s string) *FlowCreate {
	fc.stateTable = &s
	return fc
}

// SetInputs sets the inputs field.
func (fc *FlowCreate) SetInputs(s []string) *FlowCreate {
	fc.inputs = &s
	return fc
}

// SetID sets the id field.
func (fc *FlowCreate) SetID(u uuid.UUID) *FlowCreate {
	fc.id = &u
	return fc
}

// AddQuestionIDs adds the questions edge to Question by ids.
func (fc *FlowCreate) AddQuestionIDs(ids ...uuid.UUID) *FlowCreate {
	if fc.questions == nil {
		fc.questions = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		fc.questions[ids[i]] = struct{}{}
	}
	return fc
}

// AddQuestions adds the questions edges to Question.
func (fc *FlowCreate) AddQuestions(q ...*Question) *FlowCreate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return fc.AddQuestionIDs(ids...)
}

// Save creates the Flow in the database.
func (fc *FlowCreate) Save(ctx context.Context) (*Flow, error) {
	if fc.state == nil {
		return nil, errors.New("ent: missing required field \"state\"")
	}
	if fc.stateTable == nil {
		return nil, errors.New("ent: missing required field \"stateTable\"")
	}
	if err := flow.StateTableValidator(*fc.stateTable); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"stateTable\": %v", err)
	}
	if fc.questions == nil {
		return nil, errors.New("ent: missing required edge \"questions\"")
	}
	return fc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FlowCreate) SaveX(ctx context.Context) *Flow {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FlowCreate) sqlSave(ctx context.Context) (*Flow, error) {
	var (
		f     = &Flow{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: flow.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: flow.FieldID,
			},
		}
	)
	if value := fc.id; value != nil {
		f.ID = *value
		_spec.ID.Value = *value
	}
	if value := fc.state; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  *value,
			Column: flow.FieldState,
		})
		f.State = *value
	}
	if value := fc.stateTable; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: flow.FieldStateTable,
		})
		f.StateTable = *value
	}
	if value := fc.inputs; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: flow.FieldInputs,
		})
		f.Inputs = *value
	}
	if nodes := fc.questions; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   flow.QuestionsTable,
			Columns: []string{flow.QuestionsColumn},
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
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return f, nil
}
