package student

import "project/internal/db"

func (student Student) AddStudent() int {
	db.DB.Db.Create(&student)
	return student.Id
}
