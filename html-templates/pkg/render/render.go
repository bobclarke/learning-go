package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bobclarke/html-templates/pkg/config"
	"github.com/bobclarke/html-templates/pkg/models"
)

// declare a var called called app of type pointer to config.AppConfig
var app *config.AppConfig

// Templates is invoked by the main function and an instance of the AppConfig struct is passed
// which contains populated templates
func Templates(a *config.AppConfig) {
	app = a
}

// RenderTemplates is the HTML template parser
func RenderTemplates(w http.ResponseWriter, tmpl string, data *models.TemplateData) {

	templateCache := app.TemplateCache // the app variable is our AppConfig struct which was passed in from main

	template, ok := templateCache[tmpl] // Look in our template cache for our template
	if !ok {
		log.Fatalf("Problem indexing %s in template cache\n", tmpl)
	}

	buf := new(bytes.Buffer)    // Create a new buffer
	template.Execute(buf, data) // Execute our template and write the output to our new buffer

	_, err := buf.WriteTo(w) // Write our buffer to our http.ResponseWriter
	if err != nil {
		fmt.Println("Problem writing buffer", err)
	}
}

// CreateTemplateCache is called inkoked at main.go. This function creates a Map which it uses for the cache.
// It iterrates through ./templates/*.page.html
func CreateTemplateCache() (map[string]*template.Template, error) {

	// Set up a cache to hold pasred templates
	myCache := map[string]*template.Template{}

	// Get a list of pages
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// Iterrate through each page
	// Create a template
	// Load the page into the template
	// Load the layouts into the template
	// Add the template to the template cache
	for _, page := range pages {

		name := filepath.Base(page) // Get the page filename

		ts := template.New(name) // Create a template

		ts, err := ts.ParseFiles(page) // Load the page into the template
		if err != nil {
			return myCache, err
		}

		// Check if there are any layout files, if there are read these into out template set
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		// Load the template set for our page into our map
		myCache[name] = ts

	}
	return myCache, nil
}
