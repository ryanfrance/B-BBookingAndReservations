package render

import (
	"net/http"
	"testing"

	"github.com/ryanfrance/B-BBookingAndReservations/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	app.Session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("Flash value of 123 not found in")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"

	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("Error writing template to browser")
	}

	err = RenderTemplate(&ww, r, "imposter.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("Imposter template has rendered and should not exist")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/something", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = app.Session.Load(ctx, "X-Session")
	r = r.WithContext(ctx)

	return r, nil
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error()
	}
}
