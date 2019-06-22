package Managers

import (
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/Sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
)

type ExperienceManager struct {
	ExperienceCommand *Commands.ExperienceCommand
	ExperienceContentCommand *Commands.ExperienceContentCommand
}

func (this *ExperienceManager) Get(uuid uuid.UUID) (*DomainModels.ExperienceDomainModel){
	experienceDatModel := this.ExperienceCommand.Get(uuid)
	if experienceDatModel == nil {
		return nil
	}

	experienceContent := this.ExperienceContentCommand.GetByExperienceUuid(uuid)
	Sort.SortExperienceContentBySortOrder(experienceContent)

	var experienceDomainModel DomainModels.ExperienceDomainModel
	experienceDomainModel.ToDomainModel(experienceDatModel, experienceContent)
	return &experienceDomainModel
}

func (this *ExperienceManager) GetAll() (*[]DomainModels.ExperienceDomainModel) {
	experienceDataModels := this.ExperienceCommand.GetAll()
	var experienceDomainModels []DomainModels.ExperienceDomainModel
	for _, dataModel := range *experienceDataModels {
		domainModel := this.Get(dataModel.Uuid)
		experienceDomainModels = append(experienceDomainModels, *domainModel)
	}

	Sort.SortExperienceByStartDate(&experienceDomainModels)
	return &experienceDomainModels
}

func (this *ExperienceManager) Update(model *DomainModels.ExperienceDomainModel) (*DomainModels.ExperienceDomainModel) {
	if !this.ExperienceCommand.Exists(model.Uuid) {
		return nil
	}

	for _, content := range *model.Content {
		if !this.ExperienceContentCommand.Exists(content.Uuid) {
			return nil
		}
	}

	return this.Insert(model)
}

func (this *ExperienceManager) Insert(model *DomainModels.ExperienceDomainModel) (*DomainModels.ExperienceDomainModel) {
	dataModel := this.ExperienceCommand.Upsert(model.ToDataModel())

	for _, content := range *model.Content {
		this.ExperienceContentCommand.Upsert(&content)
	}

	return this.Get(dataModel.Uuid)
}

func (this *ExperienceManager) Delete(uuid uuid.UUID) bool {
	if !this.ExperienceCommand.Exists(uuid) {
		return false
	}

	content := this.ExperienceContentCommand.GetByExperienceUuid(uuid)
	for _, model := range *content {
		this.ExperienceContentCommand.Delete(model.Uuid)
	}

	return this.ExperienceCommand.Delete(uuid)	
}