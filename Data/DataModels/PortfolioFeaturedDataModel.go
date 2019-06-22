package DataModels

import (
	"github.com/google/uuid"
)

type PortfolioFeaturedDataModel struct {
	tableName struct{} `sql:"portfolio_featured"`
	PortfolioUuid uuid.UUID `sql:"portfolio_id, pk, fk:portfolio.id"`
	SortOrder int `sql:"sort_order, notnull"`
}