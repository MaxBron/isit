package dormitory

import "project/internal/db"

func (dormitory Dormitory) AddDormitory() int {
	db.DB.Db.Create(&dormitory)
	return dormitory.RoomNumber
}
