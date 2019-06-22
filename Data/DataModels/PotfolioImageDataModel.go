package DataModels

import (
	"github.com/google/uuid"
)

type PortfolioImageDataModel struct {
	tableName struct{} `sql:"Portfolio_Image"`
	Uuid uuid.UUID `sql:"id, pk"`
	PortfolioUuid uuid.UUID `sql:"portfolio_id, fk:Portfolio.id, notnull"`
	Image string `sql:"image, notnull"`
	SortOrder int `sql:"sort_order, notnull"` 
}