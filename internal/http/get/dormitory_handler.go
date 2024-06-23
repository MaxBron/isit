package get

import (
	"encoding/json"
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	"project/internal/models/dormitory"
)

func DormitoryHandler(w http.ResponseWriter, r *http.Request) {
	if !post.IsAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var dormitories []dormitory.Dormitory
	database.DB.Db.Find(&dormitories)
	jsonResp, _ := json.Marshal(dormitories)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}
