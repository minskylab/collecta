// Code generated by entc, DO NOT EDIT.

package input

import (
	"fmt"

	"github.com/minskylab/collecta/ent/schema"
)

const (
	// Label holds the string label denoting the input type in the database.
	Label = "input"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKind holds the string denoting the kind vertex property in the database.
	FieldKind = "kind"
	// FieldMultiple holds the string denoting the multiple vertex property in the database.
	FieldMultiple = "multiple"
	// FieldDefaults holds the string denoting the defaults vertex property in the database.
	FieldDefaults = "defaults"
	// FieldOptions holds the string denoting the options vertex property in the database.
	FieldOptions = "options"

	// Table holds the table name of the input in the database.
	Table = "inputs"
	// QuestionTable is the table the holds the question relation/edge.
	QuestionTable = "inputs"
	// QuestionInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionInverseTable = "questions"
	// QuestionColumn is the table column denoting the question relation/edge.
	QuestionColumn = "question_input"
)

// Columns holds all SQL columns for input fields.
var Columns = []string{
	FieldID,
	FieldKind,
	FieldMultiple,
	FieldDefaults,
	FieldOptions,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Input type.
var ForeignKeys = []string{
	"question_input",
}

var (
	fields = schema.Input{}.Fields()

	// descMultiple is the schema descriptor for multiple field.
	descMultiple = fields[2].Descriptor()
	// DefaultMultiple holds the default value on creation for the multiple field.
	DefaultMultiple = descMultiple.Default.(bool)
)

// Kind defines the type for the kind enum field.
type Kind string

// Kind values.
const (
	KindText         Kind = "Text"
	KindOptions      Kind = "Options"
	KindSatisfaction Kind = "Satisfaction"
	KindBoolean      Kind = "Boolean"
)

func (s Kind) String() string {
	return string(s)
}

// KindValidator is a validator for the "k" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindText, KindOptions, KindSatisfaction, KindBoolean:
		return nil
	default:
		return fmt.Errorf("input: invalid enum value for kind field: %q", k)
	}
}
