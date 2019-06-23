package Controllers

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/google/uuid"
	"github.com/go-chi/render"
	"github.com/go-pg/pg"
	"github.com/go-chi/chi"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Utilities"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/Managers"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
)

type EducationController struct {
	config *Utilities.Config
	educationManager *Managers.EducationManager
}

func NewEducationController(
	db *pg.DB,
	config *Utilities.Config,
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
		config: config,
		educationManager: educationManager, 
	}
}

func (this *EducationController) Routes() (*chi.Mux) {
	router := chi.NewRouter()
	router.Get("/{education_uuid}", this.Get)
	router.Get("/", this.GetAll)
	router.Post("/", this.Create)
	router.Put("/", this.Update)
	router.Delete("/{education_uuid}", this.Delete)
	return router
}

func (this *EducationController) Get(w http.ResponseWriter, r *http.Request) {
	uuidUnParse := chi.URLParam(r, "education_uuid")
	uuid, err := uuid.Parse(uuidUnParse)
	if err != nil {
		http.Error(w, "Invalid Education Uuid", http.StatusBadRequest)
		return
	}

	result := this.educationManager.Get(uuid)
	if result == nil {
		http.Error(w, "Error Processing Request", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, result)
}

func (this *EducationController) GetAll(w http.ResponseWriter, r *http.Request) {
	result := this.educationManager.GetAll()
	render.JSON(w, r, result)
}

func (this *EducationController) Create(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("APIKey")
	if apiKey == this.config.ApiKey {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		var newEducation DomainModels.EducationDomainModel
		err = json.Unmarshal(body, &newEducation)
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		result := this.educationManager.Insert(&newEducation)
		if result == nil {
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
	} else {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
	}
}

func (this *EducationController) Update(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("APIKey")
	if apiKey == this.config.ApiKey {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		var newEducation DomainModels.EducationDomainModel
		err = json.Unmarshal(body, &newEducation)
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		result := this.educationManager.Update(&newEducation)
		if result == nil {
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
	} else {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
	}
}

func (this *EducationController) Delete(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("APIKey")
	if apiKey == this.config.ApiKey {
		uuidUnParse := chi.URLParam(r, "education_uuid")
		uuid, err := uuid.Parse(uuidUnParse)
		if err != nil {
			http.Error(w, "Invalid Education Uuid", http.StatusBadRequest)
			return
		}

		if !this.educationManager.Delete(uuid) {
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

   		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
	}
}