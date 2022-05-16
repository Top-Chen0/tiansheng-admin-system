// 自动生成模板EmployeeStruct
package Employeemanagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// EmployeeStruct 结构体
// 如果含有time.Time 请自行import time包
type EmployeeStruct struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:员工姓名;"`
}

// TableName EmployeeStruct 表名
func (EmployeeStruct) TableName() string {
	return "employee_Struct"
}
