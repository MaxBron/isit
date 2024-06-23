package models

import "gorm.io/gorm"

type Dormitory struct {
	gorm.Model
	RoomNumber int `json:"room_number"`
	StudentId  int `json:"student_id"`
}
