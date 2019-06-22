package Managers

import (
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
)

type AboutManager struct {
	AboutCommand *Commands.AboutCommand
	AboutDescriptionCommand *Commands.AboutDescriptionCommand
}

func (this *AboutManager) Get(uuid uuid.UUID) (*DomainModels.AboutDomainModel){
	about := this.AboutCommand.Get(uuid)
	if about == nil {
		return nil
	}

	aboutDesc := this.AboutDescriptionCommand.GetByAboutUuid(uuid)

	var domainModel DomainModels.AboutDomainModel
	domainModel.ToDomainModel(about, aboutDesc)
	return &domainModel
}

func (this *AboutManager) Update(model *DomainModels.AboutDomainModel) (*DomainModels.AboutDomainModel) {
	if !this.AboutCommand.Exists(model.Uuid) {
		return nil
	}

	for _, desc := range *model.Description {
		if !this.AboutDescriptionCommand.Exists(desc.Uuid) {
			return nil
		}
	}

	this.AboutCommand.Upsert(model.ToDataModel())
	for _, desc := range *model.Description {
		this.AboutDescriptionCommand.Upsert(&desc)
	}

	return this.Get(model.Uuid)
}