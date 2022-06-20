package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/Employeemanagement"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Employeemanagementui"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Order_Information"
	"github.com/flipped-aurora/gin-vue-admin/server/service/Order_information_inquiry"
	"github.com/flipped-aurora/gin-vue-admin/server/service/emplyeeStructDetail0"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup                    system.ServiceGroup
	ExampleServiceGroup                   example.ServiceGroup
	EmployeemanagementServiceGroup        Employeemanagement.ServiceGroup
	EmployeemanagementuiServiceGroup      Employeemanagementui.ServiceGroup
	Order_InformationServiceGroup         Order_Information.ServiceGroup
	Order_information_inquiryServiceGroup Order_information_inquiry.ServiceGroup
	EmplyeeStructDetail0ServiceGroup      emplyeeStructDetail0.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
