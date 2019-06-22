package DataModels

import (
	"time"
	"github.com/google/uuid"
)

type EducationDataModel struct {
	tableName struct{} `sql:"Education"`
	Uuid uuid.UUID `sql:"id, pk"`
	SchoolName string `sql:"school_name, notnull"`
	StartDate time.Time `sql:"start_date, notnull"`
	EndDate time.Time `sql:"end_date"`
	Location string `sql:"location, notnull"`
	Certification string `sql:"certification, notnull"`
	SchoolLogo string `sql:"school_logo, notnull"` 
}