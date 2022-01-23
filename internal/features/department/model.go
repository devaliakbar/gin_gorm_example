package department

type CreateDepartmentInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateDepartmentInput struct {
	Name string `json:"name"`
}
