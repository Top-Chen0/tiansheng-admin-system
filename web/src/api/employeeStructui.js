import service from '@/utils/request'

// @Tags EmployeeStructui
// @Summary 创建EmployeeStructui
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmployeeStructui true "创建EmployeeStructui"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /employeeStructui/createEmployeeStructui [post]
export const createEmployeeStructui = (data) => {
  return service({
    url: '/employeeStructui/createEmployeeStructui',
    method: 'post',
    data
  })
}

// @Tags EmployeeStructui
// @Summary 删除EmployeeStructui
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmployeeStructui true "删除EmployeeStructui"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /employeeStructui/deleteEmployeeStructui [delete]
export const deleteEmployeeStructui = (data) => {
  return service({
    url: '/employeeStructui/deleteEmployeeStructui',
    method: 'delete',
    data
  })
}

// @Tags EmployeeStructui
// @Summary 删除EmployeeStructui
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EmployeeStructui"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /employeeStructui/deleteEmployeeStructui [delete]
export const deleteEmployeeStructuiByIds = (data) => {
  return service({
    url: '/employeeStructui/deleteEmployeeStructuiByIds',
    method: 'delete',
    data
  })
}

// @Tags EmployeeStructui
// @Summary 更新EmployeeStructui
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmployeeStructui true "更新EmployeeStructui"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /employeeStructui/updateEmployeeStructui [put]
export const updateEmployeeStructui = (data) => {
  return service({
    url: '/employeeStructui/updateEmployeeStructui',
    method: 'put',
    data
  })
}

// @Tags EmployeeStructui
// @Summary 用id查询EmployeeStructui
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EmployeeStructui true "用id查询EmployeeStructui"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /employeeStructui/findEmployeeStructui [get]
export const findEmployeeStructui = (params) => {
  return service({
    url: '/employeeStructui/findEmployeeStructui',
    method: 'get',
    params
  })
}

// @Tags EmployeeStructui
// @Summary 分页获取EmployeeStructui列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EmployeeStructui列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /employeeStructui/getEmployeeStructuiList [get]
export const getEmployeeStructuiList = (params) => {
  return service({
    url: '/employeeStructui/getEmployeeStructuiList',
    method: 'get',
    params
  })
}
