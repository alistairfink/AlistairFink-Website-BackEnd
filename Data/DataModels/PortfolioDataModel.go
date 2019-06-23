package DataModels

import (
	"time"
	"github.com/google/uuid"
)

type PortfolioDataModel struct {
	tableName struct{} `sql:"portfolio"`
	Uuid uuid.UUID `sql:"id, pk"`
	Name string `sql:"name, notnull"`
	Thumbnail string `sql:"thumbnail, notnull"`
	Year time.Time `sql:"year, notnull"`
}