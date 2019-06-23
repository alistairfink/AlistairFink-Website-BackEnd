package Managers

import (
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
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
	Sort.SortPortfolioDescriptionBySortOrder(description)
	images := this.PortfolioImageCommand.GetByPortfolioUuid(uuid)
	Sort.SortPortfolioImageBySortOrder(images)
	videos := this.PortfolioVideoCommand.GetByPortfolioUuid(uuid)
	Sort.SortPortfolioVideoBySortOrder(videos)

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
	featured := this.PortfolioCommand.GetFeatured()
	Sort.SortPortfolioFeaturedBySortOrder(featured)

	// Featured table isn't too big so this might be ok
	// If this is a problem redo this to join portfolio_featured and portfolio
	var portfolioDomainModels []DomainModels.PortfolioDomainModel
	for _, portfolioItem := range *featured {
		portfolioDataModel := this.PortfolioCommand.Get(portfolioItem.PortfolioUuid)
		var domainModel DomainModels.PortfolioDomainModel
		domainModel.ToDomainModel(portfolioDataModel, &[]DataModels.PortfolioDescriptionDataModel{}, &[]DataModels.PortfolioImageDataModel{}, &[]DataModels.PortfolioVideoDataModel{})
		portfolioDomainModels = append(portfolioDomainModels, domainModel)
	}

	return &portfolioDomainModels
}

func (this *PortfolioManager) UpdateFeatured(models *[]DataModels.PortfolioFeaturedDataModel) (*[]DomainModels.PortfolioDomainModel) {
	existingFeats := this.PortfolioCommand.GetFeatured()
	keepFeat := make(map[uuid.UUID]bool)
	for _, feat := range *models {
		keepFeat[feat.PortfolioUuid] = true
		if !this.PortfolioCommand.Exists(feat.PortfolioUuid) {
			return nil
		} 
	}

	for _, feat := range *models {
		this.PortfolioCommand.UpsertFeatured(&feat)
	}

	for _, feat := range *existingFeats {
		if !keepFeat[feat.PortfolioUuid] {
			this.PortfolioCommand.DeleteFeatured(feat.PortfolioUuid)
		}
	}

	return this.GetFeatured()
}

func (this *PortfolioManager) Update(model *DomainModels.PortfolioDomainModel) (*DomainModels.PortfolioDomainModel) {
	if !this.PortfolioCommand.Exists(model.Uuid) {
		return nil
	}

	existingVideo := this.PortfolioVideoCommand.GetByPortfolioUuid(model.Uuid)
	keepVideo := make(map[uuid.UUID]bool)
	for _, video := range *model.Video {
		keepVideo[video.Uuid] = true
		if (video.Uuid != uuid.Nil && !this.PortfolioVideoCommand.Exists(video.Uuid)) || video.PortfolioUuid != model.Uuid {
			return nil
		}
	}

	for _, video := range *existingVideo {
		if !keepVideo[video.Uuid] {
			this.PortfolioVideoCommand.Delete(video.Uuid)
		}
	}

	existingImage := this.PortfolioImageCommand.GetByPortfolioUuid(model.Uuid)
	keepImage := make(map[uuid.UUID]bool)
	for _, image := range *model.Image {
		keepImage[image.Uuid] = true
		if (image.Uuid != uuid.Nil && !this.PortfolioImageCommand.Exists(image.Uuid)) || image.PortfolioUuid != model.Uuid {
			return nil
		}
	}

	for _, image := range *existingImage {
		if !keepImage[image.Uuid] {
			this.PortfolioImageCommand.Delete(image.Uuid)
		}
	}

	existingDesc := this.PortfolioDescriptionCommand.GetByPortfolioUuid(model.Uuid)
	keepDesc := make(map[uuid.UUID]bool)
	for _, desc := range *model.Description {
		keepDesc[desc.Uuid] = true
		if (desc.Uuid != uuid.Nil && !this.PortfolioDescriptionCommand.Exists(desc.Uuid)) || desc.PortfolioUuid != model.Uuid {
			return nil
		}
	}

	for _, desc := range *existingDesc {
		if !keepDesc[desc.Uuid] {
			this.PortfolioDescriptionCommand.Delete(desc.Uuid)
		}
	}

	return this.Insert(model)
}

func (this *PortfolioManager) Insert(model *DomainModels.PortfolioDomainModel) (*DomainModels.PortfolioDomainModel) {
	dataModel := this.PortfolioCommand.Upsert(model.ToDataModel())
	
	for _, video := range *model.Video {
		video.PortfolioUuid = dataModel.Uuid
		this.PortfolioVideoCommand.Upsert(&video)
	}

	for _, image := range *model.Image {
		image.PortfolioUuid = dataModel.Uuid
		this.PortfolioImageCommand.Upsert(&image)
	}

	for _, desc := range *model.Description {
		desc.PortfolioUuid = dataModel.Uuid
		this.PortfolioDescriptionCommand.Upsert(&desc)
	}

	return this.Get(dataModel.Uuid)
}

func (this *PortfolioManager) Delete(uuid uuid.UUID) bool {
	if !this.PortfolioCommand.Exists(uuid) {
		return false
	}
	
	existingFeatured := this.PortfolioCommand.GetFeatured()
	for _, feat := range *existingFeatured {
		if feat.PortfolioUuid == uuid {
			this.PortfolioCommand.DeleteFeatured(uuid)
			break
		}
	}

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

	return this.PortfolioCommand.Delete(uuid)
}