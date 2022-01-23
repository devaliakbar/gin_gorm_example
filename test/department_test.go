package test

import (
	"testing"

	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/Valiben/gin_unit_test/utils"
	"github.com/devaliakbar/gin_gorm_example/internal/features/department"
)

func TestCreateDepartment(t *testing.T) {
	input := department.CreateDepartmentInput{Name: "Testt"}

	var res department.Department

	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/department", "json", input, &res)

	if err != nil {
		t.Errorf("TestPostMappingClientNotFound: %v\n", err)
		return
	}

	t.Log(res)
}

func TestGetAllDepartment(t *testing.T) {
	var res []department.Department

	err := unitTest.TestHandlerUnMarshalResp(utils.GET, "/departments", "json", nil, &res)

	if err != nil {
		t.Errorf("TestPostMappingClientNotFound: %v\n", err)
		return
	}

	t.Log(res)
}
