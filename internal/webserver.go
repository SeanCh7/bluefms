package internal

import (
	"html/template"
	"net/http"
	"time"
)

// ServerConfig is a struct containing info needed to launch a webserver
type ServerConfig struct {
	Addr    string
	Timeout time.Duration
}

func New(config ServerConfig) http.Server {
	return http.Server{
		Addr:         config.Addr,
		ReadTimeout:  config.Timeout,
		WriteTimeout: config.Timeout,
	}
}

func registerTemplates() {
	indexTempl := template.New("index")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	title := "BlueFMS Home"
	p, _ := loadpa
}

type Plugin interface {
	RegisterRoutes(mux *http.ServeMux)
}
