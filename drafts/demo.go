package drafts

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/minskylab/collecta"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/input"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type basicQuestion struct {
	questionTitle string
	questionDescription string
	inputType input.Kind
	isMultiple bool
	options map[string]string
}

func genQuestion(ctx context.Context, db *collecta.DB, q basicQuestion) (*ent.Question, error) {
	newQuestion, err := db.Ent.Question.Create().
		SetID(uuid.New()).
		SetTitle(q.questionTitle).
		SetDescription(q.questionDescription).
		SetAnonymous(false).
		SetHash(q.questionTitle).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to create an input")
	}

	_, err = db.Ent.Input.Create().
		SetID(uuid.New()).
		SetKind(q.inputType).
		SetMultiple(q.isMultiple).
		SetOptions(q.options).
		SetQuestion(newQuestion).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to create an input")
	}

	return newQuestion, nil
}


func generateUTECDemoSurvey(ctx context.Context, db *collecta.DB, domainID uuid.UUID, userID uuid.UUID) (*ent.Survey, error) {
	q1, err := genQuestion(ctx, db, basicQuestion{
		questionTitle:      "PREGUNTA 1/8",
		questionDescription: "Selecciona las dinámicas que usó el profesor que aportaron a la clase.",
		inputType:           input.KindOptions,
		isMultiple:          true,
		options:             map[string]string{
			"rooms": "Breakout Rooms",
			"chat": "Chat",
			"game": "Juegos-Simulación",
			"kahoot": "Kahoot u otra encuesta",
			"others": "Otra",
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "error at gen question q1")
	}
	log.Info("qi= ", q1.ID)

	q2, err := genQuestion(ctx, db, basicQuestion{
		questionTitle:      "PREGUNTA 2/8",
		questionDescription: "Por problemas de conectividad, ¿crees que es necesario repetir la clase?",
		inputType:           input.KindBoolean,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error at gen question q2")
	}
	log.Info("qi= ", q2.ID)

	q3, err := genQuestion(ctx, db, basicQuestion{
		questionTitle:      "PREGUNTA 3/8",
		questionDescription: "Si tienes algún comentario extra, por favor escríbelo a continuación.",
		inputType:           input.KindText,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error at gen question q3")
	}
	log.Info("qi= ", q3.ID)

	q4, err := genQuestion(ctx, db, basicQuestion{
		questionTitle:      "PREGUNTA 4/8",
		questionDescription: "¿Que tan provechosa fue esta clase para tu aprendizaje?",
		inputType:           input.KindSatisfaction,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error at gen questionq4")
	}
	log.Info("qi= ", q4.ID)

	flowProgram :=
	 	"00 " + q1.ID.String() + " -> " + q2.ID.String() +
		"01 " + q2.ID.String() + " -> " + q3.ID.String() +
		"02 " + q3.ID.String() + " -> " + q4.ID.String()


	qFlow, err := db.Ent.Flow.Create().
		SetID(uuid.New()).
		SetInputs([]string{}).
		SetState(q1.ID).
		SetStateTable(flowProgram).
		AddQuestions(
			q1, q2, q3, q4,
		).
		Save(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error at create flow")
	}

	log.Info("flow= ", qFlow.ID)

	return db.Ent.Survey.Create().
		SetID(uuid.New()).
		SetOwnerID(domainID).
		SetTitle("Feedback por Sesión | Estudiantes").
		SetDescription("<br>{{.Name}}</br>, responde esta pequeña encuesta sobre tu clase de Teoría de Decisiones del día martes 10 de Marzo.").
		SetForID(userID).
		SetDone(false).
		SetLastInteraction(time.Now()).
		SetTags([]string{"demo", "collecta", "utec", "teoriadedecisiones"}).
		SetFlow(qFlow).
		SetMetadata(map[string]string{
			"creator": "Collecta Labs",
		}).
		Save(ctx)
}