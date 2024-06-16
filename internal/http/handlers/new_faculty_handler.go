package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/internal/faculty"
)

func NewFacultyHandler(w http.ResponseWriter, r *http.Request) {
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
	faculty := faculty.Faculty{}
	json.Unmarshal(buf.Bytes(), &faculty)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"faculty_name":"%s"}`, faculty.AddFaculty())))
}
