// 自动生成模板OrderLnformation
package Order_Information

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// OrderLnformation 结构体
// 如果含有time.Time 请自行import time包
type OrderLnformation struct {
	global.GVA_MODEL
	Worknumber *int   `json:"worknumber" form:"worknumber" gorm:"column:worknumber;comment:工号;size:19;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:姓名;size:191;"`
	Age        *int   `json:"age" form:"age" gorm:"column:age;comment:年龄;size:19;"`
	Sex        *int   `json:"sex" form:"sex" gorm:"column:sex;comment:性别;size:19;"`
	Department *int   `json:"department" form:"department" gorm:"column:department;comment:部门;size:19;"`
	Job        *int   `json:"job" form:"job" gorm:"column:job;comment:岗位名称;size:19;"`
	//Dateofentry  *time.Time `json:"dateofentry" form:"dateofentry" gorm:"column:dateofentry;comment:入职日期;"`
	Socialsecuritynumber string `json:"socialsecuritynumber" form:"socialsecuritynumber" gorm:"column:socialsecuritynumber;comment:身份证号;size:191;"`
	//Dateofbirth  *time.Time `json:"dateofbirth" form:"dateofbirth" gorm:"column:dateofbirth;comment:出生日期;"`
	BirthMonth            *int       `json:"birthMonth" form:"birthMonth" gorm:"column:birth_month;comment:生日月份;size:19;"`
	Origin                string     `json:"origin" form:"origin" gorm:"column:origin;comment:籍贯;size:191;"`
	Householdregistration string     `json:"householdregistration" form:"householdregistration" gorm:"column:householdregistration;comment:户口所在地;size:191;"`
	Idaddress             string     `json:"idaddress" form:"idaddress" gorm:"column:idaddress;comment:身份证详细地址;size:191;"`
	Address               string     `json:"address" form:"address" gorm:"column:address;comment:现住详细地址;size:191;"`
	Ethnicgroup           string     `json:"ethnicgroup" form:"ethnicgroup" gorm:"column:ethnicgroup;comment:民族;size:191;"`
	Marriage              *int       `json:"marriage" form:"marriage" gorm:"column:marriage;comment:婚姻状况;size:19;"`
	Politicalface         *int       `json:"politicalface" form:"politicalface" gorm:"column:politicalface;comment:政治面貌;size:19;"`
	Emergencycontacts     string     `json:"emergencycontacts" form:"emergencycontacts" gorm:"column:emergencycontacts;comment:紧急联系人;size:191;"`
	EcNumber              *int       `json:"ecNumber" form:"ecNumber" gorm:"column:ec_number;comment:紧急联系人电话;size:19;"`
	Number                *int       `json:"number" form:"number" gorm:"column:number;comment:手机号码;size:19;"`
	EcRelationship        string     `json:"ecRelationship" form:"ecRelationship" gorm:"column:ec_relationship;comment:与联系人的关系;size:191;"`
	Degree                *int       `json:"degree" form:"degree" gorm:"column:degree;comment:学历;size:19;"`
	Graduatingschool      string     `json:"graduatingschool" form:"graduatingschool" gorm:"column:graduatingschool;comment:毕业院校;size:191;"`
	Graduationtime        *time.Time `json:"graduationtime" form:"graduationtime" gorm:"column:graduationtime;comment:毕业时间;"`
	Specialized           string     `json:"specialized" form:"specialized" gorm:"column:specialized;comment:专业;size:191;"`
	Specialty             string     `json:"specialty" form:"specialty" gorm:"column:specialty;comment:特长;size:191;"`
	Personalskills        string     `json:"personalskills" form:"personalskills" gorm:"column:personalskills;comment:个人技能;size:191;"`
	Professionaltitle     string     `json:"professionaltitle" form:"professionaltitle" gorm:"column:professionaltitle;comment:专业职称;size:191;"`
	Zip                   *int       `json:"zip" form:"zip" gorm:"column:zip;comment:邮编;size:19;"`
	Departmentmanager     *int       `json:"departmentmanager" form:"departmentmanager" gorm:"column:departmentmanager;comment:部门经理;size:19;"`
	Socialsecurity        *int       `json:"socialsecurity" form:"socialsecurity" gorm:"column:socialsecurity;comment:社保状态;size:19;"`
	EmployeeImg           string     `json:"employeeImg" form:"employeeImg" gorm:"column:employee_img;comment:员工照片;size:191;"`
}

// TableName OrderLnformation 表名
func (OrderLnformation) TableName() string {
	return "order_Information"
}
