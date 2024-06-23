package post

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"project/internal/database"
	"project/internal/models"
)

var IsAuth bool
var Role string

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	user := models.User{}
	json.Unmarshal(buf.Bytes(), &user)
	hashPassword := sha256.Sum256([]byte(user.Password))
	hashString := hex.EncodeToString(hashPassword[:])
	if database.DB.Db.First(&user, "login = ? AND password = ?", user.Login, hashString).Error != nil {
		w.Write([]byte("wrong login or password"))
	} else {
		IsAuth = true
		Role = user.Role
		w.Write([]byte("OK"))
	}

}
