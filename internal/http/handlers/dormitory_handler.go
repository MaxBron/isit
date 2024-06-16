package handlers

import (
	"encoding/json"
	"net/http"
	"project/internal/db"
	"project/internal/dormitory"
)

func DormitoryHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var dormitories []dormitory.Dormitory
	db.DB.Db.Find(&dormitories)
	jsonResp, _ := json.Marshal(dormitories)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}
