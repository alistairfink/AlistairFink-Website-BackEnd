package Managers

import (
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
)

type AboutManager struct {
	AboutCommand *Commands.AboutCommand
	AboutDescriptionCommand *Commands.AboutDescriptionCommand
}

// func (this *AboutManager) Get(uuid uuid.UUID) (*DomainModels.AboutDomainModel){

// }

// func (this *AboutManager) Update(model *DomainModels.AboutDomainModel) (*DomainModels.AboutDomainModel) {

// }