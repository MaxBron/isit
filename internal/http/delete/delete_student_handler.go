package delete

import (
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	student "project/internal/models/student"
	"strconv"
)

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	if !post.IsAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	if post.Role != "admin" {
		w.Write([]byte("No access"))
		return
	}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	database.DB.Db.Unscoped().Delete(&student.Student{}, id)
	w.Write([]byte("deleted"))
}
