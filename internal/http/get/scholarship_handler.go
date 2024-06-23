package get

import (
	"encoding/json"
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	"project/internal/models/scholarship"
)

func ScholarshipHandler(w http.ResponseWriter, r *http.Request) {
	if !post.IsAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var scholarships []scholarship.Scholarship
	database.DB.Db.Find(&scholarships)
	jsonResp, _ := json.Marshal(scholarships)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}
