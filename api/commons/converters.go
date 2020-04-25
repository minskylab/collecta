package commons

import (
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/ent"
)

func DomainToGQL(e *ent.Domain) *model.Domain {
	return &model.Domain{
		ID:             e.ID.String(),
		Tags:           e.Tags,
		Name:           e.Name,
		Email:          e.Email,
		Domain:         e.Domain,
		CollectaDomain: e.CollectaDomain,
	}
}

func UserToGQL(e *ent.User) *model.User {
	return &model.User{
		ID:           e.ID.String(),
		Name:         e.Name,
		Username:     e.Username,
		LastActivity: e.LastActivity,
		Picture:      e.Picture,
		Roles:        e.Roles,
	}
}

func AccountToGQL(e *ent.Account) *model.Account {
	return &model.Account{
		ID:       e.ID.String(),
		Type:     e.Type.String(),
		Sub:      e.Sub,
		RemoteID: e.RemoteID,
		// Secret: e.Secret,
	}
}

func ContactToGQL(e *ent.Contact) *model.Contact {
	return &model.Contact{
		ID:          e.ID.String(),
		Name:        &e.Name,
		Value:       e.Value,
		Kind:        e.Kind.String(),
		Principal:   e.Principal,
		Validated:   e.Validated,
		FromAccount: e.FromAccount,
	}
}

func FlowToGQL(e *ent.Flow) *model.Flow {
	return &model.Flow{
		ID:               e.ID.String(),
		State:            e.State.String(),
		StateTable:       e.StateTable,
		Inputs:           e.Inputs,
		InitialState:     e.InitialState.String(),
		TerminationState: e.TerminationState.String(),
		PastState:        e.PastState.String(),
	}
}

func QuestionToGQL(e *ent.Question) *model.Question {
	meta := map[string]interface{}{}
	for k, v := range e.Metadata {
		meta[k] = v
	}
	return &model.Question{
		ID:          e.ID.String(),
		Hash:        e.Hash,
		Title:       e.Title,
		Description: e.Description,
		Anonymous:   e.Anonymous,
		Validator:   e.Validator,
		Metadata:    meta,
	}
}

func AnswerToGQL(e *ent.Answer) *model.Answer {
	return &model.Answer{
		ID:        e.ID.String(),
		At:        e.At,
		Responses: e.Responses,
		Valid:     e.Valid,
	}
}

func InputToGQL(e *ent.Input) *model.Input {
	defs := make([]*string, 0)
	for _, s := range e.Defaults {
		defs = append(defs, &s)
	}

	opts := map[string]interface{}{}

	for k, v := range e.Options {
		opts[k] = v
	}

	return &model.Input{
		ID:       e.ID.String(),
		Kind:     e.Kind.String(),
		Multiple: e.Multiple,
		Defaults: defs,
		Options:  opts,
	}
}

func SurveyToGQL(e *ent.Survey) *model.Survey {
	meta := map[string]interface{}{}
	for k, v := range e.Metadata {
		meta[k] = v
	}
	return &model.Survey{
		ID:              e.ID.String(),
		Tags:            e.Tags,
		LastInteraction: e.LastInteraction,
		DueDate:         e.DueDate,
		Title:           e.Title,
		Description:     e.Description,
		Done:            e.Done,
		IsPublic:        e.IsPublic,
		Metadata:        meta,
	}
}
