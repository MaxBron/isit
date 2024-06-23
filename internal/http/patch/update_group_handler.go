package patch

import (
	"bytes"
	"encoding/json"
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	studentGroup "project/internal/models/studentgroup"
)

func UpdateGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !post.IsAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	if post.Role != "admin" {
		w.Write([]byte("No access"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	group := studentGroup.StudentGroup{}
	json.Unmarshal(buf.Bytes(), &group)
	database.DB.Db.Save(&group)
	w.Write([]byte("updated"))
}
