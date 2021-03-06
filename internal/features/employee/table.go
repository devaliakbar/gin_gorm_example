package employee

import (
	"github.com/devaliakbar/gin_gorm_example/internal/features/department"
)

type Employee struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name" gorm:"not null"`
	DepartmentId uint   `json:"department_id" gorm:"not null"`
	Department   department.Department
}
