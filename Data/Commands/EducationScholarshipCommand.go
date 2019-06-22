package Commands

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

type EducationScholarshipCommand struct {
	DB *pg.DB
}

func (this *EducationScholarshipCommand) Get(uuid uuid.UUID) (*DataModels.EducationScholarshipDataModel) { 
	if !this.Exists(uuid) {
		return nil
	}

	var models []DataModels.EducationScholarshipDataModel
	err := this.DB.Model(&models).Where("id = ?", uuid).Select()
	if err != nil {
		panic(err)
	}

	return &models[0]
}

func (this *EducationScholarshipCommand) GetByEducationUuid(educationUuid uuid.UUID) (*[]DataModels.EducationScholarshipDataModel) {
	var models []DataModels.EducationScholarshipDataModel
	err := this.DB.Model(&models).Where("education_id = ?", educationUuid).Select()
	if err != nil {
		panic(err)
	}

	return &models
}

func (this *EducationScholarshipCommand) GetAll() (*[]DataModels.EducationScholarshipDataModel) {
	var models []DataModels.EducationScholarshipDataModel
	err := this.DB.Model(&models).Select()
	if err != nil {
		panic(err)
	}

	return &models
}

func (this *EducationScholarshipCommand) Upsert(model *DataModels.EducationScholarshipDataModel) (*DataModels.EducationScholarshipDataModel) {
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

func (this *EducationScholarshipCommand) Delete(uuid uuid.UUID) bool {
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

func (this *EducationScholarshipCommand) Exists(uuid uuid.UUID) bool {
	var models []DataModels.EducationScholarshipDataModel
	exists, err := this.DB.Model(&models).Where("id = ?", uuid).Exists()
	if err != nil {
		panic(err)
	}

	return exists
}