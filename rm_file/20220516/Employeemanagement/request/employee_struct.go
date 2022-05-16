package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Employeemanagement"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EmployeeStructSearch struct {
	Employeemanagement.EmployeeStruct
	request.PageInfo
}
