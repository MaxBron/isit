package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/internal/dormitory"
)

func NewDormitoryHandler(w http.ResponseWriter, r *http.Request) {
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
	dormitory := dormitory.Dormitory{}
	json.Unmarshal(buf.Bytes(), &dormitory)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"room_number":"%d"}`, dormitory.AddDormitory())))
}
