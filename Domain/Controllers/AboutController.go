package Controllers

import (
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/go-pg/pg"
	"github.com/go-chi/render"
	"github.com/go-chi/chi"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Utilities"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/Commands"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/Managers"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
)

type AboutController struct {
	config *Utilities.Config
	aboutManager *Managers.AboutManager
}

func NewAboutController(
	db *pg.DB,
	config *Utilities.Config,
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
		config: config,
		aboutManager: aboutManager,
	}
}

func (this *AboutController) Routes() (*chi.Mux) {
	router := chi.NewRouter()
	router.Get("/", this.Get)
	router.Put("/", this.Update)
	return router
}

func (this *AboutController) Get(w http.ResponseWriter, r *http.Request) {
	result := this.aboutManager.Get()
	render.JSON(w, r, result)
}

func (this *AboutController) Update(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("APIKey")
	if apiKey == this.config.ApiKey {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		var aboutDomainModel DomainModels.AboutDomainModel
		err = json.Unmarshal(body, &aboutDomainModel)
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Model", http.StatusBadRequest)
			return
		}

		result := this.aboutManager.Update(&aboutDomainModel)
		if result == nil {
			http.Error(w, "Invalid UUID", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
	} else {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
	}
}