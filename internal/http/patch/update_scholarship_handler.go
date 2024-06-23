package patch

import (
	"bytes"
	"encoding/json"
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	"project/internal/models/scholarship"
)

func UpdateScholarshipHandler(w http.ResponseWriter, r *http.Request) {
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
	scholarship := scholarship.Scholarship{}
	json.Unmarshal(buf.Bytes(), &scholarship)
	database.DB.Db.Save(&scholarship)
	w.Write([]byte("updated"))
}
