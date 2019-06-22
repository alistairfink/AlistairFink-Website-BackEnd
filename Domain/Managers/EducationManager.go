package Managers

import (
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
)

type EducationManager struct {
	EducationCommand *Commands.EducationCommand
	EducationScholarshipCommand *Commands.EducationScholarshipCommand
	EducationAwardCommand *Commands.EducationAwardCommand
	EducationExtracurricularCommand *Commands.EducationExtracurricularCommand
}

// func (this *EducationManager) Get(uuid uuid.UUID) (*DomainModels.EducationDomainModel){

// }

// func (this *EducationManager) GetAll() (*[]DomainModels.EducationDomainModel) {
	
// }

// func (this *EducationManager) Update(model *DomainModels.EducationDomainModel) (*DomainModels.EducationDomainModel) {

// }

// func (this *EducationManager) Insert(model *DomainModels.EducationDomainModel) (*DomainModels.EducationDomainModel) {

// }

// func (this *EducationManager) Delete(uuid uuid.UUID) bool {

// }