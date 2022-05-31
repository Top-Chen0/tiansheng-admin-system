// 自动生成模板EmployeeStructui
package Employeemanagementui

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

//操作数据库增加字段
// EmployeeStructui 结构体
// 如果含有time.Time 请自行import time包
type EmployeeStructui struct {
	global.GVA_MODEL
	Worknumber            *int       `json:"worknumber" form:"worknumber" gorm:"column:worknumber;comment:工号;size:19;"`
	Name                  string     `json:"name" form:"name" gorm:"column:name;comment:姓名;size:191;"`
	Age                   *int       `json:"age" form:"age" gorm:"column:age;comment:年龄;"`
	Sex                   *int       `json:"sex" form:"sex" gorm:"column:sex;comment:性别;size:19;"`
	Department            *int       `json:"department" form:"department" gorm:"column:department;comment:部门;"`
	Job                   *int       `json:"job" form:"job" gorm:"column:job;comment:岗位名称;"`
	Dateofentry           *time.Time `json:"dateofentry" form:"dateofentry" gorm:"column:dateofentry;comment:入职日期;"`
	Socialsecuritynumber  string     `json:"socialsecuritynumber" form:"socialsecuritynumber" gorm:"column:socialsecuritynumber;comment:身份证号;"`
	Dateofbirth           *time.Time `json:"dateofbirth" form:"dateofbirth" gorm:"column:dateofbirth;comment:出生日期;"`
	Birth_month           *int       `json:"birth_month" form:"birth_month" gorm:"column:birth_month;comment:生日月份;"`
	Origin                string     `json:"origin" form:"origin" gorm:"column:origin;comment:籍贯;"`
	Householdregistration string     `json:"householdregistration" form:"householdregistration" gorm:"column:householdregistration;comment:户口所在地;"`
	Idaddress             string     `json:"idaddress" form:"idaddress" gorm:"column:idaddress;comment:身份证详细地址;"`
	Address               string     `json:"address" form:"address" gorm:"column:address;comment:现住详细地址;"`
	Ethnicgroup           string     `json:"ethnicgroup" form:"ethnicgroup" gorm:"column:ethnicgroup;comment:民族;"`
	Marriage              *int       `json:"marriage" form:"marriage" gorm:"column:marriage;comment:婚姻状况;"`
	Politicalface         *int       `json:"politicalface" form:"politicalface" gorm:"column:politicalface;comment:政治面貌;"`
	Emergencycontacts     string     `json:"emergencycontacts" form:"emergencycontacts" gorm:"column:emergencycontacts;comment:紧急联系人;"`
	Ec_number             *int       `json:"ec_number" form:"ec_number" gorm:"column:ec_number;comment:紧急联系人电话;size:11;"`
	Number                *int       `json:"number" form:"number" gorm:"column:number;comment:手机号码;size:11;"`
	Ec_relationship       string     `json:"ec_relationship" form:"ec_relationship" gorm:"column:ec_relationship;comment:与联系人的关系;"`
	Degree                *int       `json:"degree" form:"degree" gorm:"column:degree;comment:学历;"`
	Graduatingschool      string     `json:"graduatingschool" form:"graduatingschool" gorm:"column:graduatingschool;comment:毕业院校;"`
	Graduationtime        *time.Time `json:"graduationtime" form:"graduationtime" gorm:"column:graduationtime;comment:毕业时间;"`
	Specialized           string     `json:"specialized" form:"specialized" gorm:"column:specialized;comment:专业;"`
	Specialty             string     `json:"specialty" form:"specialty" gorm:"column:specialty;comment:特长;"`
	Personalskills        string     `json:"personalskills" form:"personalskills" gorm:"column:personalskills;comment:个人技能;"`
	Professionaltitle     string     `json:"professionaltitle" form:"professionaltitle" gorm:"column:professionaltitle;comment:专业职称;"`
	Zip                   *int       `json:"zip" form:"zip" gorm:"column:zip;comment:邮编;size:6;"`
	Departmentmanager     *int       `json:"departmentmanager" form:"departmentmanager" gorm:"column:departmentmanager;comment:部门经理;"`
	Socialsecurity        *int       `json:"socialsecurity" form:"socialsecurity" gorm:"column:socialsecurity;comment:社保状态;"`
	EmployeeImg           string     `json:"employeeImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:员工照片;"`
}

// TableName EmployeeStructui 表名
func (EmployeeStructui) TableName() string {
	return "employee_struct"
}
