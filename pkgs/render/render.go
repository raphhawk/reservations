package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/raphhawk/reservation/pkgs/config"
	"github.com/raphhawk/reservation/pkgs/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders template from html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var cache map[string]*template.Template
	if app.UseCache {
		cache = app.TemplateCache
	} else {
		cache, _ = CreateTemplateCache()
	}
	t, ok := cache[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	td = AddDefaultData(td)
	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to buffer: ", err)
	}
}

// CreateTemplateCache creates template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		globalTemplate, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(globalTemplate) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}
		cache[name] = ts
	}
	return cache, nil
}
