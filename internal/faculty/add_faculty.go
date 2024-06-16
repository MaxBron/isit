package faculty

import "project/internal/db"

func (faculty Faculty) AddFaculty() string {
	db.DB.Db.Create(&faculty)
	return faculty.Name
}
