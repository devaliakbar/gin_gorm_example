package employeeModels

import departmentModels "github.com/devaliakbar/gin_gorm_example/lib/features/department/models"

type Employee struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name"`
	DepartmentId uint
	Department   departmentModels.Department
}
