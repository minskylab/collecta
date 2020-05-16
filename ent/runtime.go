// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/minskylab/collecta/ent/account"
	"github.com/minskylab/collecta/ent/answer"
	"github.com/minskylab/collecta/ent/contact"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/ent/person"
	"github.com/minskylab/collecta/ent/question"
	"github.com/minskylab/collecta/ent/schema"
	"github.com/minskylab/collecta/ent/short"
	"github.com/minskylab/collecta/ent/survey"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescSub is the schema descriptor for sub field.
	accountDescSub := accountFields[2].Descriptor()
	// account.SubValidator is a validator for the "sub" field. It is called by the builders before save.
	account.SubValidator = accountDescSub.Validators[0].(func(string) error)
	answerFields := schema.Answer{}.Fields()
	_ = answerFields
	// answerDescAt is the schema descriptor for at field.
	answerDescAt := answerFields[1].Descriptor()
	// answer.DefaultAt holds the default value on creation for the at field.
	answer.DefaultAt = answerDescAt.Default.(func() time.Time)
	contactFields := schema.Contact{}.Fields()
	_ = contactFields
	// contactDescValue is the schema descriptor for value field.
	contactDescValue := contactFields[2].Descriptor()
	// contact.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	contact.ValueValidator = contactDescValue.Validators[0].(func(string) error)
	// contactDescFromAccount is the schema descriptor for fromAccount field.
	contactDescFromAccount := contactFields[6].Descriptor()
	// contact.DefaultFromAccount holds the default value on creation for the fromAccount field.
	contact.DefaultFromAccount = contactDescFromAccount.Default.(bool)
	domainFields := schema.Domain{}.Fields()
	_ = domainFields
	// domainDescName is the schema descriptor for name field.
	domainDescName := domainFields[1].Descriptor()
	// domain.NameValidator is a validator for the "name" field. It is called by the builders before save.
	domain.NameValidator = domainDescName.Validators[0].(func(string) error)
	// domainDescEmail is the schema descriptor for email field.
	domainDescEmail := domainFields[2].Descriptor()
	// domain.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	domain.EmailValidator = domainDescEmail.Validators[0].(func(string) error)
	flowFields := schema.Flow{}.Fields()
	_ = flowFields
	// flowDescStateTable is the schema descriptor for stateTable field.
	flowDescStateTable := flowFields[2].Descriptor()
	// flow.StateTableValidator is a validator for the "stateTable" field. It is called by the builders before save.
	flow.StateTableValidator = flowDescStateTable.Validators[0].(func(string) error)
	inputFields := schema.Input{}.Fields()
	_ = inputFields
	// inputDescMultiple is the schema descriptor for multiple field.
	inputDescMultiple := inputFields[2].Descriptor()
	// input.DefaultMultiple holds the default value on creation for the multiple field.
	input.DefaultMultiple = inputDescMultiple.Default.(bool)
	personFields := schema.Person{}.Fields()
	_ = personFields
	// personDescName is the schema descriptor for name field.
	personDescName := personFields[1].Descriptor()
	// person.NameValidator is a validator for the "name" field. It is called by the builders before save.
	person.NameValidator = personDescName.Validators[0].(func(string) error)
	// personDescLastActivity is the schema descriptor for lastActivity field.
	personDescLastActivity := personFields[2].Descriptor()
	// person.DefaultLastActivity holds the default value on creation for the lastActivity field.
	person.DefaultLastActivity = personDescLastActivity.Default.(func() time.Time)
	questionFields := schema.Question{}.Fields()
	_ = questionFields
	// questionDescTitle is the schema descriptor for title field.
	questionDescTitle := questionFields[2].Descriptor()
	// question.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	question.TitleValidator = questionDescTitle.Validators[0].(func(string) error)
	// questionDescAnonymous is the schema descriptor for anonymous field.
	questionDescAnonymous := questionFields[6].Descriptor()
	// question.DefaultAnonymous holds the default value on creation for the anonymous field.
	question.DefaultAnonymous = questionDescAnonymous.Default.(bool)
	shortFields := schema.Short{}.Fields()
	_ = shortFields
	// shortDescKey is the schema descriptor for key field.
	shortDescKey := shortFields[0].Descriptor()
	// short.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	short.KeyValidator = shortDescKey.Validators[0].(func(string) error)
	surveyFields := schema.Survey{}.Fields()
	_ = surveyFields
	// surveyDescDueDate is the schema descriptor for dueDate field.
	surveyDescDueDate := surveyFields[3].Descriptor()
	// survey.DefaultDueDate holds the default value on creation for the dueDate field.
	survey.DefaultDueDate = surveyDescDueDate.Default.(func() time.Time)
	// surveyDescTitle is the schema descriptor for title field.
	surveyDescTitle := surveyFields[4].Descriptor()
	// survey.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	survey.TitleValidator = surveyDescTitle.Validators[0].(func(string) error)
	// surveyDescDone is the schema descriptor for done field.
	surveyDescDone := surveyFields[7].Descriptor()
	// survey.DefaultDone holds the default value on creation for the done field.
	survey.DefaultDone = surveyDescDone.Default.(bool)
	// surveyDescIsPublic is the schema descriptor for isPublic field.
	surveyDescIsPublic := surveyFields[8].Descriptor()
	// survey.DefaultIsPublic holds the default value on creation for the isPublic field.
	survey.DefaultIsPublic = surveyDescIsPublic.Default.(bool)
}