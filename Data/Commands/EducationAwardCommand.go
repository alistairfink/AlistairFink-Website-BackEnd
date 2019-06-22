package Commands

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

type EducationAwardCommand struct {
	DB *pg.DB
}

func (this *EducationAwardCommand) Get(uuid uuid.UUID) (*DataModels.EducationAwardDataModel) { 
	if !this.Exists(uuid) {
		return nil
	}

	var models []DataModels.EducationAwardDataModel
	err := this.DB.Model(&models).Where("id = ?", uuid).Select()
	if err != nil {
		panic(err)
	}

	return &models[0]
}

func (this *EducationAwardCommand) GetByEducationUuid(educationUuid uuid.UUID) (*[]DataModels.EducationAwardDataModel) {
	var models []DataModels.EducationAwardDataModel
	err := this.DB.Model(&models).Where("education_id = ?", educationUuid).Select()
	if err != nil {
		panic(err)
	}

	return &models
}

func (this *EducationAwardCommand) GetAll() (*[]DataModels.EducationAwardDataModel) {
	var models []DataModels.EducationAwardDataModel
	err := this.DB.Model(&models).Select()
	if err != nil {
		panic(err)
	}

	return &models
}

func (this *EducationAwardCommand) Upsert(model *DataModels.EducationAwardDataModel) (*DataModels.EducationAwardDataModel) {
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

func (this *EducationAwardCommand) Delete(uuid uuid.UUID) bool {
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

func (this *EducationAwardCommand) Exists(uuid uuid.UUID) bool {
	var models []DataModels.EducationAwardDataModel
	exists, err := this.DB.Model(&models).Where("id = ?", uuid).Exists()
	if err != nil {
		panic(err)
	}

	return exists
}