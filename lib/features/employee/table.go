package employee

import (
	"github.com/devaliakbar/gin_gorm_example/lib/features/department"
)

type Employee struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name"`
	DepartmentId uint
	Department   department.Department
}
