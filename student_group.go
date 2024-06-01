package main

import "gorm.io/gorm"

type StudentGroup struct {
	gorm.Model
	Number      int    `json:"number"`
	FacultyName string `json:"faculty_name"`
}

func (studentGroup StudentGroup) AddStudentGroup() int {
	DB.Db.Create(&studentGroup)
	return studentGroup.Number
}
