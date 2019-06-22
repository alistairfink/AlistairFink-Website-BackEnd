package DataModels

import (
	"github.com/google/uuid"
)

type ExperienceContentDataModel struct {
	tableName struct{} `sql:"Experience_Content"`
	Uuid uuid.UUID `sql:"id, pk"`
	ExperienceUuid uuid.UUID `sql:"experience_id, fk:Experience.id, notnull"`
	Content string `sql:"content, notnull"`
	SortOrder int `sql:"sort_order, notnull"`
}