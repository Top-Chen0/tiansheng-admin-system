package Employeemanagementui

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Employeemanagementui"
	EmployeemanagementuiReq "github.com/flipped-aurora/gin-vue-admin/server/model/Employeemanagementui/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EmployeeStructuiApi struct {
}

var employeeStructuiService = service.ServiceGroupApp.EmployeemanagementuiServiceGroup.EmployeeStructuiService

// CreateEmployeeStructui 创建EmployeeStructui
// @Tags EmployeeStructui
// @Summary 创建EmployeeStructui
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Employeemanagementui.EmployeeStructui true "创建EmployeeStructui"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /employeeStructui/createEmployeeStructui [post]
func (employeeStructuiApi *EmployeeStructuiApi) CreateEmployeeStructui(c *gin.Context) {
	var employeeStructui Employeemanagementui.EmployeeStructui
	_ = c.ShouldBindJSON(&employeeStructui)
	if err := employeeStructuiService.CreateEmployeeStructui(employeeStructui); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEmployeeStructui 删除EmployeeStructui
// @Tags EmployeeStructui
// @Summary 删除EmployeeStructui
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Employeemanagementui.EmployeeStructui true "删除EmployeeStructui"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /employeeStructui/deleteEmployeeStructui [delete]
func (employeeStructuiApi *EmployeeStructuiApi) DeleteEmployeeStructui(c *gin.Context) {
	var employeeStructui Employeemanagementui.EmployeeStructui
	_ = c.ShouldBindJSON(&employeeStructui)
	if err := employeeStructuiService.DeleteEmployeeStructui(employeeStructui); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEmployeeStructuiByIds 批量删除EmployeeStructui
// @Tags EmployeeStructui
// @Summary 批量删除EmployeeStructui
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EmployeeStructui"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /employeeStructui/deleteEmployeeStructuiByIds [delete]
func (employeeStructuiApi *EmployeeStructuiApi) DeleteEmployeeStructuiByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := employeeStructuiService.DeleteEmployeeStructuiByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEmployeeStructui 更新EmployeeStructui
// @Tags EmployeeStructui
// @Summary 更新EmployeeStructui
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json.

// @Param data body Employeemanagementui.EmployeeStructui true "更新EmployeeStructui"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /employeeStructui/updateEmployeeStructui [put]
func (employeeStructuiApi *EmployeeStructuiApi) UpdateEmployeeStructui(c *gin.Context) {
	var employeeStructui Employeemanagementui.EmployeeStructui
	_ = c.ShouldBindJSON(&employeeStructui)
	if err := employeeStructuiService.UpdateEmployeeStructui(employeeStructui); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEmployeeStructui 用id查询EmployeeStructui
// @Tags EmployeeStructui
// @Summary 用id查询EmployeeStructui
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Employeemanagementui.EmployeeStructui true "用id查询EmployeeStructui"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /employeeStructui/findEmployeeStructui [get]
func (employeeStructuiApi *EmployeeStructuiApi) FindEmployeeStructui(c *gin.Context) {
	var employeeStructui Employeemanagementui.EmployeeStructui
	_ = c.ShouldBindQuery(&employeeStructui)
	if err, reemployeeStructui := employeeStructuiService.GetEmployeeStructui(employeeStructui.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reemployeeStructui": reemployeeStructui}, c)
	}
}

// GetEmployeeStructuiList 分页获取EmployeeStructui列表
// @Tags EmployeeStructui
// @Summary 分页获取EmployeeStructui列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query EmployeemanagementuiReq.EmployeeStructuiSearch true "分页获取EmployeeStructui列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /employeeStructui/getEmployeeStructuiList [get]
func (employeeStructuiApi *EmployeeStructuiApi) GetEmployeeStructuiList(c *gin.Context) {
	var pageInfo EmployeemanagementuiReq.EmployeeStructuiSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := employeeStructuiService.GetEmployeeStructuiInfoList(pageInfo); err != nil {
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
