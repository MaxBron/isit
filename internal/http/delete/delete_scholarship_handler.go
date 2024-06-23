package delete

import (
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	"project/internal/models/scholarship"
)

func DeleteScholarshipHandler(w http.ResponseWriter, r *http.Request) {

	if !post.IsAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	if post.Role != "admin" {
		w.Write([]byte("No access"))
		return
	}

	typeScholarship := r.URL.Query().Get("type")
	database.DB.Db.Unscoped().Delete(&scholarship.Scholarship{}, typeScholarship)
	w.Write([]byte("deleted"))
}
