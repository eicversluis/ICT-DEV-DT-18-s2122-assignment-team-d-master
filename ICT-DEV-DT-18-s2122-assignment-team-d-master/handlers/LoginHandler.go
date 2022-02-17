package handlers

import (
	"eindopdracht/repositories"
	"html/template"
	"log"
	"net/http"
	"time"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseGlob("./templates/*")
	if err != nil {
		log.Panic("Error parsing templates: " + err.Error())
	}
	err = templ.ExecuteTemplate(w, "login.gohtml", nil)
	if err != nil {
		log.Panic("Error serving template: " + err.Error())
	}
}

func DoLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := repositories.AuthenticateTeacher(username, password)

	if err == nil {
		http.SetCookie(w, &http.Cookie{Name: "username", Value: user.Username, Secure: true, Expires: time.Now().Add(time.Duration(time.Now().Day()))})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		log.Println(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "username", Value: "", Expires: time.Unix(0, 0), Secure: true})
	//http.Redirect(w, r, "/login", http.StatusSeeOther)
}
