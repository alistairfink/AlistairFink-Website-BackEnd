package Controllers

import (
	"github.com/go-pg/pg"
	"github.com/go-chi/chi"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/Managers"
)

type AboutController struct {
	aboutManager *Managers.AboutManager
}

func NewAboutController(
	db *pg.DB,
) (*AboutController) {
	aboutCommand := &Commands.AboutCommand {
		DB: db,
	}
	aboutDescriptionCommand := &Commands.AboutDescriptionCommand {
		DB: db,
	}

	aboutManager := &Managers.AboutManager {
		AboutCommand: aboutCommand,
		AboutDescriptionCommand: aboutDescriptionCommand,
	}

	return &AboutController {
		aboutManager: aboutManager,
	}
}

func (this *AboutController) Routes() (*chi.Mux) {
	router := chi.NewRouter()

	return router
}