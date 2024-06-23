package patch

import (
	"bytes"
	"encoding/json"
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	"project/internal/models/dormitory"
)

func UpdateDormitoryHandler(w http.ResponseWriter, r *http.Request) {
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
	dormitory := dormitory.Dormitory{}
	json.Unmarshal(buf.Bytes(), &dormitory)
	database.DB.Db.Save(&dormitory)
	w.Write([]byte("updated"))
}
