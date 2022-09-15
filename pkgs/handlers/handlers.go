package handlers

import (
	"net/http"

	"github.com/raphhawk/reservation/pkgs/config"
	"github.com/raphhawk/reservation/pkgs/models"
	"github.com/raphhawk/reservation/pkgs/render"
)

var Repo *Repository

type Repository struct {
	app *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (re *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	re.app.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (re *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hellow Again"

	remoteIP := re.app.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
