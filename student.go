package main

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	GroupNumber  int    `json:"group_number"`
	Location     string `json:"location"`
	CourseNumber int    `json:"course_number"`
	SurName      string `json:"sur_name"`
	LastName     string `json:"last_name"`
	FirstName    string `json:"first_name"`
	Id           int    `json:"id"`
}

func (student Student) AddStudent() int {
	DB.Db.Create(&student)
	return student.Id
}
