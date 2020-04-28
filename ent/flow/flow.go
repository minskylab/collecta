// Code generated by entc, DO NOT EDIT.

package flow

import (
	"github.com/minskylab/collecta/ent/schema"
)

const (
	// Label holds the string label denoting the flow type in the database.
	Label = "flow"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldState holds the string denoting the state vertex property in the database.
	FieldState = "state"
	// FieldStateTable holds the string denoting the statetable vertex property in the database.
	FieldStateTable = "state_table"
	// FieldInitialState holds the string denoting the initialstate vertex property in the database.
	FieldInitialState = "initial_state"
	// FieldTerminationState holds the string denoting the terminationstate vertex property in the database.
	FieldTerminationState = "termination_state"
	// FieldPastState holds the string denoting the paststate vertex property in the database.
	FieldPastState = "past_state"
	// FieldInputs holds the string denoting the inputs vertex property in the database.
	FieldInputs = "inputs"

	// Table holds the table name of the flow in the database.
	Table = "flows"
	// SurveyTable is the table the holds the survey relation/edge.
	SurveyTable = "flows"
	// SurveyInverseTable is the table name for the Survey entity.
	// It exists in this package in order to avoid circular dependency with the "survey" package.
	SurveyInverseTable = "surveys"
	// SurveyColumn is the table column denoting the survey relation/edge.
	SurveyColumn = "survey_flow"
	// QuestionsTable is the table the holds the questions relation/edge.
	QuestionsTable = "questions"
	// QuestionsInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionsInverseTable = "questions"
	// QuestionsColumn is the table column denoting the questions relation/edge.
	QuestionsColumn = "flow_questions"
)

// Columns holds all SQL columns for flow fields.
var Columns = []string{
	FieldID,
	FieldState,
	FieldStateTable,
	FieldInitialState,
	FieldTerminationState,
	FieldPastState,
	FieldInputs,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Flow type.
var ForeignKeys = []string{
	"survey_flow",
}

var (
	fields = schema.Flow{}.Fields()

	// descStateTable is the schema descriptor for stateTable field.
	descStateTable = fields[2].Descriptor()
	// StateTableValidator is a validator for the "stateTable" field. It is called by the builders before save.
	StateTableValidator = descStateTable.Validators[0].(func(string) error)
)
