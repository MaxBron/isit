package delete

import (
	"net/http"
	"project/internal/database"
	"project/internal/http/handlers/post"
	"project/internal/models"
	"strconv"
)

func DeleteDormitoryHandler(w http.ResponseWriter, r *http.Request) {
	if !post.IsAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	if post.Role != "admin" {
		w.Write([]byte("No access"))
		return
	}

	number, _ := strconv.Atoi(r.URL.Query().Get("number"))
	database.DB.Db.Unscoped().Delete(&models.Dormitory{}, number)
	w.Write([]byte("deleted"))
}
