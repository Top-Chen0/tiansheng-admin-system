import service from '@/utils/request'

// @Tags EmployeeModification
// @Summary 创建EmployeeModification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmployeeModification true "创建EmployeeModification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /employeeModification/createEmployeeModification [post]
export const createEmployeeModification = (data) => {
  return service({
    url: '/employeeModification/createEmployeeModification',
    method: 'post',
    data
  })
}

// @Tags EmployeeModification
// @Summary 删除EmployeeModification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmployeeModification true "删除EmployeeModification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /employeeModification/deleteEmployeeModification [delete]
export const deleteEmployeeModification = (data) => {
  return service({
    url: '/employeeModification/deleteEmployeeModification',
    method: 'delete',
    data
  })
}

// @Tags EmployeeModification
// @Summary 删除EmployeeModification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EmployeeModification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /employeeModification/deleteEmployeeModification [delete]
export const deleteEmployeeModificationByIds = (data) => {
  return service({
    url: '/employeeModification/deleteEmployeeModificationByIds',
    method: 'delete',
    data
  })
}

// @Tags EmployeeModification
// @Summary 更新EmployeeModification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmployeeModification true "更新EmployeeModification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /employeeModification/updateEmployeeModification [put]
export const updateEmployeeModification = (data) => {
  return service({
    url: '/employeeModification/updateEmployeeModification',
    method: 'put',
    data
  })
}

// @Tags EmployeeModification
// @Summary 用id查询EmployeeModification
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EmployeeModification true "用id查询EmployeeModification"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /employeeModification/findEmployeeModification [get]
export const findEmployeeModification = (params) => {
  return service({
    url: '/employeeModification/findEmployeeModification',
    method: 'get',
    params
  })
}

// @Tags EmployeeModification
// @Summary 分页获取EmployeeModification列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EmployeeModification列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /employeeModification/getEmployeeModificationList [get]
export const getEmployeeModificationList = (params) => {
  return service({
    url: '/employeeModification/getEmployeeModificationList',
    method: 'get',
    params
  })
}
