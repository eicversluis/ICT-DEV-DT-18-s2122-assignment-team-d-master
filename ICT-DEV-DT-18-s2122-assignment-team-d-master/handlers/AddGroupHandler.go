package handlers

import (
	"eindopdracht/repositories"
	. "eindopdracht/types"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func AddGroupHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	Data := repositories.GetTeachers()
	err = templ.ExecuteTemplate(w, "addgroup.gohtml", Data)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}

func AddGroupPostHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	r.ParseForm()
	g := Group{Name: r.Form.Get("name")}
	id, err := strconv.Atoi(r.Form.Get("teacher"))
	g.TeacherID = uint(id)
	err = repositories.CreateGroup(g)
	if err != nil {
		http.Redirect(w, r, "/groepen/nieuw", http.StatusSeeOther)
		return
	}
	err = templ.ExecuteTemplate(w, "groupadded.gohtml", nil)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}
