package handlers

import (
	"net/http"

	"github.com/bobclarke/html-templates/pkg/config"
	"github.com/bobclarke/html-templates/pkg/models"
	"github.com/bobclarke/html-templates/pkg/render"
)

// -----------------------------------------------------
// Repository pattern - Start
// -----------------------------------------------------

// Declare a new type called Repository which contains one field called App of type *config.AppConfig
type Repository struct {
	App *config.AppConfig
}

// NewRepo will be called by main to instantiate a new instance of Repository which will be loaded with an instance of AppConfig (which was also passed in by main)
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// We need to be able to set the receiver for the handlers, the SetRepositoryForHandlers initialises a Repo var for this
var Repo *Repository

func SetRepositoryForHandlers(r *Repository) {
	Repo = r
	// Beacase r points a concrete instance of Repo (instantiated by main) abd then because the handlers are
	// called from main using handlers.Repo.Home and handlers.About.Home.... the m Receiver is asociated with m
}

// -----------------------------------------------------
// Repository pattern - End
// -----------------------------------------------------

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	sm := make(map[string]string)
	sm["foo"] = "bar"
	render.RenderTemplates(w, "home.page.html", &models.TemplateData{
		StringMap: sm,
	})
}

// About is the home page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.html", &models.TemplateData{})
}
