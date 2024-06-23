package post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/internal/database"
	"project/internal/models"
)

func NewFacultyHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	if Role != "admin" {
		w.Write([]byte("No access"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	faculty := models.Faculty{}
	json.Unmarshal(buf.Bytes(), &faculty)
	database.DB.Db.Create(&faculty)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"faculty_name":"%s"}`, faculty.Name)))
}
