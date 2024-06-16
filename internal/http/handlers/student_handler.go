package handlers

import (
	"encoding/json"
	"net/http"
	"project/internal/db"
	"project/internal/student"
)

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var students []student.Student
	db.DB.Db.Find(&students)
	jsonResp, _ := json.Marshal(students)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}
