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

var app *config.AppConfig

// Templates is invoked by the main function and a instance of the AppConfig struct is passed containing populated templates
func Templates(a *config.AppConfig) {
	app = a
}

// RenderTemplates is the HTML template parser
func RenderTemplates(w http.ResponseWriter, tmpl string, data *models.TemplateData) {

	templateCache := app.TemplateCache // the app variable is our AppConfig struct whihc was passed in from main

	template, ok := templateCache[tmpl] // Look in our template cache for our template
	if !ok {
		log.Fatalf("Problem indexing %s in template cache\n", tmpl)
	}

	buf := new(bytes.Buffer) // Create a new buffer

	template.Execute(buf, data) // Execute our template and write the output to our new buffer

	_, err := buf.WriteTo(w) // Write our buffer to our http.ResponseWriter
	if err != nil {
		fmt.Println("Problem writing buffer", err)
	}
}

// CreateTemplateCache is called by main
func CreateTemplateCache() (map[string]*template.Template, error) {

	// Set up a cache to hold pasred templates
	myCache := map[string]*template.Template{}

	// Get all the template files with *.page.html
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {

		// Get the filename from the full path
		name := filepath.Base(page)

		//ts, err = template.New(name).Funcs(functions).ParseFiles(page)
		// Create a tenplate set (This is actually just a Template that we're going to appent multiple templates to - I think)
		ts := template.New(name)

		// Load a page into our template set
		ts, err := ts.ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Check if there are any layout files, if there are read these into out template set
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		} else if len(matches) > 0 {
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
