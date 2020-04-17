package drafts

import (
	"context"

	"github.com/google/uuid"
	"github.com/minskylab/collecta"
	"github.com/minskylab/collecta/ent"
)

func GenerateUTECDemo(ctx context.Context, db *collecta.DB, domainID uuid.UUID, userID uuid.UUID) (*ent.Survey, error) {
	return generateUtecDemoSurvey(ctx , db , domainID, userID )
}
