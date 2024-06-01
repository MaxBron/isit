package main

import "gorm.io/gorm"

type Scholarship struct {
	gorm.Model
	Type      string `json:"type"`
	StudentId int    `json:"student_id"`
	Amount    int    `json:"amount"`
}

func (scholarship Scholarship) AddScholarship() (int, string) {
	DB.Db.Create(&scholarship)
	return scholarship.StudentId, scholarship.Type
}
