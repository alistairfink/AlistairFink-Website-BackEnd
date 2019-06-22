package DomainModels

import (
	"time"
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

type PortfolioDomainModel struct {
	Uuid uuid.UUID
	Name string
	Thumbnail string
	Year time.Time
	Description *[]DataModels.PortfolioDescriptionDataModel
	Image *[]DataModels.PortfolioImageDataModel
	Video *[]DataModels.PortfolioVideoDataModel
}

func (this *PortfolioDomainModel) ToDomainModel (
	portfolioDataModel *DataModels.PortfolioDataModel,
	descriptionDataModels *[]DataModels.PortfolioDescriptionDataModel,
	imageDataModels *[]DataModels.PortfolioImageDataModel,
	videoDataModels *[]DataModels.PortfolioVideoDataModel,
) {
	this.Uuid = portfolioDataModel.Uuid
	this.Name = portfolioDataModel.Name
	this.Thumbnail = portfolioDataModel.Thumbnail
	this.Year = portfolioDataModel.Year
	this.Description = descriptionDataModels
	this.Image = imageDataModels
	this.Video = videoDataModels
}

func (this *PortfolioDomainModel) ToDataModel() (*DataModels.PortfolioDataModel) {
	return &DataModels.PortfolioDataModel {
		Uuid: this.Uuid,
		Name: this.Name,
		Thumbnail: this.Thumbnail,
		Year: this.Year,
	}
}