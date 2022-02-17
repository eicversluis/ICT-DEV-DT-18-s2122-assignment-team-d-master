package main

import (
	"eindopdracht/handlers"
	"eindopdracht/helpers"
	"eindopdracht/repositories"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load dotenv when in development
	if os.Getenv("APP_MODE") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	if repositories.Connected() {
		log.Println("Starting Database Connection")
	}

	r := chi.NewRouter()
	r.Use(helpers.CheckAuth)

	r.Get("/login", handlers.LoginPage)
	r.Get("/logout", handlers.Logout)
	r.Post("/login", handlers.DoLogin)

	r.Get("/", handlers.HomeHandler)
	r.Get("/home", handlers.HomeHandler)

	r.Get("/studenten", handlers.StudentsHandler)
	r.Get("/studenten/{id}", handlers.StudentHandler)
	r.Get("/studenten/nieuw", handlers.AddStudentHandler)
	r.Post("/studenten/nieuw", handlers.AddStudentPostHandler)

	r.Get("/groepen", handlers.GroupsHandler)
	r.Get("/groepen/{id}", handlers.GroupHandler)
	r.Get("/groepen/nieuw", handlers.AddGroupHandler)
	r.Post("/groepen/nieuw", handlers.AddGroupPostHandler)

	r.Get("/modules", handlers.ModulesHandler)
	r.Get("/modules/{id}", handlers.ModuleHandler)
	r.Get("/modules/nieuw", handlers.AddModuleHandler)
	r.Post("/modules/nieuw", handlers.AddModulePostHandler)

	r.Get("/cijfers/{mid}/{gid}", handlers.EnterGradesHandler)
	r.Post("/cijfers/{mid}/{gid}", handlers.EnterGradesPostHandler)

	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	log.Println("Started server")
	http.ListenAndServe(":8080", r)
}
