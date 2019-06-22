package Managers

import (
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
)

type ExperienceManager struct {
	ExperienceCommand *Commands.ExperienceCommand
	ExperienceContentCommand *Commands.ExperienceContentCommand
}

// func (this *ExperienceManager) Get(uuid uuid.UUID) (*DomainModels.ExperienceDomainModel){

// }

// func (this *ExperienceManager) GetAll() (*[]DomainModels.ExperienceDomainModel) {
	
// }

// func (this *ExperienceManager) Update(model *DomainModels.ExperienceDomainModel) (*DomainModels.ExperienceDomainModel) {

// }

// func (this *ExperienceManager) Insert(model *DomainModels.ExperienceDomainModel) (*DomainModels.ExperienceDomainModel) {

// }

// func (this *ExperienceManager) Delete(uuid uuid.UUID) bool {
	
// }