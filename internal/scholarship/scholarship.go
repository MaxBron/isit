package scholarship

import (
	"gorm.io/gorm"
)

type Scholarship struct {
	gorm.Model
	Type      string `json:"type"`
	StudentId int    `json:"student_id"`
	Amount    int    `json:"amount"`
}
