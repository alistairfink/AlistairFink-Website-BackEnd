package Controllers

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/go-pg/pg"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Utilities"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/Managers"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
)

type ExperienceController struct {
	config *Utilities.Config
	experienceManager *Managers.ExperienceManager
}

func NewExperienceController(
	db *pg.DB,
	config *Utilities.Config,
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
		config: config,
		experienceManager: experienceManager,
	}
}

func (this *ExperienceController) Routes() (*chi.Mux) {
	router := chi.NewRouter()
	router.Get("/{experience_uuid}", this.Get)
	router.Get("/", this.GetAll)
	router.Post("/", this.Create)
	router.Put("/", this.Update)
	router.Delete("/{experience_uuid}", this.Delete)
	return router
}

func (this *ExperienceController) Get(w http.ResponseWriter, r *http.Request) {
	uuidUnParsed := chi.URLParam(r, "experience_uuid")
	uuid, err := uuid.Parse(uuidUnParsed)
	if err != nil {
		http.Error(w, "Invalid Experience Uuid", http.StatusBadRequest)
		return
	}

	result := this.experienceManager.Get(uuid)
	if result == nil {
		http.Error(w, "Error Processing Request", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, result)
}

func (this *ExperienceController) GetAll(w http.ResponseWriter, r *http.Request) {
	result := this.experienceManager.GetAll()
	render.JSON(w, r, result)
}

func (this *ExperienceController) Create(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("APIKey")
	if apiKey == this.config.ApiKey {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		var newExperience DomainModels.ExperienceDomainModel
		err = json.Unmarshal(body, &newExperience)
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		result := this.experienceManager.Insert(&newExperience)
		if result == nil {
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
	} else {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
	}
}

func (this *ExperienceController) Update(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("APIKey")
	if apiKey == this.config.ApiKey {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		var newExperience DomainModels.ExperienceDomainModel
		err = json.Unmarshal(body, &newExperience)
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		result := this.experienceManager.Update(&newExperience)
		if result == nil {
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
	} else {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
	}
}

func (this *ExperienceController) Delete(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("APIKey")
	if apiKey == this.config.ApiKey {
		uuidUnParse := chi.URLParam(r, "experience_uuid")
		uuid, err := uuid.Parse(uuidUnParse)
		if err != nil {
			http.Error(w, "Invalid Education Uuid", http.StatusBadRequest)
			return
		}

		if !this.experienceManager.Delete(uuid) {
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

   		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
	}
}