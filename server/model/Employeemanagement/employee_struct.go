// 自动生成模板EmployeeStruct
package Employeemanagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// EmployeeStruct 结构体
// 如果含有time.Time 请自行import time包
type EmployeeStruct struct {
	global.GVA_MODEL
	Worknumber *int   `json:"worknumber" form:"worknumber" gorm:"column:worknumber;comment:工号;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:姓名;size:191;"`
	Age        *int   `json:"age" form:"age" gorm:"column:age;comment:年龄;size:3;"`
	Sex        *int   `json:"sex" form:"sex" gorm:"column:sex;comment:性别;"`
}

// TableName EmployeeStruct 表名
func (EmployeeStruct) TableName() string {
	return "employee_struct"
}
