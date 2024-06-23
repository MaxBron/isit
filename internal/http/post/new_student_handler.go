package post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/internal/database"
	"project/internal/models"
)

func NewStudentHandler(w http.ResponseWriter, r *http.Request) {
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
	student := models.Student{}
	json.Unmarshal(buf.Bytes(), &student)
	database.DB.Db.Create(&student)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"student_id":"%d"}`, student.ID)))
}
