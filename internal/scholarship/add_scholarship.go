package scholarship

import "project/internal/db"

func (scholarship Scholarship) AddScholarship() (int, string) {
	db.DB.Db.Create(&scholarship)
	return scholarship.StudentId, scholarship.Type
}
