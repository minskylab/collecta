// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Account struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Sub      string `json:"sub"`
	RemoteID string `json:"remoteID"`
	Secret   string `json:"secret"`
	Owner    *User  `json:"owner"`
}

type Answer struct {
	ID        string    `json:"id"`
	At        time.Time `json:"at"`
	Responses []string  `json:"responses"`
	Valid     bool      `json:"valid"`
	Question  *Question `json:"question"`
}

type Contact struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Value       string  `json:"value"`
	Kind        string  `json:"kind"`
	Principal   bool    `json:"principal"`
	Validated   bool    `json:"validated"`
	FromAccount bool    `json:"fromAccount"`
	Owner       *User   `json:"owner"`
}

type Domain struct {
	ID             string    `json:"id"`
	Tags           []string  `json:"tags"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Domain         string    `json:"domain"`
	CollectaDomain string    `json:"collectaDomain"`
	Surveys        []*Survey `json:"surveys"`
	Users          []*User   `json:"users"`
	Admins         []*User   `json:"admins"`
}

type DomainCreator struct {
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Domain         string   `json:"domain"`
	CollectaDomain string   `json:"collectaDomain"`
	Tags           []string `json:"tags"`
}

type Flow struct {
	ID               string      `json:"id"`
	State            string      `json:"state"`
	StateTable       string      `json:"stateTable"`
	InitialState     string      `json:"initialState"`
	TerminationState string      `json:"terminationState"`
	PastState        string      `json:"pastState"`
	Inputs           []string    `json:"inputs"`
	Questions        []*Question `json:"questions"`
}

type Input struct {
	ID       string                 `json:"id"`
	Kind     string                 `json:"kind"`
	Multiple bool                   `json:"multiple"`
	Defaults []*string              `json:"defaults"`
	Options  map[string]interface{} `json:"options"`
	Question *Question              `json:"question"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type MetadataPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Question struct {
	ID          string          `json:"id"`
	Hash        string          `json:"hash"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Anonymous   bool            `json:"anonymous"`
	Metadata    []*MetadataPair `json:"metadata"`
	Validator   string          `json:"validator"`
	Answers     []*Answer       `json:"answers"`
	Input       *Input          `json:"input"`
	Flow        *Flow           `json:"flow"`
}

type QuestionCreator struct {
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Kind        InputType              `json:"kind"`
	Multiple    *bool                  `json:"multiple"`
	Anonymous   *bool                  `json:"anonymous"`
	Options     map[string]interface{} `json:"options"`
}

type Short struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Survey struct {
	ID              string          `json:"id"`
	Tags            []string        `json:"tags"`
	LastInteraction time.Time       `json:"lastInteraction"`
	DueDate         time.Time       `json:"dueDate"`
	Title           string          `json:"title"`
	Description     string          `json:"description"`
	Metadata        []*MetadataPair `json:"metadata"`
	Done            bool            `json:"done"`
	IsPublic        bool            `json:"isPublic"`
	Flow            *Flow           `json:"flow"`
	For             *User           `json:"for"`
	Owner           *Domain         `json:"owner"`
}

type SurveyDomain struct {
	ByID         *string `json:"byID"`
	ByDomainName *string `json:"byDomainName"`
}

type SurveyGenerator struct {
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Tags        []string               `json:"tags"`
	Questions   []*QuestionCreator     `json:"questions"`
	Target      *SurveyTargetUsers     `json:"target"`
	Metadata    map[string]interface{} `json:"metadata"`
	Logic       *string                `json:"logic"`
}

type SurveyTargetUsers struct {
	TargetKind SurveyAudenceKind `json:"targetKind"`
	Whitelist  []string          `json:"whitelist"`
}

type SuveyGenerationResult struct {
	How     int       `json:"how"`
	Surveys []*Survey `json:"surveys"`
}

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	LastActivity time.Time `json:"lastActivity"`
	Picture      string    `json:"picture"`
	Roles        []string  `json:"roles"`
	Accounts     *Account  `json:"accounts"`
	Contacts     *Contact  `json:"contacts"`
	Surveys      []*Survey `json:"surveys"`
	Domains      []*Domain `json:"domains"`
	AdminOf      []*Domain `json:"adminOf"`
}

type InputType string

const (
	InputTypeOption       InputType = "OPTION"
	InputTypeText         InputType = "TEXT"
	InputTypeBoolean      InputType = "BOOLEAN"
	InputTypeSatisfaction InputType = "SATISFACTION"
)

var AllInputType = []InputType{
	InputTypeOption,
	InputTypeText,
	InputTypeBoolean,
	InputTypeSatisfaction,
}

func (e InputType) IsValid() bool {
	switch e {
	case InputTypeOption, InputTypeText, InputTypeBoolean, InputTypeSatisfaction:
		return true
	}
	return false
}

func (e InputType) String() string {
	return string(e)
}

func (e *InputType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = InputType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid InputType", str)
	}
	return nil
}

func (e InputType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SurveyAudenceKind string

const (
	SurveyAudenceKindPublic SurveyAudenceKind = "PUBLIC"
	SurveyAudenceKindDomain SurveyAudenceKind = "DOMAIN"
	SurveyAudenceKindClose  SurveyAudenceKind = "CLOSE"
)

var AllSurveyAudenceKind = []SurveyAudenceKind{
	SurveyAudenceKindPublic,
	SurveyAudenceKindDomain,
	SurveyAudenceKindClose,
}

func (e SurveyAudenceKind) IsValid() bool {
	switch e {
	case SurveyAudenceKindPublic, SurveyAudenceKindDomain, SurveyAudenceKindClose:
		return true
	}
	return false
}

func (e SurveyAudenceKind) String() string {
	return string(e)
}

func (e *SurveyAudenceKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SurveyAudenceKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SurveyAudenceKind", str)
	}
	return nil
}

func (e SurveyAudenceKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
