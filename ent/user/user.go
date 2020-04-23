// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"github.com/minskylab/collecta/ent/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldLastActivity holds the string denoting the lastactivity vertex property in the database.
	FieldLastActivity = "last_activity"
	// FieldUsername holds the string denoting the username vertex property in the database.
	FieldUsername = "username"
	// FieldPicture holds the string denoting the picture vertex property in the database.
	FieldPicture = "picture"
	// FieldRoles holds the string denoting the roles vertex property in the database.
	FieldRoles = "roles"

	// Table holds the table name of the user in the database.
	Table = "users"
	// AccountsTable is the table the holds the accounts relation/edge.
	AccountsTable = "accounts"
	// AccountsInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountsInverseTable = "accounts"
	// AccountsColumn is the table column denoting the accounts relation/edge.
	AccountsColumn = "user_accounts"
	// ContactsTable is the table the holds the contacts relation/edge.
	ContactsTable = "contacts"
	// ContactsInverseTable is the table name for the Contact entity.
	// It exists in this package in order to avoid circular dependency with the "contact" package.
	ContactsInverseTable = "contacts"
	// ContactsColumn is the table column denoting the contacts relation/edge.
	ContactsColumn = "user_contacts"
	// SurveysTable is the table the holds the surveys relation/edge.
	SurveysTable = "surveys"
	// SurveysInverseTable is the table name for the Survey entity.
	// It exists in this package in order to avoid circular dependency with the "survey" package.
	SurveysInverseTable = "surveys"
	// SurveysColumn is the table column denoting the surveys relation/edge.
	SurveysColumn = "user_surveys"
	// DomainsTable is the table the holds the domains relation/edge. The primary key declared below.
	DomainsTable = "domain_users"
	// DomainsInverseTable is the table name for the Domain entity.
	// It exists in this package in order to avoid circular dependency with the "domain" package.
	DomainsInverseTable = "domains"
	// AdminOfTable is the table the holds the adminOf relation/edge. The primary key declared below.
	AdminOfTable = "domain_admins"
	// AdminOfInverseTable is the table name for the Domain entity.
	// It exists in this package in order to avoid circular dependency with the "domain" package.
	AdminOfInverseTable = "domains"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldLastActivity,
	FieldUsername,
	FieldPicture,
	FieldRoles,
}

var (
	// DomainsPrimaryKey and DomainsColumn2 are the table columns denoting the
	// primary key for the domains relation (M2M).
	DomainsPrimaryKey = []string{"domain_id", "user_id"}
	// AdminOfPrimaryKey and AdminOfColumn2 are the table columns denoting the
	// primary key for the adminOf relation (M2M).
	AdminOfPrimaryKey = []string{"domain_id", "user_id"}
)

var (
	fields = schema.User{}.Fields()

	// descName is the schema descriptor for name field.
	descName = fields[1].Descriptor()
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator = descName.Validators[0].(func(string) error)

	// descLastActivity is the schema descriptor for lastActivity field.
	descLastActivity = fields[2].Descriptor()
	// DefaultLastActivity holds the default value on creation for the lastActivity field.
	DefaultLastActivity = descLastActivity.Default.(func() time.Time)
)
