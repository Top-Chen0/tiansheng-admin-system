package Employeemanagementui

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Employeemanagementui"
	EmployeemanagementuiReq "github.com/flipped-aurora/gin-vue-admin/server/model/Employeemanagementui/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EmployeeStructuiService struct {
}

// CreateEmployeeStructui 创建EmployeeStructui记录
// Author [piexlmax](https://github.com/piexlmax)
func (employeeStructuiService *EmployeeStructuiService) CreateEmployeeStructui(employeeStructui Employeemanagementui.EmployeeStructui) (err error) {
	err = global.GVA_DB.Create(&employeeStructui).Error
	return err
}

// DeleteEmployeeStructui 删除EmployeeStructui记录
// Author [piexlmax](https://github.com/piexlmax)
func (employeeStructuiService *EmployeeStructuiService) DeleteEmployeeStructui(employeeStructui Employeemanagementui.EmployeeStructui) (err error) {
	err = global.GVA_DB.Delete(&employeeStructui).Error
	return err
}

// DeleteEmployeeStructuiByIds 批量删除EmployeeStructui记录
// Author [piexlmax](https://github.com/piexlmax)
func (employeeStructuiService *EmployeeStructuiService) DeleteEmployeeStructuiByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Employeemanagementui.EmployeeStructui{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEmployeeStructui 更新EmployeeStructui记录
// Author [piexlmax](https://github.com/piexlmax)
func (employeeStructuiService *EmployeeStructuiService) UpdateEmployeeStructui(employeeStructui Employeemanagementui.EmployeeStructui) (err error) {
	err = global.GVA_DB.Save(&employeeStructui).Error
	return err
}

// GetEmployeeStructui 根据id获取EmployeeStructui记录
// Author [piexlmax](https://github.com/piexlmax)
func (employeeStructuiService *EmployeeStructuiService) GetEmployeeStructui(id uint) (err error, employeeStructui Employeemanagementui.EmployeeStructui) {
	err = global.GVA_DB.Where("id = ?", id).First(&employeeStructui).Error
	return
}

// GetEmployeeStructuiInfoList 分页获取EmployeeStructui记录
// Author [piexlmax](https://github.com/piexlmax)
func (employeeStructuiService *EmployeeStructuiService) GetEmployeeStructuiInfoList(info EmployeemanagementuiReq.EmployeeStructuiSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Employeemanagementui.EmployeeStructui{})
	var employeeStructuis []Employeemanagementui.EmployeeStructui
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Worknumber != nil {
		db = db.Where("worknumber = ?", info.Worknumber)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Age != nil {
		db = db.Where("age = ?", info.Age)
	}
	if info.Age != nil {
		db = db.Where("age > ? and age < ?", info.Age)
	}
	if info.Sex != nil {
		db = db.Where("sex = ?", info.Sex)
	}
	if info.Department != nil {
		db = db.Where("department = ?", info.Department)
	}
	if info.Job != nil {
		db = db.Where("job = ?", info.Job)
	}
	if info.Dateofentry != nil {
		db = db.Where("dateofentry = ?", info.Dateofentry)
	}
	if info.Socialsecuritynumber != "" {
		db = db.Where("socialsecuritynumber LIKE ?", "%"+info.Socialsecuritynumber+"%")
	}
	if info.Dateofbirth != nil {
		db = db.Where("dateofbirth = ?", info.Dateofbirth)
	}
	if info.Birth_month != nil {
		db = db.Where("birth_month = ?", info.Birth_month)
	}
	if info.Origin != "" {
		db = db.Where("origin LIKE ?", "%"+info.Origin+"%")
	}
	if info.Householdregistration != "" {
		db = db.Where("householdregistration LIKE ?", "%"+info.Householdregistration+"%")
	}
	if info.Idaddress != "" {
		db = db.Where("idaddress LIKE ?", "%"+info.Idaddress+"%")
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.Ethnicgroup != "" {
		db = db.Where("ethnicgroup = ?", info.Ethnicgroup)
	}
	if info.Marriage != nil {
		db = db.Where("marriage = ?", info.Marriage)
	}
	if info.Politicalface != nil {
		db = db.Where("politicalface = ?", info.Politicalface)
	}
	if info.Emergencycontacts != "" {
		db = db.Where("emergencycontacts LIKE ?", "%"+info.Emergencycontacts+"%")
	}
	if info.Ec_number != nil {
		db = db.Where("ec_number = ?", info.Ec_number)
	}
	if info.Number != nil {
		db = db.Where("number = ?", info.Number)
	}
	if info.Ec_relationship != "" {
		db = db.Where("ec_relationship LIKE ?", "%"+info.Ec_relationship+"%")
	}
	if info.Degree != nil {
		db = db.Where("degree = ?", info.Degree)
	}
	if info.Graduatingschool != "" {
		db = db.Where("graduatingschool LIKE ?", "%"+info.Graduatingschool+"%")
	}
	if info.Graduationtime != nil {
		db = db.Where("graduationtime = ?", info.Graduationtime)
	}
	if info.Specialized != "" {
		db = db.Where("specialized LIKE ?", "%"+info.Specialized+"%")
	}
	if info.Specialty != "" {
		db = db.Where("specialty LIKE ?", "%"+info.Specialty+"%")
	}
	if info.Personalskills != "" {
		db = db.Where("personalskills LIKE ?", "%"+info.Personalskills+"%")
	}
	if info.Professionaltitle != "" {
		db = db.Where("professionaltitle LIKE ?", "%"+info.Professionaltitle+"%")
	}
	if info.Zip != nil {
		db = db.Where("zip = ?", info.Zip)
	}
	if info.Departmentmanager != nil {
		db = db.Where("departmentmanager = ?", info.Departmentmanager)
	}
	if info.Socialsecurity != nil {
		db = db.Where("socialsecurity = ?", info.Socialsecurity)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&employeeStructuis).Error
	return err, employeeStructuis, total
}
