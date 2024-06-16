package handlers

import (
	"encoding/json"
	"net/http"
	"project/internal/db"
	"project/internal/scholarship"
)

func ScholarshipHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var scholarships []scholarship.Scholarship
	db.DB.Db.Find(&scholarships)
	jsonResp, _ := json.Marshal(scholarships)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}
