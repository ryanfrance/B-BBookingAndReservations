package render

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
	"github.com/ryanfrance/B-BBookingAndReservations/internal/config"
	"github.com/ryanfrance/B-BBookingAndReservations/internal/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig
var pathToTemplates = "./templates"

// NewRenderer sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	app = a
}

func AddDefaultData(templateData *models.TemplateData, r *http.Request) *models.TemplateData {
	templateData.Flash = app.Session.PopString(r.Context(), "flash")
	templateData.Error = app.Session.PopString(r.Context(), "error")
	templateData.Warning = app.Session.PopString(r.Context(), "warning")
	templateData.CSRFToken = nosurf.Token(r)
	return templateData
}

// Template renders templates using html/template
func Template(w http.ResponseWriter, r *http.Request, tmpl string, templateData *models.TemplateData) error {
	var templateCache map[string]*template.Template

	if app.UseCache {
		// get the template cache from AppConfig
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// get requested template from cache
	template, ok := templateCache[tmpl]

	if !ok {
		return errors.New("cannot retrieve template from cache")
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData, r)

	_ = template.Execute(buf, templateData)

	// render the template
	_, err := buf.WriteTo(w)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))

	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
