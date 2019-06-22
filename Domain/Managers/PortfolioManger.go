package Managers

import (
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/Sort"
)

type PortfolioManager struct {
	PortfolioCommand *Commands.PortfolioCommand
	PortfolioDescriptionCommand *Commands.PortfolioDescriptionCommand
	PortfolioImageCommand *Commands.PortfolioImageCommand
	PortfolioVideoCommand *Commands.PortfolioVideoCommand
}

func (this *PortfolioManager) Get(uuid uuid.UUID) (*DomainModels.PortfolioDomainModel){
	portfolioDataModel := this.PortfolioCommand.Get(uuid)
	if portfolioDataModel == nil {
		return nil
	}

	description := this.PortfolioDescriptionCommand.GetByPortfolioUuid(uuid)
	images := this.PortfolioImageCommand.GetByPortfolioUuid(uuid)
	videos := this.PortfolioVideoCommand.GetByPortfolioUuid(uuid)

	var domainModel DomainModels.PortfolioDomainModel
	domainModel.ToDomainModel(portfolioDataModel, description, images, videos)
	return &domainModel
}

func (this *PortfolioManager) GetAll() (*[]DomainModels.PortfolioDomainModel) {
	portfolioItems := this.PortfolioCommand.GetAll()

	domainModels := []DomainModels.PortfolioDomainModel{}
	for _, item := range *portfolioItems {
		var model DomainModels.PortfolioDomainModel
		model.ToDomainModel(&item, nil, nil, nil)
		domainModels = append(domainModels, model)
	}

	Sort.SortPortfolioByYear(&domainModels)
	return &domainModels
}

func (this *PortfolioManager) GetFeatured() (*[]DomainModels.PortfolioDomainModel) {
	portfolioItems := this.PortfolioCommand.GetFeatured()

	domainModels := []DomainModels.PortfolioDomainModel{}
	for _, dataModel := range *portfolioItems {
		var model DomainModels.PortfolioDomainModel
		model.ToDomainModel(&dataModel, nil, nil, nil)
		domainModels = append(domainModels, model)
	}

	return &domainModels
}

func (this *PortfolioManager) Update(model *DomainModels.PortfolioDomainModel) (*DomainModels.PortfolioDomainModel) {
	if !this.PortfolioCommand.Exists(model.Uuid) {
		return nil
	}

	for _, video := range *model.Video {
		if !this.PortfolioVideoCommand.Exists(video.Uuid) {
			return nil
		}
	}

	for _, image := range *model.Image {
		if !this.PortfolioImageCommand.Exists(image.Uuid) {
			return nil
		}
	}

	for _, desc := range *model.Description {
		if !this.PortfolioDescriptionCommand.Exists(desc.Uuid) {
			return nil
		}
	}

	return this.Insert(model)
}

func (this *PortfolioManager) Insert(model *DomainModels.PortfolioDomainModel) (*DomainModels.PortfolioDomainModel) {
	this.PortfolioCommand.Upsert(model.ToDataModel())
	
	for _, video := range *model.Video {
		this.PortfolioVideoCommand.Upsert(&video)
	}

	for _, image := range *model.Image {
		this.PortfolioImageCommand.Upsert(&image)
	}

	for _, desc := range *model.Description {
		this.PortfolioDescriptionCommand.Upsert(&desc)
	}

	return this.Get(model.Uuid)
}

func (this *PortfolioManager) Delete(uuid uuid.UUID) bool {
	portfolioDescriptions := this.PortfolioDescriptionCommand.GetByPortfolioUuid(uuid)
	portfolioImages := this.PortfolioImageCommand.GetByPortfolioUuid(uuid)
	portfolioVideos := this.PortfolioVideoCommand.GetByPortfolioUuid(uuid)

	for _, video := range *portfolioVideos {
		this.PortfolioVideoCommand.Delete(video.Uuid)
	}

	for _, image := range *portfolioImages {
		this.PortfolioImageCommand.Delete(image.Uuid)
	}

	for _, desc := range *portfolioDescriptions {
		this.PortfolioDescriptionCommand.Delete(desc.Uuid)
	}

	this.PortfolioCommand.Delete(uuid)

	return true
}