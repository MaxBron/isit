package post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/internal/database"
	"project/internal/models"
)

func NewScholarshipHandler(w http.ResponseWriter, r *http.Request) {
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
	scholarship := models.Scholarship{}
	json.Unmarshal(buf.Bytes(), &scholarship)
	database.DB.Db.Create(&scholarship)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"student_id":"%d"}, {"type":"%s"}`, 0, scholarship.Type)))
}
