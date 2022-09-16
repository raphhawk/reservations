package main

//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
import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/raphhawk/reservation/pkgs/config"
	"github.com/raphhawk/reservation/pkgs/handlers"
	"github.com/raphhawk/reservation/pkgs/render"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//turn this to true in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	cache, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("error creating template cache: ", err)
	}

	app.TemplateCache = cache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Listening on port %s", port))

	srv := &http.Server{
		Addr:    port,
		Handler: Routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
