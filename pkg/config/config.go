// to prevent import circle

package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig hold the application config
// all the member inside the AppConfig are available to globe
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
