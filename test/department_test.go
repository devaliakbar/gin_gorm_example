package test

import (
	"fmt"
	"testing"

	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/Valiben/gin_unit_test/utils"
	"github.com/devaliakbar/gin_gorm_example/internal/features/department"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**Test create department**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestCreateDepartment(t *testing.T) {
	type resMdl struct {
		Success bool                  `json:"success"`
		Data    department.Department `json:"data"`
	}

	var res resMdl

	///Testing without 'NAME'(Required parameter)
	input := department.CreateDepartmentInput{}

	unitTest.TestHandlerUnMarshalResp(utils.POST, "/department", "json", input, &res)

	if res.Success {
		t.Error("Should Fail to create employee, because department name is required")
		return
	}

	///Testing with correct data
	input = department.CreateDepartmentInput{Name: "Testt"}

	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/department", "json", input, &res)

	if err != nil {
		t.Errorf("Failed to request create employee: %v\n", err)
		return
	}

	if res.Data.Name != input.Name {
		t.Error("Unexpected department created")
		return
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**Test get all department**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestGetAllDepartment(t *testing.T) {
	type resMdl struct {
		Success bool                    `json:"success"`
		Data    []department.Department `json:"data"`
	}

	var res resMdl

	err := unitTest.TestHandlerUnMarshalResp(utils.GET, "/departments", "json", nil, &res)

	if err != nil {
		t.Errorf("Failed to request get all department: %v\n", err)
		return
	}

	if !res.Success || len(res.Data) == 0 {
		t.Error("Failed to get all department")
		return
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**Test delete department**///
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestDeleteDepartment(t *testing.T) {
	type resMdl struct {
		Success bool                  `json:"success"`
		Data    department.Department `json:"data"`
	}

	///Creating a department for testing delete
	dep, depErr := createDepartment()

	if depErr != nil {
		t.Errorf("Failed to create department: %v\n", depErr)
		return
	}

	///Deleting department
	var delRes resMdl

	err := unitTest.TestHandlerUnMarshalResp(utils.DELETE, fmt.Sprintf("/department/%d", dep.ID), "json", nil, &delRes)

	if err != nil {
		t.Errorf("Failed to request delete employee: %v\n", err)
		return
	}

	if delRes.Data.ID != dep.ID {
		t.Error("Unexpected department deleted")
		return
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///**Create Department**///
///For creating dummy department (might use on other files)
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func createDepartment() (department.Department, error) {
	type resMdl struct {
		Success bool                  `json:"success"`
		Data    department.Department `json:"data"`
	}

	var depRes resMdl

	depInput := department.CreateDepartmentInput{Name: "Testt"}

	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/department", "json", depInput, &depRes)

	return depRes.Data, err
}
