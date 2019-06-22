package Commands

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

type AboutDescriptionCommand struct {
	DB *pg.DB
}

func (this *AboutDescriptionCommand) Get(uuid uuid.UUID) (*DataModels.AboutDescriptionDataModel) { 
	if !this.Exists(uuid) {
		return nil
	}

	var models []DataModels.AboutDescriptionDataModel
	err := this.DB.Model(&models).Where("id = ?", uuid).Select()
	if err != nil {
		panic(err)
	}

	return &models[0]
}

func (this *AboutDescriptionCommand) GetByAboutUuid(aboutUuid uuid.UUID) (*[]DataModels.AboutDescriptionDataModel) {
	var models []DataModels.AboutDescriptionDataModel
	err := this.DB.Model(&models).Where("about_id = ?", aboutUuid).Select()
	if err != nil {
		panic(err)
	}

	return &models
}

func (this *AboutDescriptionCommand) GetAll() (*[]DataModels.AboutDescriptionDataModel) {
	var models []DataModels.AboutDescriptionDataModel
	err := this.DB.Model(&models).Select()
	if err != nil {
		panic(err)
	}

	return &models
}

func (this *AboutDescriptionCommand) Upsert(model *DataModels.AboutDescriptionDataModel) (*DataModels.AboutDescriptionDataModel) {
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

func (this *AboutDescriptionCommand) Delete(uuid uuid.UUID) bool {
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

func (this *AboutDescriptionCommand) Exists(uuid uuid.UUID) bool {
	var models []DataModels.AboutDescriptionDataModel
	exists, err := this.DB.Model(&models).Where("id = ?", uuid).Exists()
	if err != nil {
		panic(err)
	}

	return exists
}