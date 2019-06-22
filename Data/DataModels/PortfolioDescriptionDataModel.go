package DataModels

import (
	"github.com/google/uuid"
)

type PortfolioDescriptionDataModel struct {
	tableName struct{} `sql:"Portfolio_Description"`
	Uuid uuid.UUID `sql:"id, pk"`
	PortfolioUuid uuid.UUID `sql:"portfolio_id, fk:Portfolio.id, notnull"`
	Content string `sql:"content, notnull"`
	SortOrder int `sql:"sort_order, notnull"`
}