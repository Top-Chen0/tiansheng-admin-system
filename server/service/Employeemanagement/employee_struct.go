package Employeemanagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Employeemanagement"
	EmployeemanagementReq "github.com/flipped-aurora/gin-vue-admin/server/model/Employeemanagement/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EmployeeStructService struct {
}

// CreateEmployeeStruct 创建EmployeeStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (employeeStructService *EmployeeStructService) CreateEmployeeStruct(employeeStruct Employeemanagement.EmployeeStruct) (err error) {
	err = global.GVA_DB.Create(&employeeStruct).Error
	return err
}

// DeleteEmployeeStruct 删除EmployeeStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (employeeStructService *EmployeeStructService) DeleteEmployeeStruct(employeeStruct Employeemanagement.EmployeeStruct) (err error) {
	err = global.GVA_DB.Delete(&employeeStruct).Error
	return err
}

// DeleteEmployeeStructByIds 批量删除EmployeeStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (employeeStructService *EmployeeStructService) DeleteEmployeeStructByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Employeemanagement.EmployeeStruct{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEmployeeStruct 更新EmployeeStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (employeeStructService *EmployeeStructService) UpdateEmployeeStruct(employeeStruct Employeemanagement.EmployeeStruct) (err error) {
	err = global.GVA_DB.Save(&employeeStruct).Error
	return err
}

// GetEmployeeStruct 根据id获取EmployeeStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (employeeStructService *EmployeeStructService) GetEmployeeStruct(id uint) (err error, employeeStruct Employeemanagement.EmployeeStruct) {
	err = global.GVA_DB.Where("id = ?", id).First(&employeeStruct).Error
	return
}

// GetEmployeeStructInfoList 分页获取EmployeeStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (employeeStructService *EmployeeStructService) GetEmployeeStructInfoList(info EmployeemanagementReq.EmployeeStructSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Employeemanagement.EmployeeStruct{})
	var employeeStructs []Employeemanagement.EmployeeStruct
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Worknumber != nil {
		db = db.Where("worknumber <> ?", info.Worknumber)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Age != nil {
		db = db.Where("age = ?", info.Age)
	}
	if info.Sex != nil {
		db = db.Where("sex = ?", info.Sex)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&employeeStructs).Error
	return err, employeeStructs, total
}
