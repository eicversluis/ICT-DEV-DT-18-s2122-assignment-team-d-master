package handlers

import (
	"eindopdracht/repositories"
	. "eindopdracht/types"
	"html/template"
	"log"
	"net/http"
)

func AddModuleHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	Data := repositories.GetGroups()
	err = templ.ExecuteTemplate(w, "addmodule.gohtml", Data)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}

func AddModulePostHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	r.ParseForm()
	m := Module{Name: r.Form.Get("name")}
	err = repositories.CreateModule(m)
	if err != nil {
		http.Redirect(w, r, "/modules/nieuw", http.StatusSeeOther)
		return
	}
	err = templ.ExecuteTemplate(w, "moduleadded.gohtml", nil)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}
