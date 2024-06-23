package delete

import (
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	"project/internal/models/faculty"
)

func DeleteFacultyHandler(w http.ResponseWriter, r *http.Request) {
	if !post.IsAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	if post.Role != "admin" {
		w.Write([]byte("No access"))
		return
	}

	name := r.URL.Query().Get("name")
	database.DB.Db.Unscoped().Delete(&faculty.Faculty{}, name)
	w.Write([]byte("deleted"))
}
