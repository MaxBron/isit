package faculty

import (
	"gorm.io/gorm"
)

type Faculty struct {
	gorm.Model
	Name    string `json:"name"`
	Catalog string `json:"catalog"`
}
