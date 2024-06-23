package models

import (
	"gorm.io/gorm"
)

type Scholarship struct {
	gorm.Model
	Type   string `json:"type"`
	Amount int    `json:"amount"`
}
