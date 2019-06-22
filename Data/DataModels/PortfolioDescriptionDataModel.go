package DataModels

import (
	"github.com/google/uuid"
)

type PortfolioDescriptionDataModel struct {
	tableName struct{} `sql:"portfolio_description"`
	Uuid uuid.UUID `sql:"id, pk"`
	PortfolioUuid uuid.UUID `sql:"portfolio_id, fk:portfolio.id, notnull"`
	Content string `sql:"content, notnull"`
	SortOrder int `sql:"sort_order, notnull"`
}