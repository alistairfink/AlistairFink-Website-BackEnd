package DataModels

import (
	"github.com/google/uuid"
)

type EducationScholarshipDataModel struct {
	tableName struct{} `sql:"education_scholarship"`
	Uuid uuid.UUID `sql:"id, pk"`
	Educationuuid uuid.UUID `sql:"education_id, fk:education.id, notnull"`
	Name string `sql:"name, notnull"`
	SortOrder int `sql:"sort_order, notnull"`
}