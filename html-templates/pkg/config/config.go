package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache       bool
	TemplateCache  map[string]*template.Template
	Logger         log.Logger
	IsProduction   bool
	SessionManager *scs.SessionManager
}
