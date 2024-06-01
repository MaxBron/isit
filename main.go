package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	_ "github.com/lib/pq"
)

var isAuth bool
var role string

func signInHandler(w http.ResponseWriter, r *http.Request) {

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	user := User{}
	json.Unmarshal(buf.Bytes(), &user)
	if DB.Db.First(&user, "login = ? AND password = ?", user.Login, user.Password).Error != nil {
		w.Write([]byte("wrong login or password"))
	} else {
		isAuth = true
		role = user.role
		w.Write([]byte("OK"))
	}

}

func newStudentHandler(w http.ResponseWriter, r *http.Request) {

	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	if role != "admin" {
		w.Write([]byte("No access"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	student := Student{}
	json.Unmarshal(buf.Bytes(), &student)
	w.Header().Set("Content-Type", "applictaion/json")

	w.Write([]byte(fmt.Sprintf(`{"student_id":"%d"}`, student.AddStudent())))

}

func newFacultyHandler(w http.ResponseWriter, r *http.Request) {

	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	if role != "admin" {
		w.Write([]byte("No access"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	faculty := Faculty{}
	json.Unmarshal(buf.Bytes(), &faculty)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"faculty_name":"%s"}`, faculty.AddFaculty())))
}

func newDormitoryHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	if role != "admin" {
		w.Write([]byte("No access"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	dormitory := Dormitory{}
	json.Unmarshal(buf.Bytes(), &dormitory)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"room_number":"%d"}`, dormitory.AddDormitory())))
}

func newScholarshipHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	if role != "admin" {
		w.Write([]byte("No access"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	scholarship := Scholarship{}
	json.Unmarshal(buf.Bytes(), &scholarship)
	w.Header().Set("Content-Type", "applictaion/json")
	studentId, scholarshipType := scholarship.AddScholarship()
	w.Write([]byte(fmt.Sprintf(`{"student_id":"%d"}, {"type":"%s"}`, studentId, scholarshipType)))
}

func newStudentGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	if role != "admin" {
		w.Write([]byte("No access"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	studentGroup := StudentGroup{}
	json.Unmarshal(buf.Bytes(), &studentGroup)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"number":"%d"}`, studentGroup.AddStudentGroup())))
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var students []Student
	DB.Db.Find(&students)
	jsonResp, _ := json.Marshal(students)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}

func facultyHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var faculties []Faculty
	DB.Db.Find(&faculties)
	jsonResp, _ := json.Marshal(faculties)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}

func scholarshipHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var scholarships []Scholarship
	DB.Db.Find(&scholarships)
	jsonResp, _ := json.Marshal(scholarships)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}

func studentGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var studentGroups []StudentGroup
	DB.Db.Find(&studentGroups)
	jsonResp, _ := json.Marshal(studentGroups)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}

func dormitoryHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var dormitories []Dormitory
	DB.Db.Find(&dormitories)
	jsonResp, _ := json.Marshal(dormitories)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}

func main() {

	ConnectDB()
	r := chi.NewRouter()
	r.Get("/student", studentHandler)
	r.Get("/faculty", facultyHandler)
	r.Get("/scholarship", scholarshipHandler)
	r.Get("/studentgroup", studentGroupHandler)
	r.Get("/dormitory", dormitoryHandler)
	r.Post("/signin", signInHandler)
	r.Post("/newstudent", newStudentHandler)
	r.Post("/newfaculty", newFacultyHandler)
	r.Post("/newdormitory", newDormitoryHandler)
	r.Post("/newscholarship", newScholarshipHandler)
	r.Post("/newstudentgroup", newStudentGroupHandler)

	http.ListenAndServe(":3000", r)

}
