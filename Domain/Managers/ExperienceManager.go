package Managers

import (
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
)

type ExperienceManager struct {
	ExperienceCommand *Commands.ExperienceCommand
	ExperienceContentCommand *Commands.ExperienceContentCommand
}

func (this *ExperienceManager) Get(uuid uuid.UUID) (*DomainModels.ExperienceDomainModel){
	// TODO: Don't forget to sort this and getall
	return nil
}

func (this *ExperienceManager) GetAll() (*[]DomainModels.ExperienceDomainModel) {
	return nil
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
	content := this.ExperienceContentCommand.GetByExperienceUuid(uuid)
	for _, model := range *content {
		this.ExperienceContentCommand.Delete(model.Uuid)
	}

	this.ExperienceCommand.Delete(uuid)
	return true	
}