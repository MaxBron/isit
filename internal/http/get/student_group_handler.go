package get

import (
	"encoding/json"
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	studentGroup "project/internal/models/studentgroup"
)

func StudentGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !post.IsAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	var studentGroups []studentGroup.StudentGroup
	database.DB.Db.Find(&studentGroups)
	jsonResp, _ := json.Marshal(studentGroups)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(jsonResp))
}
