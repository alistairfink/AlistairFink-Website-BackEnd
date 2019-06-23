package Managers

import (
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/Sort"
)

type AboutManager struct {
	AboutCommand *Commands.AboutCommand
	AboutDescriptionCommand *Commands.AboutDescriptionCommand
}

func (this *AboutManager) Get() (*DomainModels.AboutDomainModel){
	about := this.AboutCommand.Get()

	aboutDesc := this.AboutDescriptionCommand.GetByAboutUuid(about.Uuid)

	var domainModel DomainModels.AboutDomainModel
	Sort.SortAboutDescriptionBySortOrder(aboutDesc)
	domainModel.ToDomainModel(about, aboutDesc)
	return &domainModel
}

func (this *AboutManager) Update(model *DomainModels.AboutDomainModel) (*DomainModels.AboutDomainModel) {
	if !this.AboutCommand.Exists(model.Uuid) {
		return nil
	}

	for _, desc := range *model.Description {
		if desc.AboutUuid != model.Uuid {
			return nil
		}
	}

	this.AboutCommand.Upsert(model.ToDataModel())
	existing := this.AboutDescriptionCommand.GetByAboutUuid(model.Uuid)

	toKeep := make(map[uuid.UUID]bool)
	for _, desc := range *model.Description {
		dataModel := this.AboutDescriptionCommand.Upsert(&desc)
		toKeep[dataModel.Uuid] = true
	}

	for _, exist := range *existing {
		if !toKeep[exist.Uuid] {
			this.AboutDescriptionCommand.Delete(exist.Uuid)
		}
	}

	return this.Get()
}