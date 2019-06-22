package DataModels

import (
	"github.com/google/uuid"
)

type AboutDescriptionDataModel struct {
	tableName struct{} `sql:"about_description"`
	Uuid uuid.UUID `sql:"id, pk"`
	AboutUuid uuid.UUID `sql:"about_id, fk:About.id, notnull"`
	Content string `sql:"content, notnull"`
	SortOrder int `sql:"sort_order, notnull"`
}