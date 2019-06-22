package Commands

import (
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
)

type PortfolioCommand struct {
	DB *pg.DB
}

func (this *PortfolioCommand) Get(uuid uuid.UUID) (*DataModels.PortfolioDataModel) { 
	if !this.Exists(uuid) {
		return nil
	}

	var models []DataModels.PortfolioDataModel
	err := this.DB.Model(&models).Where("id = ?", uuid).Select()
	if err != nil {
		panic(err)
	}

	return &models[0]
}

func (this *PortfolioCommand) GetAll() (*[]DataModels.PortfolioDataModel) {
	var models []DataModels.PortfolioDataModel
	err := this.DB.Model(&models).Select()
	if err != nil {
		panic(err)
	}

	return &models
}

func (this *PortfolioCommand) GetFeatured() (*[]DataModels.PortfolioDataModel) {
	var models []DataModels.PortfolioFeaturedDataModel
	err := this.DB.Model(&models).Select()
	if err != nil {
		panic(err)
	}

	var portfolioModels []DataModels.PortfolioDataModel
	for _, featuredItem := range models {
		portfolioModels = append(portfolioModels, *this.Get(featuredItem.PortfolioUuid))
	}

	return &portfolioModels
}

func (this *PortfolioCommand) Upsert(model *DataModels.PortfolioDataModel) (*DataModels.PortfolioDataModel) {
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

func (this *PortfolioCommand) Delete(uuid uuid.UUID) bool {
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

func (this *PortfolioCommand) Exists(uuid uuid.UUID) bool {
	var models []DataModels.PortfolioDataModel
	exists, err := this.DB.Model(&models).Where("id = ?", uuid).Exists()
	if err != nil {
		panic(err)
	}

	return exists
}