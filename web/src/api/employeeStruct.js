import service from '@/utils/request'

// @Tags EmployeeStruct
// @Summary 创建EmployeeStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmployeeStruct true "创建EmployeeStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /employeeStruct/createEmployeeStruct [post]
export const createEmployeeStruct = (data) => {
  return service({
    url: '/employeeStruct/createEmployeeStruct',
    method: 'post',
    data
  })
}

// @Tags EmployeeStruct
// @Summary 删除EmployeeStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmployeeStruct true "删除EmployeeStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /employeeStruct/deleteEmployeeStruct [delete]
export const deleteEmployeeStruct = (data) => {
  return service({
    url: '/employeeStruct/deleteEmployeeStruct',
    method: 'delete',
    data
  })
}

// @Tags EmployeeStruct
// @Summary 删除EmployeeStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EmployeeStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /employeeStruct/deleteEmployeeStruct [delete]
export const deleteEmployeeStructByIds = (data) => {
  return service({
    url: '/employeeStruct/deleteEmployeeStructByIds',
    method: 'delete',
    data
  })
}

// @Tags EmployeeStruct
// @Summary 更新EmployeeStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EmployeeStruct true "更新EmployeeStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /employeeStruct/updateEmployeeStruct [put]
export const updateEmployeeStruct = (data) => {
  return service({
    url: '/employeeStruct/updateEmployeeStruct',
    method: 'put',
    data
  })
}

// @Tags EmployeeStruct
// @Summary 用id查询EmployeeStruct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EmployeeStruct true "用id查询EmployeeStruct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /employeeStruct/findEmployeeStruct [get]
export const findEmployeeStruct = (params) => {
  return service({
    url: '/employeeStruct/findEmployeeStruct',
    method: 'get',
    params
  })
}

// @Tags EmployeeStruct
// @Summary 分页获取EmployeeStruct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EmployeeStruct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /employeeStruct/getEmployeeStructList [get]
export const getEmployeeStructList = (params) => {
  return service({
    url: '/employeeStruct/getEmployeeStructList',
    method: 'get',
    params
  })
}
