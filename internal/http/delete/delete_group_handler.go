package delete

import (
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	studentGroup "project/internal/models/studentgroup"
	"strconv"
)

func DeleteGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !post.IsAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	if post.Role != "admin" {
		w.Write([]byte("No access"))
		return
	}

	number, _ := strconv.Atoi(r.URL.Query().Get("number"))
	database.DB.Db.Unscoped().Delete(&studentGroup.StudentGroup{}, number)
	w.Write([]byte("deleted"))
}
