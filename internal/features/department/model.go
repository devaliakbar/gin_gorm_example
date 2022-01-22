package department

type createDepartmentInput struct {
	Name string `json:"name" binding:"required"`
}

type updateDepartmentInput struct {
	Name string `json:"name"`
}
