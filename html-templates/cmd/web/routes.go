package main

import (
	"net/http"

	"github.com/bobclarke/html-templates/pkg/config"
	"github.com/bobclarke/html-templates/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

// Routes register our URI patterns and return a http handler. NOTE, presently we do not use the passed config.AppConfig
func routes(config *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Get("/home", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
