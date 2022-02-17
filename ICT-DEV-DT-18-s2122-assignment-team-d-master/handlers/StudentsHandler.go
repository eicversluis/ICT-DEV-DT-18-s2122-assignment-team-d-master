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

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	Students := repositories.GetStudents()
	err = templ.ExecuteTemplate(w, "students.gohtml", Students)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}

type StudentData struct {
	Student  types.Student
	AvgGrade int
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	StudentID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	Student := repositories.GetStudentById(StudentID, "Grades.Module")
	AvgGrade := repositories.GetAverageGrade(StudentID)
	err = templ.ExecuteTemplate(w, "student.gohtml", StudentData{Student: Student, AvgGrade: AvgGrade})
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}
