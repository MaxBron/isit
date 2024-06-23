package get

import (
	"encoding/json"
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	"project/internal/models/student"
)

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	if !post.IsAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var students []student.Student
	database.DB.Db.Find(&students)
	jsonResp, _ := json.Marshal(students)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}
