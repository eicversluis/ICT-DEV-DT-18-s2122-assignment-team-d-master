package handlers

import (
	"eindopdracht/repositories"
	"eindopdracht/types"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type ModuleData struct {
	Groups []types.Group
	Module types.Module
}

func ModulesHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	Modules := repositories.GetModules()
	err = templ.ExecuteTemplate(w, "modules.gohtml", Modules)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}

func ModuleHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	data := ModuleData{}
	data.Module, err = repositories.GetModule(id, "Grades.Student", "Grades.Student.Group")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err = templ.ExecuteTemplate(w, "404.gohtml", nil)
		if err != nil {
			log.Panic("Error serving template:" + err.Error())
		}
	}
	data.Groups = repositories.GetGroups()
	err = templ.ExecuteTemplate(w, "module.gohtml", data)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}
