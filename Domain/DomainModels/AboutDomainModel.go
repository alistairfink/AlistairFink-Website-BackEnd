package DomainModels

import (
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

type AboutDomainModel struct {
	Uuid uuid.UUID
	Image string
	Description *[]DataModels.AboutDescriptionDataModel
}

func (this *AboutDomainModel) ToDomainModel (
	aboutDataModel *DataModels.AboutDataModel, 
	aboutDescriptionDataModels *[]DataModels.AboutDescriptionDataModel,
) {
	this.Uuid = aboutDataModel.Uuid
	this.Image = aboutDataModel.Image
	this.Description = aboutDescriptionDataModels
}

func (this *AboutDomainModel) ToDataModel() (*DataModels.AboutDataModel) {
	return &DataModels.AboutDataModel {
		Uuid: this.Uuid,
		Image: this.Image,
	}
}