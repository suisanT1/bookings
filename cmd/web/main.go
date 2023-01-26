package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/suisanT1/bookings/pkg/config"
	"github.com/suisanT1/bookings/pkg/handlers"
	"github.com/suisanT1/bookings/pkg/render"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	// change this to true when in production mode
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// create a template
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	// store temp cache in application
	app.TemplateCache = tc

	app.UseCache = false

	// create repository variable
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println("Starting application on port: ", port)
	// importing routes
	serve := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)

}
