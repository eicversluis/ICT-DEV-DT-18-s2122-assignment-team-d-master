package handlers

import (
	"eindopdracht/repositories"
	. "eindopdracht/types"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type GradesData struct {
	Module Module
	Group  Group
	Grades map[uint]Grade
}

func EnterGradesHandler(w http.ResponseWriter, r *http.Request) {
	moduleId, err := strconv.Atoi(chi.URLParam(r, "mid"))
	groupId, err := strconv.Atoi(chi.URLParam(r, "gid"))
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	Data := GradesData{}
	Data.Module, err = repositories.GetModule(moduleId)
	Data.Group, err = repositories.GetGroup(groupId)
	_grades, err := repositories.GetModuleGradesForGroup(moduleId, groupId)
	gmap := make(map[uint]Grade)
	for _, grade := range _grades {
		gmap[grade.StudentID] = grade
	}
	Data.Grades = gmap
	err = templ.ExecuteTemplate(w, "entergrades.gohtml", Data)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}

func EnterGradesPostHandler(w http.ResponseWriter, r *http.Request) {
	moduleId, err := strconv.Atoi(chi.URLParam(r, "mid"))
	groupId, err := strconv.Atoi(chi.URLParam(r, "gid"))
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	err = r.ParseForm()
	Group, err := repositories.GetGroup(groupId, "Students")
	grades := []Grade{}
	for _, student := range Group.Students {
		g, err := strconv.ParseFloat(r.Form.Get("grade-"+strconv.Itoa(int(student.ID))), 64)
		if err != nil {
			log.Panic("Error parsing grades: " + err.Error())
		}
		if g != 0 {
			grades = append(grades, Grade{StudentID: student.ID, Value: g, ModuleID: uint(moduleId)})
		}
	}
	repositories.InsertGrades(grades)
	err = templ.ExecuteTemplate(w, "gradesentered.gohtml", nil)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}
