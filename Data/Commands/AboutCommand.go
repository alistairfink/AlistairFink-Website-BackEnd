package Commands

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

type AboutCommand struct {
	DB *pg.DB
}

func (this *AboutCommand) Get(uuid uuid.UUID) (*DataModels.AboutDataModel) { 
	if !this.Exists(uuid) {
		return nil
	}

	var models []DataModels.AboutDataModel
	err := this.DB.Model(&models).Where("id = ?", uuid).Select()
	if err != nil {
		panic(err)
	}

	return &models[0]
}

func (this *AboutCommand) GetAll() (*[]DataModels.AboutDataModel) {
	var models []DataModels.AboutDataModel
	err := this.DB.Model(&models).Select()
	if err != nil {
		panic(err)
	}

	return &models
}

func (this *AboutCommand) Upsert(model *DataModels.AboutDataModel) (*DataModels.AboutDataModel) {
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

func (this *AboutCommand) Delete(uuid uuid.UUID) bool {
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

func (this *AboutCommand) Exists(uuid uuid.UUID) bool {
	var models []DataModels.AboutDataModel
	exists, err := this.DB.Model(&models).Where("id = ?", uuid).Exists()
	if err != nil {
		panic(err)
	}

	return exists
}