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

	existingScholarships := this.EducationScholarshipCommand.GetByEducationUuid(model.Uuid)
	keepScholarships := make(map[uuid.UUID]bool)
	for _, scholarship := range *model.Scholarship {
		keepScholarships[scholarship.Uuid] = true
		if (scholarship.Uuid != uuid.Nil && !this.EducationScholarshipCommand.Exists(scholarship.Uuid)) || scholarship.EducationUuid != model.Uuid {
			return nil
		}
	}

	for _, scholarship := range *existingScholarships {
		if !keepScholarships[scholarship.Uuid] {
			this.EducationScholarshipCommand.Delete(scholarship.Uuid)
		}
	}


	existingAwards := this.EducationAwardCommand.GetByEducationUuid(model.Uuid)
	keepAwards := make(map[uuid.UUID]bool)
	for _, award := range *model.Award {
		keepAwards[award.Uuid] = true
		if (award.Uuid != uuid.Nil && !this.EducationAwardCommand.Exists(award.Uuid)) || award.EducationUuid != model.Uuid {
			return nil
		}
	}

	for _, award := range *existingAwards {
		if !keepAwards[award.Uuid] {
			this.EducationAwardCommand.Delete(award.Uuid)
		}
	}

	existingExtracurriculars := this.EducationExtracurricularCommand.GetByEducationUuid(model.Uuid)
	keepExtracurriculars := make(map[uuid.UUID]bool)
	for _, extracarricular := range *model.Extracurricular {
		keepExtracurriculars[extracarricular.Uuid] = true
		if (extracarricular.Uuid != uuid.Nil && !this.EducationExtracurricularCommand.Exists(extracarricular.Uuid)) || extracarricular.EducationUuid != model.Uuid {
			return nil
		}
	}

	for _, extracurricular := range *existingExtracurriculars {
		if !keepExtracurriculars[extracurricular.Uuid] {
			this.EducationExtracurricularCommand.Delete(extracurricular.Uuid)
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