package DataModels

import (
	"github.com/google/uuid"
)

type PortfolioVideoDataModel struct {
	tableName struct{} `sql:"portfolio_video"`
	Uuid uuid.UUID `sql:"id, pk"`
	PortfolioUuid uuid.UUID `sql:"portfolio_id, fk:portfolio.id, notnull"`
	Video string `sql:"video, notnull"`
	SortOrder int `sql:"sort_order, notnull"`
}