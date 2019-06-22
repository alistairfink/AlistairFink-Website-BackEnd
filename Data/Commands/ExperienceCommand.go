package Commands

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

type ExperienceCommand struct {
	DB *pg.DB
}

func (this *ExperienceCommand) Get(uuid uuid.UUID) (*DataModels.ExperienceDataModel) { 
	if !this.Exists(uuid) {
		return nil
	}

	var models []DataModels.ExperienceDataModel
	err := this.DB.Model(&models).Where("id = ?", uuid).Select()
	if err != nil {
		panic(err)
	}

	return &models[0]
}

func (this *ExperienceCommand) GetAll() (*[]DataModels.ExperienceDataModel) {
	var models []DataModels.ExperienceDataModel
	err := this.DB.Model(&models).Select()
	if err != nil {
		panic(err)
	}

	return &models
}

func (this *ExperienceCommand) Upsert(model *DataModels.ExperienceDataModel) (*DataModels.ExperienceDataModel) {
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

func (this *ExperienceCommand) Delete(uuid uuid.UUID) bool {
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

func (this *ExperienceCommand) Exists(uuid uuid.UUID) bool {
	var models []DataModels.ExperienceDataModel
	exists, err := this.DB.Model(&models).Where("id = ?", uuid).Exists()
	if err != nil {
		panic(err)
	}

	return exists
}