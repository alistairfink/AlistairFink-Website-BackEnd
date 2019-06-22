package DataModels

import (
	"github.com/google/uuid"
)

type EducationExtracurricularDataModel struct {
	tableName struct{} `sql:"education_extracurricular"`
	Uuid uuid.UUID `sql:"id, pk"`
	EducationUuid uuid.UUID `sql:"education_id, fk:Education.id, notnull"`
	Name string `sql:"name, notnull"`
	SortOrder int `sql:"sort_order, notnull"`
}