package drafts

import (
	"context"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/collecta/commons"
	"github.com/minskylab/collecta/collecta/surveys"
	"github.com/minskylab/collecta/db"
	"github.com/minskylab/collecta/ent"
	"github.com/pkg/errors"
)

func generateUTECDemoSurvey(ctx context.Context, db *db.DB, domainID uuid.UUID, userID uuid.UUID) (*ent.Survey, error) {
	gen := model.SurveyGenerator{
		Title:       "Feedback | Teoría de Decisiones",
		Description: "<br>{{.Name}}</br>, responde esta pequeña encuesta sobre tu clase de Teoría de Decisiones del día martes 10 de Marzo.",
		Tags:        []string{"Teo1.01", "UTEC", "2020-1"},
		Questions:   []*model.QuestionCreator{
			{
				Title:       "Dinamicas en clase",
				Description: "Selecciona las dinámicas que usó el profesor que aportaron a la clase.",
				Kind:        model.InputTypeOption,
				Multiple:    commons.PtrBool(true),
				Anonymous:   commons.PtrBool(false),
				Options:     commons.MapToPairs(map[string]string{
							"rooms":  "Breakout Rooms",
							"chat":   "Chat",
							"game":   "Juegos-Simulación",
							"kahoot": "Kahoot u otra encuesta",
							"others": "Otra",
						}),
			},
			{
				Title:       "Acerca de la conectividad",
				Description: "Por problemas de conectividad, ¿crees que es necesario repetir la clase?",
				Kind:        model.InputTypeBoolean,
			},
			{
				Title:       "Satisfacción personal",
				Description: "¿Que tan provechosa fue esta clase para tu aprendizaje?",
				Kind:        model.InputTypeSatisfaction,
			},
			{
				Title:       "Opinion Extra",
				Description: "Si tienes algún comentario extra, por favor escríbelo a continuación.",
				Kind:        model.InputTypeText,
			},
		},
		Target:      &model.SurveyTargetUsers{
			TargetKind: model.SurveyAudenceKindClose,
			Whitelist:  []string{userID.String()},
		},
		Metadata:    commons.MapToPairs(map[string]string{
			"type": "example",
			"classroom": "S001",
		}),
	}

	s, err := surveys.GenerateSurveys(ctx, db, domainID, gen)
	if err != nil {
		return nil, errors.Wrap(err, "error at generate surveys")
	}

	if len(s) < 1 {
		return nil, errors.New("problem occurred at generate your survey")
	}

	return s[0], nil
}
