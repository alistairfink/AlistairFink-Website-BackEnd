package Commands

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

type EducationCommand struct {
	DB *pg.DB
}

func (this *EducationCommand) Get(uuid uuid.UUID) (*DataModels.EducationDataModel) { 
	if !this.Exists(uuid) {
		return nil
	}

	var models []DataModels.EducationDataModel
	err := this.DB.Model(&models).Where("id = ?", uuid).Select()
	if err != nil {
		panic(err)
	}

	return &models[0]
}

func (this *EducationCommand) GetAll() (*[]DataModels.EducationDataModel) {
	var models []DataModels.EducationDataModel
	err := this.DB.Model(&models).Select()
	if err != nil {
		panic(err)
	}

	return &models
}

func (this *EducationCommand) Upsert(model *DataModels.EducationDataModel) (*DataModels.EducationDataModel) {
	if this.Exists(model.Uuid) {
		_, err := this.DB.Model(model).Where("id = ?", model.Uuid).Update(model)
		if err != nil {
			panic(err)
		}
	} else {
		err := this.DB.Insert(model)
		if err != nil {
			panic(err)
		}
	}

	return this.Get(model.Uuid)
}

func (this *EducationCommand) Delete(uuid uuid.UUID) bool {
	model := this.Get(uuid)
	if model == nil {
		return false
	}

	err := this.DB.Delete(model)
	if err != nil {
		panic(err)
	}

	return true
}

func (this *EducationCommand) Exists(uuid uuid.UUID) bool {
	var models []DataModels.EducationDataModel
	exists, err := this.DB.Model(&models).Where("id = ?", uuid).Exists()
	if err != nil {
		panic(err)
	}

	return exists
}