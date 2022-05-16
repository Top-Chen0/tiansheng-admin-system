package Employeemanagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Employeemanagement"
	EmployeemanagementReq "github.com/flipped-aurora/gin-vue-admin/server/model/Employeemanagement/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EmployeeStructApi struct {
}

var employeeStructService = service.ServiceGroupApp.EmployeemanagementServiceGroup.EmployeeStructService

// CreateEmployeeStruct 创建EmployeeStruct
// @Tags EmployeeStruct
// @Summary 创建EmployeeStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Employeemanagement.EmployeeStruct true "创建EmployeeStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /employeeStruct/createEmployeeStruct [post]
func (employeeStructApi *EmployeeStructApi) CreateEmployeeStruct(c *gin.Context) {
	var employeeStruct Employeemanagement.EmployeeStruct
	_ = c.ShouldBindJSON(&employeeStruct)
	if err := employeeStructService.CreateEmployeeStruct(employeeStruct); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEmployeeStruct 删除EmployeeStruct
// @Tags EmployeeStruct
// @Summary 删除EmployeeStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Employeemanagement.EmployeeStruct true "删除EmployeeStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /employeeStruct/deleteEmployeeStruct [delete]
func (employeeStructApi *EmployeeStructApi) DeleteEmployeeStruct(c *gin.Context) {
	var employeeStruct Employeemanagement.EmployeeStruct
	_ = c.ShouldBindJSON(&employeeStruct)
	if err := employeeStructService.DeleteEmployeeStruct(employeeStruct); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEmployeeStructByIds 批量删除EmployeeStruct
// @Tags EmployeeStruct
// @Summary 批量删除EmployeeStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EmployeeStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /employeeStruct/deleteEmployeeStructByIds [delete]
func (employeeStructApi *EmployeeStructApi) DeleteEmployeeStructByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := employeeStructService.DeleteEmployeeStructByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEmployeeStruct 更新EmployeeStruct
// @Tags EmployeeStruct
// @Summary 更新EmployeeStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Employeemanagement.EmployeeStruct true "更新EmployeeStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /employeeStruct/updateEmployeeStruct [put]
func (employeeStructApi *EmployeeStructApi) UpdateEmployeeStruct(c *gin.Context) {
	var employeeStruct Employeemanagement.EmployeeStruct
	_ = c.ShouldBindJSON(&employeeStruct)
	if err := employeeStructService.UpdateEmployeeStruct(employeeStruct); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEmployeeStruct 用id查询EmployeeStruct
// @Tags EmployeeStruct
// @Summary 用id查询EmployeeStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Employeemanagement.EmployeeStruct true "用id查询EmployeeStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /employeeStruct/findEmployeeStruct [get]
func (employeeStructApi *EmployeeStructApi) FindEmployeeStruct(c *gin.Context) {
	var employeeStruct Employeemanagement.EmployeeStruct
	_ = c.ShouldBindQuery(&employeeStruct)
	if err, reemployeeStruct := employeeStructService.GetEmployeeStruct(employeeStruct.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reemployeeStruct": reemployeeStruct}, c)
	}
}

// GetEmployeeStructList 分页获取EmployeeStruct列表
// @Tags EmployeeStruct
// @Summary 分页获取EmployeeStruct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EmployeemanagementReq.EmployeeStructSearch true "分页获取EmployeeStruct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /employeeStruct/getEmployeeStructList [get]
func (employeeStructApi *EmployeeStructApi) GetEmployeeStructList(c *gin.Context) {
	var pageInfo EmployeemanagementReq.EmployeeStructSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := employeeStructService.GetEmployeeStructInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
