package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/internal/student"
)

func NewStudentHandler(w http.ResponseWriter, r *http.Request) {

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
	student := student.Student{}
	json.Unmarshal(buf.Bytes(), &student)
	w.Header().Set("Content-Type", "applictaion/json")

	w.Write([]byte(fmt.Sprintf(`{"student_id":"%d"}`, student.AddStudent())))

}
