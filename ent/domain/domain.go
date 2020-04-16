// Code generated by entc, DO NOT EDIT.

package domain

import (
	"github.com/minskylab/collecta/ent/schema"
)

const (
	// Label holds the string label denoting the domain type in the database.
	Label = "domain"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTags holds the string denoting the tags vertex property in the database.
	FieldTags = "tags"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email vertex property in the database.
	FieldEmail = "email"
	// FieldDomain holds the string denoting the domain vertex property in the database.
	FieldDomain = "domain"

	// Table holds the table name of the domain in the database.
	Table = "domains"
	// SurveysTable is the table the holds the surveys relation/edge.
	SurveysTable = "surveys"
	// SurveysInverseTable is the table name for the Survey entity.
	// It exists in this package in order to avoid circular dependency with the "survey" package.
	SurveysInverseTable = "surveys"
	// SurveysColumn is the table column denoting the surveys relation/edge.
	SurveysColumn = "domain_surveys"
	// UsersTable is the table the holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "domain_users"
)

// Columns holds all SQL columns for domain fields.
var Columns = []string{
	FieldID,
	FieldTags,
	FieldName,
	FieldEmail,
	FieldDomain,
}

var (
	fields = schema.Domain{}.Fields()

	// descName is the schema descriptor for name field.
	descName = fields[2].Descriptor()
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator = descName.Validators[0].(func(string) error)

	// descEmail is the schema descriptor for email field.
	descEmail = fields[3].Descriptor()
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator = descEmail.Validators[0].(func(string) error)
)
