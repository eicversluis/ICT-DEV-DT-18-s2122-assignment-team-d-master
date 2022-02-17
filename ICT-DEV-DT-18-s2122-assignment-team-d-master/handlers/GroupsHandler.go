package handlers

import (
	"eindopdracht/repositories"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func GroupsHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	Groups := repositories.GetGroups()
	err = templ.ExecuteTemplate(w, "groups.gohtml", Groups)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}

func GroupHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	Group, err := repositories.GetGroup(id, "Student")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err = templ.ExecuteTemplate(w, "404.gohtml", nil)
		if err != nil {
			log.Panic("Error serving template:" + err.Error())
		}
	} else {
		err = templ.ExecuteTemplate(w, "group.gohtml", Group)
		if err != nil {
			log.Panic("Error serving template: " + err.Error())
		}
	}
}
