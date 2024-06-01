package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Login    string `json:"login"`
	Password string `json:"password"`
	role     string `json:"role"`
}