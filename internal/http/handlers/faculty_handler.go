package handlers

import (
	"encoding/json"
	"net/http"
	"project/internal/db"
	"project/internal/faculty"
)

func FacultyHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var faculties []faculty.Faculty
	db.DB.Db.Find(&faculties)
	jsonResp, _ := json.Marshal(faculties)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}
