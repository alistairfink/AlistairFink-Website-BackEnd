package Controllers

import (
	"github.com/go-pg/pg"
	"github.com/go-chi/chi"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/Managers"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
)

type EducationController struct {
	educationManager *Managers.EducationManager
}

func NewEducationController(
	db *pg.DB,
) (*EducationController) {
	educationCommand := &Commands.EducationCommand {
		DB: db,
	}
	educationScholarshipCommand := &Commands.EducationScholarshipCommand {
		DB: db,
	}
	educationAwardCommand := &Commands.EducationAwardCommand {
		DB: db,
	}
	educationExtracurricularCommand := &Commands.EducationExtracurricularCommand {
		DB: db,
	}

	educationManager := &Managers.EducationManager {
		EducationCommand: educationCommand,
		EducationScholarshipCommand: educationScholarshipCommand,
		EducationAwardCommand: educationAwardCommand,
		EducationExtracurricularCommand: educationExtracurricularCommand,
	}

	return &EducationController {
		educationManager: educationManager, 
	}
}

func (this *EducationController) Routes() (*chi.Mux) {
	router := chi.NewRouter()

	return router
}