package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/internal/scholarship"
)

func NewScholarshipHandler(w http.ResponseWriter, r *http.Request) {
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
	scholarship := scholarship.Scholarship{}
	json.Unmarshal(buf.Bytes(), &scholarship)
	w.Header().Set("Content-Type", "applictaion/json")
	studentId, scholarshipType := scholarship.AddScholarship()
	w.Write([]byte(fmt.Sprintf(`{"student_id":"%d"}, {"type":"%s"}`, studentId, scholarshipType)))
}
