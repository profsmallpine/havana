package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// handler is http struct for passing services to the router.
type handler struct {
	Logger *log.Logger
}

// goHome is used for handling requests to "/".
func (h *handler) goHome(w http.ResponseWriter, r *http.Request) {
	respond(h.Logger, w, r, "./tmpl/index.tmpl", nil)
}

// respond is used to parse a base template.
func respond(logger *log.Logger, w http.ResponseWriter, r *http.Request, layout string, data interface{}) {
	// Parse static files.
	tmpl := template.Must(template.New("base.tmpl").Funcs(templateFuncs).ParseFiles(
		"./tmpl/base.tmpl",
		// "./tmpl/partials/_header.tmpl",
		// "./tmpl/partials/_intro.tmpl",
		"./tmpl/partials/_about.tmpl",
		// "./tmpl/partials/_services.tmpl",
		// "./tmpl/partials/_work.tmpl",
		// "./tmpl/partials/_social.tmpl",
		// "./tmpl/partials/_contact.tmpl",
		// "./tmpl/partials/_map.tmpl",
		// "./tmpl/partials/_footer.tmpl",
		layout,
	))
	err := tmpl.Funcs(template.FuncMap{}).Execute(w, data)

	// Log template compilation failure.
	if err != nil {
		logger.Println("Template execution error: ", err.Error(), layout)
		return
	}
}

var templateFuncs = map[string]interface{}{
	"javascriptTag": javascriptTag,
	"stylesheetTag": stylesheetTag,
	"currentYear": func() int {
		return time.Now().UTC().Year()
	},
}
