package patch

import (
	"bytes"
	"encoding/json"
	"net/http"
	"project/internal/database"
	"project/internal/models/student"
)

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	// if !post.IsAuth {
	// 	w.Write([]byte("Please, log in"))
	// 	return
	// }

	// if post.Role != "admin" {
	// 	w.Write([]byte("No access"))
	// 	return
	// }

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	student := student.Student{}
	json.Unmarshal(buf.Bytes(), &student)
	database.DB.Db.Save(&student)
	w.Write([]byte("updated"))
}
