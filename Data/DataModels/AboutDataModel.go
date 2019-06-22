package DataModels

import (
	"github.com/google/uuid"
)

type AboutDataModel struct {
	tableName struct{} `sql:"about"`
	Uuid uuid.UUID `sql:"id, pk"`
	Image string `sql:"image, notnull"`
}