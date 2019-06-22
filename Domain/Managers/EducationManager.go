package Managers

import (
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/Sort"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
)

type EducationManager struct {
	EducationCommand *Commands.EducationCommand
	EducationScholarshipCommand *Commands.EducationScholarshipCommand
	EducationAwardCommand *Commands.EducationAwardCommand
	EducationExtracurricularCommand *Commands.EducationExtracurricularCommand
}

func (this *EducationManager) Get(uuid uuid.UUID) (*DomainModels.EducationDomainModel) {
	educationDataModel := this.EducationCommand.Get(uuid)
	if educationDataModel == nil {
		return nil
	}

	scholarships := this.EducationScholarshipCommand.GetByEducationUuid(uuid)
	Sort.SortEducationScholarshipBySortOrder(scholarships)
	awards := this.EducationAwardCommand.GetByEducationUuid(uuid)
	Sort.SortEducationAwardBySortOrder(awards)
	extracurriculars := this.EducationExtracurricularCommand.GetByEducationUuid(uuid)
	Sort.SortEducationExtracurricularBySortOrder(extracurriculars)

	var educationDomainModel DomainModels.EducationDomainModel
	educationDomainModel.ToDomainModel(educationDataModel, scholarships, awards, extracurriculars)
	return &educationDomainModel
}

func (this *EducationManager) GetAll() (*[]DomainModels.EducationDomainModel) {
	var educationDomainModels []DomainModels.EducationDomainModel
	educationDataModels := this.EducationCommand.GetAll()
	for _, educationDataModel := range *educationDataModels {
		educationDomainModel := this.Get(educationDataModel.Uuid)
		educationDomainModels = append(educationDomainModels, *educationDomainModel)
	}

	Sort.SortEducationByStartDate(&educationDomainModels)
	return &educationDomainModels
}

func (this *EducationManager) Update(model *DomainModels.EducationDomainModel) (*DomainModels.EducationDomainModel) {
	if !this.EducationCommand.Exists(model.Uuid) {
		return nil
	}

	for _, scholarship := range *model.Scholarship {
		if !this.EducationScholarshipCommand.Exists(scholarship.Uuid) {
			return nil
		}
	} 

	for _, award := range *model.Award {
		if !this.EducationAwardCommand.Exists(award.Uuid) {
			return nil
		}
	}

	for _, extracarricular := range *model.Extracurricular {
		if !this.EducationExtracurricularCommand.Exists(extracarricular.Uuid) {
			return nil
		}
	}

	return this.Insert(model)
}

func (this *EducationManager) Insert(model *DomainModels.EducationDomainModel) (*DomainModels.EducationDomainModel) {
	dataModel := this.EducationCommand.Upsert(model.ToDataModel())

	for _, scholarship := range *model.Scholarship {
		scholarship.EducationUuid = dataModel.Uuid
		this.EducationScholarshipCommand.Upsert(&scholarship)
	} 

	for _, award := range *model.Award {
		award.EducationUuid = dataModel.Uuid
		this.EducationAwardCommand.Upsert(&award)
	}

	for _, extracarricular := range *model.Extracurricular {
		extracarricular.EducationUuid = dataModel.Uuid
		this.EducationExtracurricularCommand.Upsert(&extracarricular)
	}

	return this.Get(dataModel.Uuid)
}

func (this *EducationManager) Delete(uuid uuid.UUID) bool {
	if !this.EducationCommand.Exists(uuid) {
		return false
	}

	scholarships := this.EducationScholarshipCommand.GetByEducationUuid(uuid)
	awards := this.EducationAwardCommand.GetByEducationUuid(uuid)
	extracurricular := this.EducationExtracurricularCommand.GetByEducationUuid(uuid)

	for _, scholarship := range *scholarships {
		this.EducationScholarshipCommand.Delete(scholarship.Uuid)
	}

	for _, award := range *awards {
		this.EducationAwardCommand.Delete(award.Uuid)
	}

	for _, extracurricularItem := range *extracurricular {
		this.EducationExtracurricularCommand.Delete(extracurricularItem.Uuid)
	} 

	return this.EducationCommand.Delete(uuid)
}