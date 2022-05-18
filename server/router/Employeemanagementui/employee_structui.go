package Employeemanagementui

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EmployeeStructuiRouter struct {
}

// InitEmployeeStructuiRouter 初始化 EmployeeStructui 路由信息
func (s *EmployeeStructuiRouter) InitEmployeeStructuiRouter(Router *gin.RouterGroup) {
	employeeStructuiRouter := Router.Group("employeeStructui").Use(middleware.OperationRecord())
	employeeStructuiRouterWithoutRecord := Router.Group("employeeStructui")
	var employeeStructuiApi = v1.ApiGroupApp.EmployeemanagementuiApiGroup.EmployeeStructuiApi
	{
		employeeStructuiRouter.POST("createEmployeeStructui", employeeStructuiApi.CreateEmployeeStructui)             // 新建EmployeeStructui
		employeeStructuiRouter.DELETE("deleteEmployeeStructui", employeeStructuiApi.DeleteEmployeeStructui)           // 删除EmployeeStructui
		employeeStructuiRouter.DELETE("deleteEmployeeStructuiByIds", employeeStructuiApi.DeleteEmployeeStructuiByIds) // 批量删除EmployeeStructui
		employeeStructuiRouter.PUT("updateEmployeeStructui", employeeStructuiApi.UpdateEmployeeStructui)              // 更新EmployeeStructui
	}
	{
		employeeStructuiRouterWithoutRecord.GET("findEmployeeStructui", employeeStructuiApi.FindEmployeeStructui)       // 根据ID获取EmployeeStructui
		employeeStructuiRouterWithoutRecord.GET("getEmployeeStructuiList", employeeStructuiApi.GetEmployeeStructuiList) // 获取EmployeeStructui列表
	}
}
