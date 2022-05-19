package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Employeemanagementui"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

//搜索结构体
type EmployeeStructuiSearch struct {
	Employeemanagementui.EmployeeStructui
	request.PageInfo
}
