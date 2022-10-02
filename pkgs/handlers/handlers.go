package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (re *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hellow Again"

	remoteIP := re.app.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation is the about reservations handler
func (re *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Ja is the Japanese Delight page handler
func (re *Repository) Ja(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "ja.page.tmpl", &models.TemplateData{})
}

// Ke is the Kerala Backwaters page handler
func (re *Repository) Ke(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "ke.page.tmpl", &models.TemplateData{})
}

// Availablity is the search-availablity page handler
func (re *Repository) Availablity(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availablity.page.tmpl", &models.TemplateData{})
}

// PostAvailablity is the search-availablity page handler
func (re *Repository) PostAvailablity(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start Date: %s \nEnd Date: %s", start, end)))
}

type jsonresponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailablityJSON is the availablity page JSON handler
func (re *Repository) AvailablityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonresponse{
		OK:      false,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "	")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact is the contact page handler
func (re *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
