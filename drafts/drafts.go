package drafts

import (
	"context"

	"github.com/minskylab/collecta/db"
	"github.com/minskylab/collecta/ent"
	"github.com/minskylab/collecta/uuid"
)

func GenerateUTECDemo(ctx context.Context, collectaDB *db.DB, domainID uuid.UUID, userID uuid.UUID) (*ent.Survey, error) {
	return generateUTECDemoSurvey(ctx, collectaDB, domainID, userID)
}
