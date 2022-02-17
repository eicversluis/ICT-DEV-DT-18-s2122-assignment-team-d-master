package handlers

import (
	"eindopdracht/repositories"
	. "eindopdracht/types"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func AddStudentHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	Data := repositories.GetGroups()
	err = templ.ExecuteTemplate(w, "addstudent.gohtml", Data)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}

func AddStudentPostHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	r.ParseForm()
	s := Student{Name: r.Form.Get("name")}
	id, err := strconv.Atoi(r.Form.Get("group"))
	s.GroupID = uint(id)
	err = repositories.CreateStudent(s)
	if err != nil {
		http.Redirect(w, r, "/studenten/nieuw", http.StatusSeeOther)
		return
	}
	err = templ.ExecuteTemplate(w, "studentadded.gohtml", nil)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}
