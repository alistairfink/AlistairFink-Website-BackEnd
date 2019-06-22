package DomainModels

import (
	"time"
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

type EducationDomainModel struct {
	Uuid uuid.UUID
	SchoolName string
	StartDate time.Time
	EndDate time.Time
	Location string
	Certification string
	SchoolLogo string
	Scholarship *[]DataModels.EducationScholarshipDataModel
	Award *[]DataModels.EducationAwardDataModel
	Extracurricular *[]DataModels.EducationExtracurricularDataModel
}

func (this *EducationDomainModel) ToDomainModel (
	educationDataModel *DataModels.EducationDataModel,
	scholarshipDataModels *[]DataModels.EducationScholarshipDataModel,
	awardDataModels *[]DataModels.EducationAwardDataModel,
	extracurricularDataModels *[]DataModels.EducationExtracurricularDataModel,
) {
	this.Uuid = educationDataModel.Uuid
	this.SchoolName = educationDataModel.SchoolName
	this.StartDate = educationDataModel.StartDate
	this.EndDate = educationDataModel.EndDate
	this.Location = educationDataModel.Location
	this.Certification = educationDataModel.Certification
	this.SchoolLogo = educationDataModel.SchoolLogo
	this.Scholarship = scholarshipDataModels
	this.Award = awardDataModels
	this.Extracurricular = extracurricularDataModels
}

func (this *EducationDomainModel) ToDataModel() (*DataModels.EducationDataModel) {
	return &DataModels.EducationDataModel {
		Uuid: this.Uuid,
		SchoolName: this.SchoolName,
		StartDate: this.StartDate,
		EndDate: this.EndDate,
		Location: this.Location,
		Certification: this.Certification,
		SchoolLogo: this.SchoolLogo,
	}
}