package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/suisanT1/bookings/pkg/config"
	"github.com/suisanT1/bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates set the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate that render html pages
func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {
	var tc map[string]*template.Template // create a variable to hold template cache

	// check if use cache is true or false for the development
	if app.UseCache {
		// gt the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[html]
	if !ok {
		log.Fatal("Could not get the template from template cache")
	}

	buf := new(bytes.Buffer) // just to see where the error is coming from

	td = AddDefaultData(td) // add default data to every page

	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

// TODO adv. way to add cache templates (no longer need to track what files are in the templates folders; it will automatics added to the templates)
// create template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)

	// create a cache variable that will stores
	myCache := map[string]*template.Template{}

	// get all of the named *.page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {

		// get the file (about.page.html) from page and leave the name itself
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// find if there is layout.html
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		// add the files together
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		// add the page name
		myCache[name] = ts
	}

	return myCache, nil
}

// TODO simple way to add cache templates
// RenderTemplate that render html pages
// func RenderTemplate(w http.ResponseWriter, html string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+html, "./templates/base.layout.html")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error parsing template:", err)
// 		return
// 	}
// }
// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we already have the template in our cache
// 	_, inMap := tc[t]
// 	if !inMap {
// 		// need to create the template
// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	} else {
// 		// we have the template in the cache
// 		fmt.Println("using cached template")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.html",
// 	}
// 	// parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	// add template to cache (map)
// 	tc[t] = tmpl
// 	return nil
// }
