package main

import (
	"net/http"
	"os"
	"fmt"
	"log"
	"strings"
	"github.com/go-pg/pg"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-chi/chi/middleware"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Utilities"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Data/DatabaseConnection"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/Middleware"
	"github.com/alistairfink/AlistairFink-Website-BackEnd/Domain/Controllers"
)

func main() {
	// Read Config
	var config *Utilities.Config
	if _, err := os.Stat("./Config.json"); err == nil {
		config = Utilities.GetConfig(".", "Config")
	} else {
		config = Utilities.GetConfig("/go/src/github.com/alistairfink/AlistairFink-Website-BackEnd/.", "Config")
	}

	// Open DB
	db := DatabaseConnection.Connect(config)
	defer DatabaseConnection.Close(db)
	
	// Router
	router := Routes(db, config)
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf(" %-10s%-10s\n", method, strings.Replace(route, "/*", "", -1))
		return nil
	}

	println("================================= Routes =================================")
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}
	println("==========================================================================\n")

	log.Fatal(http.ListenAndServe(":" + config.Port, router))
}

func Routes(db *pg.DB, config *Utilities.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use (
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		Middleware.CorsMiddleware,
	)

	// Controllers
	aboutController := Controllers.NewAboutController(db, config)
	educationController := Controllers.NewEducationController(db, config)
	experienceController := Controllers.NewExperienceController(db, config)
	portfolioController := Controllers.NewPortfolioController(db, config)

	// Paths
	router.Route("/api", func(routes chi.Router) {
		routes.Mount("/about", aboutController.Routes())
		routes.Mount("/education", educationController.Routes())
		routes.Mount("/experience", experienceController.Routes())
		routes.Mount("/portfolio", portfolioController.Routes())
	})

	return router
}