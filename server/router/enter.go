package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Employeemanagement"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Employeemanagementui"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Order_Information"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Order_information_inquiry"
	"github.com/flipped-aurora/gin-vue-admin/server/router/emplyeeStructDetail0"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System                    system.RouterGroup
	Example                   example.RouterGroup
	Employeemanagement        Employeemanagement.RouterGroup
	Employeemanagementui      Employeemanagementui.RouterGroup
	Order_Information         Order_Information.RouterGroup
	Order_information_inquiry Order_information_inquiry.RouterGroup
	EmplyeeStructDetail0      emplyeeStructDetail0.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
