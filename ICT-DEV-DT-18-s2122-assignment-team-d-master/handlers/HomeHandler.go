package handlers

import (
	"html/template"
	"log"
	http "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	err = templ.ExecuteTemplate(w, "home.gohtml", nil)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}
