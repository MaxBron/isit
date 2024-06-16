package studentGroup

import "project/internal/db"

func (studentGroup StudentGroup) AddStudentGroup() int {
	db.DB.Db.Create(&studentGroup)
	return studentGroup.Number
}
