package post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/internal/database"
	"project/internal/models"
)

func NewStudentGroupHandler(w http.ResponseWriter, r *http.Request) {
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
	studentGroup := models.StudentGroup{}
	json.Unmarshal(buf.Bytes(), &studentGroup)
	database.DB.Db.Create(&studentGroup)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"number":"%d"}`, studentGroup.Number)))
}
