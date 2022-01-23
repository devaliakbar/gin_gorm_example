package employee

type EmployeeSelection struct {
	EmployeeId         string `json:"employee_id"`
	EmployeeName       string `json:"employee_name"`
	EmployeeDepartment string `json:"employee_department"`
}

type UpdateEmployeeInput struct {
	Name         string `json:"name"`
	DepartmentId int    `json:"department_id"`
}

type CreateEmployeeInput struct {
	Name         string `json:"name" binding:"required"`
	DepartmentId int    `json:"department_id"`
}
