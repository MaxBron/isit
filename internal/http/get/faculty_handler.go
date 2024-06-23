package get

import (
	"encoding/json"
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	"project/internal/models/faculty"
)

func FacultyHandler(w http.ResponseWriter, r *http.Request) {
	if !post.IsAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var faculties []faculty.Faculty
	database.DB.Db.Find(&faculties)
	jsonResp, _ := json.Marshal(faculties)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}
