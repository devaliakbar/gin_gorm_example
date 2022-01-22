package employee

type employeeSelection struct {
	EmployeeId         string `json:"employee_id"`
	EmployeeName       string `json:"employee_name"`
	EmployeeDepartment string `json:"employee_department"`
}

type updateEmployeeInput struct {
	Name         string `json:"name"`
	DepartmentId int    `json:"department_id"`
}

type createEmployeeInput struct {
	Name         string `json:"name" binding:"required"`
	DepartmentId int    `json:"department_id"`
}
