package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Employeemanagementui"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EmployeeStructuiSearch struct {
	Employeemanagementui.EmployeeStructui
	request.PageInfo
}
