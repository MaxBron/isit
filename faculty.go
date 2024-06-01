package main

import "gorm.io/gorm"

type Faculty struct {
	gorm.Model
	Name    string `json:"name"`
	Catalog string `json:"catalog"`
}

func (faculty Faculty) AddFaculty() string {
	DB.Db.Create(&faculty)
	return faculty.Name
}
