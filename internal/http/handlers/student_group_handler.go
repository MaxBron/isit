package handlers

import (
	"encoding/json"
	"net/http"
	"project/internal/db"
	studentGroup "project/internal/studentgroup"
)

func StudentGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var studentGroups []studentGroup.StudentGroup
	db.DB.Db.Find(&studentGroups)
	jsonResp, _ := json.Marshal(studentGroups)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}
