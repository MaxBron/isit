package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"project/internal/db"
	"project/internal/user"
)

var isAuth bool
var role string

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	user := user.User{}
	json.Unmarshal(buf.Bytes(), &user)
	if db.DB.Db.First(&user, "login = ? AND password = ?", user.Login, user.Password).Error != nil {
		w.Write([]byte("wrong login or password"))
	} else {
		isAuth = true
		role = user.Role
		w.Write([]byte("OK"))
	}

}
