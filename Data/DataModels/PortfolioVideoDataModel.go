package DataModels

import (
	"github.com/google/uuid"
)

type PortfolioVideoDataModel struct {
	tableName struct{} `sql:"Portfolio_Video"`
	Uuid uuid.UUID `sql:"id, pk"`
	PortfolioUuid uuid.UUID `sql:"portfolio_id, fk:Portfolio.id, notnull"`
	Video string `sql:"video, notnull"`
	SortOrder int `sql:"sort_order, notnull"`
}