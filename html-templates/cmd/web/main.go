package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bobclarke/html-templates/pkg/config"
	"github.com/bobclarke/html-templates/pkg/handlers"
	"github.com/bobclarke/html-templates/pkg/render"
)

const portNum = ":8090" // Set the listening port
var app config.AppConfig
var sessionManager *scs.SessionManager

func main() {

	app.IsProduction = false

	// Set up the session manager
	sessionManager = scs.New()
	sessionManager.Lifetime = time.Hour * 24
	sessionManager.Cookie.Persist = false
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.IsProduction

	// Share the sessionManager config
	app.SessionManager = sessionManager

	// Create the templace cache, make it available to the whole app via the AppConfig struct
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot load template cache", err)
	}
	app.TemplateCache = templateCache

	// Using the Repository pattern I can make the AppConfig struct available to the Handlers
	repo := handlers.NewRepo(&app)
	handlers.SetRepositoryForHandlers(repo)
	render.Templates(&app)

	// Start the webserver
	fmt.Printf("Starting webserver on port %s\n", portNum)
	srv := http.Server{
		Addr:    portNum,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
