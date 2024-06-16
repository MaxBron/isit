package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/internal/studentgroup"
)

func NewStudentGroupHandler(w http.ResponseWriter, r *http.Request) {
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
	studentGroup := studentgroup.StudentGroup{}
	json.Unmarshal(buf.Bytes(), &studentGroup)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"number":"%d"}`, studentGroup.AddStudentGroup())))
}
