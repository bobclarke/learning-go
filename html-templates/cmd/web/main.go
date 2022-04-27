package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bobclarke/html-templates/pkg/config"
	"github.com/bobclarke/html-templates/pkg/handlers"
	"github.com/bobclarke/html-templates/pkg/render"
)

const portNum = ":8090"

//main is the main function of the main package
func main() {

	// Create the templace cache and make it available to the whole app
	var app config.AppConfig
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot load template cache", err)
	}
	app.TemplateCache = templateCache

	// Using the Repository pattern I can make AppConfig available to the Handlers
	repo := handlers.NewRepo(&app)
	handlers.SetRepositoryForHandlers(repo)

	render.Templates(&app)

	http.HandleFunc("/home", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting webserver on port %s\n", portNum)
	http.ListenAndServe(portNum, nil)
}
