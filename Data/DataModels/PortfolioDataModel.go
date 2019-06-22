package DataModels

import (
	"time"
	"github.com/google/uuid"
)

type PortfolioDataModel struct {
	tableName struct{} `sql:"portfolio"`
	Uuid uuid.UUID `sql:"id, pk"`
	Name string `sql:"string, notnull"`
	Thumbnail string `sql:"string, notnull"`
	Year time.Time `sql:"year, notnull"`
}