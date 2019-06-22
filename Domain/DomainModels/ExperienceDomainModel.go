package DomainModels

import (
	"time"
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

type ExperienceDomainModel struct {
	Uuid uuid.UUID
	Position string
	StartDate time.Time
	EndDate time.Time
	Company string
	Location string
	LogoImage string
	Content *[]DataModels.ExperienceContentDataModel
}

func (this *ExperienceDomainModel) ToDomainModel (
	experienceDataModel *DataModels.ExperienceDataModel,
	contentsDataModels *[]DataModels.ExperienceContentDataModel,
) {
	this.Uuid = experienceDataModel.Uuid
	this.Position = experienceDataModel.Position
	this.StartDate = experienceDataModel.StartDate
	this.EndDate = experienceDataModel.EndDate
	this.Company = experienceDataModel.Company
	this.Location = experienceDataModel.Location
	this.LogoImage = experienceDataModel.LogoImage
	this.Content = contentsDataModels
}

func (this *ExperienceDomainModel) ToDataModel() (*DataModels.ExperienceDataModel) {
	return &DataModels.ExperienceDataModel {
		Uuid: this.Uuid,
		Position: this.Position,
		StartDate: this.StartDate,
		EndDate: this.EndDate,
		Company: this.Company,
		Location: this.Location,
		LogoImage: this.LogoImage,
	}
}