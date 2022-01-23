package test

import (
	"fmt"
	"testing"

	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/Valiben/gin_unit_test/utils"
	"github.com/devaliakbar/gin_gorm_example/internal/features/department"
	"github.com/devaliakbar/gin_gorm_example/internal/features/employee"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**Test create Employee**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestCreateEmployee(t *testing.T) {
	type resMdl struct {
		Success bool              `json:"success"`
		Data    employee.Employee `json:"data"`
	}

	var res resMdl

	///Testing without 'Department Id'(Required parameter)
	input := employee.CreateEmployeeInput{Name: "Test"}

	unitTest.TestHandlerUnMarshalResp(utils.POST, "/employee", "json", input, &res)

	if res.Success {
		t.Error("Should Fail to create employee, because department id is required")
		return
	}

	dep, depErr := createDepartment()

	if depErr != nil {
		t.Errorf("Failed to create department: %v\n", depErr)
		return
	}

	///Testing Create Employee without 'Name'(Required parameter)
	input = employee.CreateEmployeeInput{DepartmentId: int(dep.ID)}

	unitTest.TestHandlerUnMarshalResp(utils.POST, "/employee", "json", input, &res)

	if res.Success {
		t.Error("Should Fail to create employee, because employee name is required")
		return
	}

	///Testing Create Employee without all required parameter
	input = employee.CreateEmployeeInput{Name: "Marshall", DepartmentId: int(dep.ID)}

	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/employee", "json", input, &res)

	if err != nil {
		t.Errorf("Failed to request create employee: %v\n", err)
		return
	}

	if !res.Success || res.Data.Name != input.Name {
		t.Error("Fail to create employee")
		return
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**Test get all Employee**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestGetAllEmployee(t *testing.T) {
	type resMdl struct {
		Success bool                `json:"success"`
		Data    []employee.Employee `json:"data"`
	}

	var res resMdl

	err := unitTest.TestHandlerUnMarshalResp(utils.GET, "/employees", "json", nil, &res)

	if err != nil {
		t.Errorf("Failed to request get all employee: %v\n", err)
		return
	}

	if !res.Success || len(res.Data) == 0 {
		t.Error("Failed to get all employee")
		return
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**Test get all Employee**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestDeleteEmployee(t *testing.T) {
	type resMdl struct {
		Success bool              `json:"success"`
		Data    employee.Employee `json:"data"`
	}

	dep, depErr := createDepartment()

	if depErr != nil {
		t.Errorf("Failed to create department: %v\n", depErr)
		return
	}

	input := employee.CreateEmployeeInput{Name: "Marshall", DepartmentId: int(dep.ID)}

	var res resMdl

	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/employee", "json", input, &res)

	if err != nil {
		t.Errorf("Failed to request create employee: %v\n", err)
		return
	}

	///Deleting department
	var delRes resMdl

	err = unitTest.TestHandlerUnMarshalResp(utils.DELETE, fmt.Sprintf("/employee/%d", res.Data.ID), "json", nil, &delRes)

	if err != nil {
		t.Errorf("Failed to request delete employee: %v\n", err)
		return
	}

	if delRes.Data.ID != res.Data.ID {
		t.Error("Unexpected employee deleted")
		return
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**Create Department**///
///For creating dummy department
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func createDepartment() (department.Department, error) {
	///Creating Department for create Employee
	type depResMdl struct {
		Success bool                  `json:"success"`
		Data    department.Department `json:"data"`
	}

	var depRes depResMdl

	depInput := department.CreateDepartmentInput{Name: "Testt"}

	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/department", "json", depInput, &depRes)

	return depRes.Data, err
}
