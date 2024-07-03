package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//built to ensure that it imports only from standard library, and doesnt import anything from other files

// Appconfig holds the application configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager    
}
