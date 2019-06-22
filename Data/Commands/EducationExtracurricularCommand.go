package Commands

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

type EducationExtracurricularCommand struct {
	DB *pg.DB
}

func (this *EducationExtracurricularCommand) Get(uuid uuid.UUID) (*DataModels.EducationExtracurricularDataModel) { 
	if !this.Exists(uuid) {
		return nil
	}

	var models []DataModels.EducationExtracurricularDataModel
	err := this.DB.Model(&models).Where("id = ?", uuid).Select()
	if err != nil {
		panic(err)
	}

	return &models[0]
}

func (this *EducationExtracurricularCommand) GetAll() (*[]DataModels.EducationExtracurricularDataModel) {
	var models []DataModels.EducationExtracurricularDataModel
	err := this.DB.Model(&models).Select()
	if err != nil {
		panic(err)
	}

	return &models
}

func (this *EducationExtracurricularCommand) Upsert(model *DataModels.EducationExtracurricularDataModel) (*DataModels.EducationExtracurricularDataModel) {
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

func (this *EducationExtracurricularCommand) Delete(uuid uuid.UUID) bool {
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

func (this *EducationExtracurricularCommand) Exists(uuid uuid.UUID) bool {
	var models []DataModels.EducationExtracurricularDataModel
	exists, err := this.DB.Model(&models).Where("id = ?", uuid).Exists()
	if err != nil {
		panic(err)
	}

	return exists
}