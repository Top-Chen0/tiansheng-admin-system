package Employeemanagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EmployeeStructRouter struct {
}

// InitEmployeeStructRouter 初始化 EmployeeStruct 路由信息
func (s *EmployeeStructRouter) InitEmployeeStructRouter(Router *gin.RouterGroup) {
	employeeStructRouter := Router.Group("employeeStruct").Use(middleware.OperationRecord())
	employeeStructRouterWithoutRecord := Router.Group("employeeStruct")
	var employeeStructApi = v1.ApiGroupApp.EmployeemanagementApiGroup.EmployeeStructApi
	{
		employeeStructRouter.POST("createEmployeeStruct", employeeStructApi.CreateEmployeeStruct)             // 新建EmployeeStruct
		employeeStructRouter.DELETE("deleteEmployeeStruct", employeeStructApi.DeleteEmployeeStruct)           // 删除EmployeeStruct
		employeeStructRouter.DELETE("deleteEmployeeStructByIds", employeeStructApi.DeleteEmployeeStructByIds) // 批量删除EmployeeStruct
		employeeStructRouter.PUT("updateEmployeeStruct", employeeStructApi.UpdateEmployeeStruct)              // 更新EmployeeStruct
	}
	{
		employeeStructRouterWithoutRecord.GET("findEmployeeStruct", employeeStructApi.FindEmployeeStruct)       // 根据ID获取EmployeeStruct
		employeeStructRouterWithoutRecord.GET("getEmployeeStructList", employeeStructApi.GetEmployeeStructList) // 获取EmployeeStruct列表
	}
}
