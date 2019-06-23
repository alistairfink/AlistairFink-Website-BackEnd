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
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DataModels"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/DomainModels"
)

type PortfolioController struct {
	config *Utilities.Config
	portfolioManager *Managers.PortfolioManager
}

func NewPortfolioController(
	db *pg.DB,
	config *Utilities.Config,
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
		config: config,
		portfolioManager: portfolioManager,
	}
}

func (this *PortfolioController) Routes() (*chi.Mux) {
	router := chi.NewRouter()
	router.Get("/{portfolio_uuid}", this.Get)
	router.Get("/", this.GetAll)
	router.Get("/featured", this.GetFeatured)
	router.Post("/", this.Create)
	router.Put("/", this.Update)
	router.Put("/featured", this.UpdateFeatured)
	router.Delete("/{portfolio_uuid}", this.Delete)
	return router
}

func (this *PortfolioController) Get(w http.ResponseWriter, r *http.Request) {
	uuidUnParse := chi.URLParam(r, "portfolio_uuid")
	uuid, err := uuid.Parse(uuidUnParse)
	if err != nil {
		http.Error(w, "Invalid Portfolio Uuid", http.StatusBadRequest)
		return
	}

	result := this.portfolioManager.Get(uuid)
	if result == nil {
		http.Error(w, "Error Processing Request", http.StatusBadRequest)
		return
	}

	render.JSON(w, r, result)
}

func (this *PortfolioController) GetFeatured(w http.ResponseWriter, r *http.Request) {
	result := this.portfolioManager.GetFeatured()
	render.JSON(w, r, result)
}

func (this *PortfolioController) GetAll(w http.ResponseWriter, r *http.Request) {
	result := this.portfolioManager.GetAll()
	render.JSON(w, r, result)
}

func (this *PortfolioController) Create(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("APIKey")
	if apiKey == this.config.ApiKey {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		var newPortfolio DomainModels.PortfolioDomainModel
		err = json.Unmarshal(body, &newPortfolio)
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		result := this.portfolioManager.Insert(&newPortfolio)
		if result == nil {
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
	} else {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
	}
}

func (this *PortfolioController) Update(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("APIKey")
	if apiKey == this.config.ApiKey {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		var newPortfolio DomainModels.PortfolioDomainModel
		err = json.Unmarshal(body, &newPortfolio)
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		result := this.portfolioManager.Update(&newPortfolio)
		if result == nil {
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
	} else {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
	}
}

func (this *PortfolioController) UpdateFeatured(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("APIKey")
	if apiKey == this.config.ApiKey {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		var newFeatured []DataModels.PortfolioFeaturedDataModel
		err = json.Unmarshal(body, &newFeatured)
		if err != nil {
			log.Print(err)
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		result := this.portfolioManager.UpdateFeatured(&newFeatured)
		if result == nil {
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

		render.JSON(w, r, result)
	} else {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
	}
}

func (this *PortfolioController) Delete(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("APIKey")
	if apiKey == this.config.ApiKey {
		uuidUnParse := chi.URLParam(r, "portfolio_uuid")
		uuid, err := uuid.Parse(uuidUnParse)
		if err != nil {
			http.Error(w, "Invalid Portfolio Uuid", http.StatusBadRequest)
			return
		}

		if !this.portfolioManager.Delete(uuid) {
			http.Error(w, "Error Processing Request", http.StatusBadRequest)
			return
		}

   		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
	}
}