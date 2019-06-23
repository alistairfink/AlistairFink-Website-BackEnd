package Controllers

import (
	"github.com/go-pg/pg"
	"github.com/go-chi/chi"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/Managers"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
)

type PortfolioController struct {
	portfolioManager *Managers.PortfolioManager
}

func NewPortfolioController(
	db *pg.DB,
) (*PortfolioController) {
	portfolioCommand := &Commands.PortfolioCommand {
		DB: db,
	}
	portfolioDescriptionCommand := &Commands.PortfolioDescriptionCommand {
		DB: db,
	}
	portfolioImageCommand := &Commands.PortfolioImageCommand {
		DB: db,
	}
	portfolioVideoCommand := &Commands.PortfolioVideoCommand {
		DB: db,
	}
	
	portfolioManager := &Managers.PortfolioManager {
		PortfolioCommand: portfolioCommand,
		PortfolioDescriptionCommand: portfolioDescriptionCommand,
		PortfolioImageCommand: portfolioImageCommand,
		PortfolioVideoCommand: portfolioVideoCommand,
	}

	return &PortfolioController {
		portfolioManager: portfolioManager,
	}
}
func (this *PortfolioController) Routes() (*chi.Mux) {
	router := chi.NewRouter()

	return router
}