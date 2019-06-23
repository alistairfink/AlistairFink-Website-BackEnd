package Controllers

import (
	"github.com/go-pg/pg"
	"github.com/go-chi/chi"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/Managers"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
)

type ExperienceController struct {
	experienceManager *Managers.ExperienceManager
}

func NewExperienceController(
	db *pg.DB,
) (*ExperienceController) {
	experienceCommand := &Commands.ExperienceCommand {
		DB: db,
	}
	experienceContentCommand := &Commands.ExperienceContentCommand {
		DB: db,
	}

	experienceManager := &Managers.ExperienceManager {
		ExperienceCommand: experienceCommand,
		ExperienceContentCommand: experienceContentCommand,
	}

	return &ExperienceController {
		experienceManager: experienceManager,
	}
}

func (this *ExperienceController) Routes() (*chi.Mux) {
	router := chi.NewRouter()

	return router
}