package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Employeemanagement"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Employeemanagementui"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup               system.ApiGroup
	ExampleApiGroup              example.ApiGroup
	EmployeemanagementApiGroup   Employeemanagement.ApiGroup
	EmployeemanagementuiApiGroup Employeemanagementui.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
