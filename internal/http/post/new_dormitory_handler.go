package post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"project/internal/database"
	"project/internal/models"
)

func NewDormitoryHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuth {
		w.Write([]byte("Please, log in"))
		return
	}

	if Role != "admin" {
		w.Write([]byte("No access"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	dormitory := models.Dormitory{}
	json.Unmarshal(buf.Bytes(), &dormitory)
	database.DB.Db.Create(&dormitory)
	w.Header().Set("Content-Type", "applictaion/json")
	w.Write([]byte(fmt.Sprintf(`{"room_number":"%d"}`, dormitory.RoomNumber)))
}
