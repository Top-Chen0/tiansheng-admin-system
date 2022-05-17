package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Employeemanagement"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Employeemanagementui"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System               system.RouterGroup
	Example              example.RouterGroup
	Employeemanagement   Employeemanagement.RouterGroup
	Employeemanagementui Employeemanagementui.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
