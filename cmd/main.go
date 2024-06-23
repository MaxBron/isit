package main

import (
	"net/http"
	"project/internal/database"
	"project/internal/http/delete"
	"project/internal/http/get"
	"project/internal/http/patch"
	"project/internal/http/post"

	_ "github.com/lib/pq"
)

func main() {
	database.ConnectDB()
	// GET
	http.Handle("GET /student", http.HandlerFunc(get.StudentHandler))
	http.Handle("GET /faculty", http.HandlerFunc(get.FacultyHandler))
	http.Handle("GET /scholarship", http.HandlerFunc(get.ScholarshipHandler))
	http.Handle("GET /studentgroup", http.HandlerFunc(get.StudentGroupHandler))
	http.Handle("GET /dormitory", http.HandlerFunc(get.DormitoryHandler))
	// POST
	http.Handle("POST /signin", http.HandlerFunc(post.SignInHandler))
	http.Handle("POST /newstudent", http.HandlerFunc(post.NewStudentHandler))
	http.Handle("POST /newfaculty", http.HandlerFunc(post.NewFacultyHandler))
	http.Handle("POST /newdormitory", http.HandlerFunc(post.NewDormitoryHandler))
	http.Handle("POST /newscholarship", http.HandlerFunc(post.NewScholarshipHandler))
	http.Handle("POST /newstudentgroup", http.HandlerFunc(post.NewStudentGroupHandler))
	// PATCH
	http.Handle("PATCH /updatestudent", http.HandlerFunc(patch.UpdateStudentHandler))
	http.Handle("PATCH /updatefaculty", http.HandlerFunc(patch.UpdateFacultyHandler))
	http.Handle("PATCH /updategroup", http.HandlerFunc(patch.UpdateGroupHandler))
	http.Handle("PATCH /updatedormitory", http.HandlerFunc(patch.UpdateDormitoryHandler))
	http.Handle("PATCH /updatescholarship", http.HandlerFunc(patch.UpdateScholarshipHandler))
	// DELETE
	http.Handle("DELETE /deletestudent", http.HandlerFunc(delete.DeleteStudentHandler))
	http.Handle("DELETE /deletegroup", http.HandlerFunc(delete.DeleteGroupHandler))
	http.Handle("DELETE /deletedormitory", http.HandlerFunc(delete.DeleteDormitoryHandler))
	http.Handle("DELETE /deletefaculty", http.HandlerFunc(delete.DeleteFacultyHandler))
	http.Handle("DELETE /deletescholarship", http.HandlerFunc(delete.DeleteScholarshipHandler))
	http.ListenAndServe(":3000", nil)
}
