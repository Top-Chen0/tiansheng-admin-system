import service from '@/utils/request'

// @Tags OrderLnformation
// @Summary 创建OrderLnformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderLnformation true "创建OrderLnformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderLnformation/createOrderLnformation [post]
export const createOrderLnformation = (data) => {
  return service({
    url: '/orderLnformation/createOrderLnformation',
    method: 'post',
    data
  })
}

// @Tags OrderLnformation
// @Summary 删除OrderLnformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderLnformation true "删除OrderLnformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderLnformation/deleteOrderLnformation [delete]
export const deleteOrderLnformation = (data) => {
  return service({
    url: '/orderLnformation/deleteOrderLnformation',
    method: 'delete',
    data
  })
}

// @Tags OrderLnformation
// @Summary 删除OrderLnformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderLnformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderLnformation/deleteOrderLnformation [delete]
export const deleteOrderLnformationByIds = (data) => {
  return service({
    url: '/orderLnformation/deleteOrderLnformationByIds',
    method: 'delete',
    data
  })
}

// @Tags OrderLnformation
// @Summary 更新OrderLnformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OrderLnformation true "更新OrderLnformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderLnformation/updateOrderLnformation [put]
export const updateOrderLnformation = (data) => {
  return service({
    url: '/orderLnformation/updateOrderLnformation',
    method: 'put',
    data
  })
}

// @Tags OrderLnformation
// @Summary 用id查询OrderLnformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OrderLnformation true "用id查询OrderLnformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderLnformation/findOrderLnformation [get]
export const findOrderLnformation = (params) => {
  return service({
    url: '/orderLnformation/findOrderLnformation',
    method: 'get',
    params
  })
}

// @Tags OrderLnformation
// @Summary 分页获取OrderLnformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OrderLnformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderLnformation/getOrderLnformationList [get]
export const getOrderLnformationList = (params) => {
  return service({
    url: '/orderLnformation/getOrderLnformationList',
    method: 'get',
    params
  })
}
