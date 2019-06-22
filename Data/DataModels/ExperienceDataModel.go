package DataModels

import (
	"time"
	"github.com/google/uuid"
)

type ExperienceDataModel struct {
	tableName struct{} `sql:"experience"`
	Uuid uuid.UUID `sql:"id, pk"`
	Position string `sql:"position, notnull"`
	StartDate time.Time `sql:"start_date, notnull"`
	EndDate time.Time `sql:"end_date"`
	Company string `sql:"company, notnull"`
	Location string `sql:"location, notnull"`
	LogoImage string `sql:"logo_image, notnull"`
}