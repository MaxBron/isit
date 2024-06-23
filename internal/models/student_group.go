package models

import (
	"gorm.io/gorm"
)

type StudentGroup struct {
	gorm.Model
	Number      int    `json:"number"`
	FacultyName string `json:"faculty_name"`
}
