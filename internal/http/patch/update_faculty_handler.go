package patch

import (
	"bytes"
	"encoding/json"
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	"project/internal/models/faculty"
)

func UpdateFacultyHandler(w http.ResponseWriter, r *http.Request) {
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
	faculty := faculty.Faculty{}
	json.Unmarshal(buf.Bytes(), &faculty)
	// database.DB.Db.Model(&faculty).Update
	database.DB.Db.Save(&faculty)
	w.Write([]byte("updated"))
}
