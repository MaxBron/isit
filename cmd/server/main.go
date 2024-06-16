package main

import (
	"net/http"
	"project/internal/db"
	"project/internal/http/handlers"

	"github.com/go-chi/chi/v5"

	_ "github.com/lib/pq"
)

func main() {
	db.ConnectDB()
	r := chi.NewRouter()
	r.Get("/student", handlers.StudentHandler)
	r.Get("/faculty", handlers.FacultyHandler)
	r.Get("/scholarship", handlers.ScholarshipHandler)
	r.Get("/studentgroup", handlers.StudentGroupHandler)
	r.Get("/dormitory", handlers.DormitoryHandler)
	r.Post("/signin", handlers.SignInHandler)
	r.Post("/newstudent", handlers.NewStudentHandler)
	r.Post("/newfaculty", handlers.NewFacultyHandler)
	r.Post("/newdormitory", handlers.NewDormitoryHandler)
	r.Post("/newscholarship", handlers.NewScholarshipHandler)
	r.Post("/newstudentgroup", handlers.NewStudentGroupHandler)
	http.ListenAndServe(":3000", r)
}
