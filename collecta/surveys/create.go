package surveys

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
	"github.com/minskylab/collecta/api/graph/model"
	"github.com/minskylab/collecta/collecta/flows"
	"github.com/minskylab/collecta/db"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/ent/input"
	"github.com/minskylab/collecta/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func createQuestion(ctx context.Context, db *db.DB, q model.QuestionCreator) (*ent.Question, error) {
	policy := bluemonday.UGCPolicy()
	policy.AllowElements("h1").AllowElements("h2").AllowElements("h3")
	policy.AllowAttrs("href").OnElements("a")
	policy.AllowElements("p")
	// TODO: Sanitize Policy should be into core struct (Collecta)

	an := false
	if q.Anonymous != nil {
		an = *q.Anonymous
	}

	h := md5.Sum([]byte(q.Title))
	hash := hex.EncodeToString(h[:])
	newQuestion, err := db.Ent.Question.Create().
		SetID(uuid.New()).
		SetTitle(policy.Sanitize(q.Title)).
		SetDescription(policy.Sanitize(q.Description)).
		SetAnonymous(an).
		SetValidator("").
		SetHash(hash).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to create an input")
	}

	var inType input.Kind
	switch q.Kind {
	case model.InputTypeText:
		inType = input.KindText
	case model.InputTypeOption:
		inType = input.KindOptions
	case model.InputTypeSatisfaction:
		inType = input.KindSatisfaction
	case model.InputTypeBoolean:
		inType = input.KindBoolean
	default:
		return nil, errors.New("invalid input type")
	}

	isMultiple := false
	if q.Multiple != nil {
		isMultiple = *q.Multiple
	}

	options := map[string]string{}

	for _, pair := range q.Options {
		options[pair.Key] = policy.Sanitize(pair.Value)
	}

	_, err = db.Ent.Input.Create().
		SetID(uuid.New()).
		SetKind(inType).
		SetMultiple(isMultiple).
		SetOptions(options).
		SetQuestion(newQuestion).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at try to create an input")
	}

	return newQuestion, nil
}

func createSurveysFromAPI(ctx context.Context, db *db.DB, domainID uuid.UUID, draft model.SurveyGenerator) ([]*ent.Survey, error) {
	if len(draft.Questions) == 0 {
		return nil, errors.New("invalid survey, you need specify questions")
	}
	log.Info("creating surveys")

	spew.Dump(draft)

	questions := make([]*ent.Question, 0)

	for _, q := range draft.Questions {
		if q != nil {
			newQ, err := createQuestion(ctx, db, *q)
			if err != nil {
				for _, rbQuestion := range questions {
					if err := db.Ent.Question.DeleteOneID(rbQuestion.ID).Exec(ctx); err != nil {
						return nil, errors.Wrap(err, "error at delete created question")
					}
				}
				return nil, errors.Wrap(err, "error at create one of your question, roll-backing operation")
			}
			questions = append(questions, newQ)
		}
	}

	// generating the flow program
	// by default, the logic is a ordered sequential by one by one question
	questionIDs := make([]uuid.UUID, 0)
	for _, q := range questions {
		questionIDs = append(questionIDs, q.ID)
	}

	flowProgram := flows.DefaultSequentialProgram(questionIDs)
	if draft.Logic != nil {
		flowProgram = *draft.Logic
	}

	// TODO: Improve this part of the Collecta API
	basicFlow, err := db.Ent.Flow.Create().
		SetID(uuid.New()).
		SetInputs([]string{}).
		SetState(questions[0].ID).                           // first question
		SetInitialState(questions[0].ID).                    // first question, TODO
		SetTerminationState(questions[len(questions)-1].ID). // last question, TODO
		SetStateTable(flowProgram).
		AddQuestions(questions...).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error at create flow")
	}

	if draft.Target == nil {
		return nil, errors.New("users target not specified")
	}

	audience := make([]*ent.User, 0)

	switch draft.Target.TargetKind {
	case model.SurveyAudenceKindPublic:
		break
	case model.SurveyAudenceKindDomain:
		d, err := db.Ent.Domain.Get(ctx, domainID)
		if err != nil {
			return nil, errors.Wrap(err, "error at fetch domain")
		}
		audience, err = d.QueryUsers().All(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "error at get all users of your domain")
		}
	case model.SurveyAudenceKindClose:
		if len(draft.Target.Whitelist) == 0 {
			return nil, errors.New("you cannot create a close survey without specify the whitelist")
		}
		for _, uID := range draft.Target.Whitelist {
			userID, err := uuid.Parse(uID)
			if err != nil {
				return nil, errors.Wrap(err, "error at parse one user ID from your audience")
			}

			u, err := db.Ent.User.Get(ctx, userID)
			if err != nil {
				return nil, errors.Wrap(err, "error at get user ")
			}

			audience = append(audience, u)
		}
	default:
		return nil, errors.New("invalid target kind for your survey")
	}

	policy := bluemonday.UGCPolicy()
	policy.AllowElements("h1").AllowElements("h2").AllowElements("h3")
	policy.AllowAttrs("href").OnElements("a")
	policy.AllowElements("p")
	// TODO: Sanitize Policy should be into core struct (Collecta)

	tags := make([]string, 0)

	for _, t := range draft.Tags {
		tags = append(tags, policy.Sanitize(t))
	}

	metadata := map[string]string{}
	for _, pair := range draft.Metadata {
		metadata[pair.Key] = pair.Value
	}

	generatedSurveys := make([]*ent.Survey, 0)
	for _, u := range audience {
		newSurvey, err := db.Ent.Survey.Create().
			SetID(uuid.New()).
			SetForID(u.ID).
			SetOwnerID(domainID).
			SetDone(false).
			SetTitle(policy.Sanitize(draft.Title)).
			SetDescription(policy.Sanitize(draft.Description)).
			SetLastInteraction(time.Now()).
			SetTags(tags).
			SetFlow(basicFlow).
			SetMetadata(metadata).
			Save(ctx)
		if err != nil {
			for _, s := range generatedSurveys {
				if err := db.Ent.Survey.DeleteOneID(s.ID).Exec(ctx); err != nil {
					return nil, errors.Wrap(err, "error at perform roll-backing")
				}
			}
			return nil, errors.Wrap(err, "error at try to create your survey, your operation will be roll-backed")
		}

		generatedSurveys = append(generatedSurveys, newSurvey)
	}

	return generatedSurveys, nil
}

func GenerateSurveys(ctx context.Context, db *db.DB, domainID uuid.UUID, draft model.SurveyGenerator) ([]*ent.Survey, error) {
	return createSurveysFromAPI(ctx, db, domainID, draft)
}
