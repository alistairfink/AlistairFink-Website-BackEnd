package Commands

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

type PortfolioDescriptionCommand struct {
	DB *pg.DB
}

func (this *PortfolioDescriptionCommand) Get(uuid uuid.UUID) (*DataModels.PortfolioDescriptionDataModel) { 
	if !this.Exists(uuid) {
		return nil
	}

	var models []DataModels.PortfolioDescriptionDataModel
	err := this.DB.Model(&models).Where("id = ?", uuid).Select()
	if err != nil {
		panic(err)
	}

	return &models[0]
}

func (this *PortfolioDescriptionCommand) GetByPortfolioUuid(portfolioUuid uuid.UUID) (*[]DataModels.PortfolioDescriptionDataModel) {
	var models []DataModels.PortfolioDescriptionDataModel
	err := this.DB.Model(&models).Where("portfolio_id = ?", portfolioUuid).Select()
	if err != nil {
		panic(err)
	}

	return &models
}

func (this *PortfolioDescriptionCommand) GetAll() (*[]DataModels.PortfolioDescriptionDataModel) {
	var models []DataModels.PortfolioDescriptionDataModel
	err := this.DB.Model(&models).Select()
	if err != nil {
		panic(err)
	}

	return &models
}

func (this *PortfolioDescriptionCommand) Upsert(model *DataModels.PortfolioDescriptionDataModel) (*DataModels.PortfolioDescriptionDataModel) {
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

func (this *PortfolioDescriptionCommand) Delete(uuid uuid.UUID) bool {
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

func (this *PortfolioDescriptionCommand) Exists(uuid uuid.UUID) bool {
	var models []DataModels.PortfolioDescriptionDataModel
	exists, err := this.DB.Model(&models).Where("id = ?", uuid).Exists()
	if err != nil {
		panic(err)
	}

	return exists
}