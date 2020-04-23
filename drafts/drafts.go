package drafts

import (
	"context"

	"github.com/google/uuid"
	"github.com/minskylab/collecta/db"
	"github.com/minskylab/collecta/ent"
)

func GenerateUTECDemo(ctx context.Context, collectaDB *db.DB, domainID uuid.UUID, userID uuid.UUID) (*ent.Survey, error) {
	return generateUTECDemoSurvey(ctx, collectaDB, domainID, userID)
}
