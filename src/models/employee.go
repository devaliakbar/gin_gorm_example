package models

type Employee struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name"`
	DepartmentId uint
	Department   Department
}
