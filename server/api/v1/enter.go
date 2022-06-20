package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Employeemanagement"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Employeemanagementui"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Order_Information"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Order_information_inquiry"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/emplyeeStructDetail0"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup                    system.ApiGroup
	ExampleApiGroup                   example.ApiGroup
	EmployeemanagementApiGroup        Employeemanagement.ApiGroup
	EmployeemanagementuiApiGroup      Employeemanagementui.ApiGroup
	Order_InformationApiGroup         Order_Information.ApiGroup
	Order_information_inquiryApiGroup Order_information_inquiry.ApiGroup
	EmplyeeStructDetail0ApiGroup      emplyeeStructDetail0.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
