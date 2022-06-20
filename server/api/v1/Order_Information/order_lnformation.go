package Order_Information

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/Order_Information"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    Order_InformationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Order_Information/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type OrderLnformationApi struct {
}

var orderLnformationService = service.ServiceGroupApp.Order_InformationServiceGroup.OrderLnformationService


// CreateOrderLnformation 创建OrderLnformation
// @Tags OrderLnformation
// @Summary 创建OrderLnformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Order_Information.OrderLnformation true "创建OrderLnformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderLnformation/createOrderLnformation [post]
func (orderLnformationApi *OrderLnformationApi) CreateOrderLnformation(c *gin.Context) {
	var orderLnformation Order_Information.OrderLnformation
	_ = c.ShouldBindJSON(&orderLnformation)
	if err := orderLnformationService.CreateOrderLnformation(orderLnformation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrderLnformation 删除OrderLnformation
// @Tags OrderLnformation
// @Summary 删除OrderLnformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Order_Information.OrderLnformation true "删除OrderLnformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orderLnformation/deleteOrderLnformation [delete]
func (orderLnformationApi *OrderLnformationApi) DeleteOrderLnformation(c *gin.Context) {
	var orderLnformation Order_Information.OrderLnformation
	_ = c.ShouldBindJSON(&orderLnformation)
	if err := orderLnformationService.DeleteOrderLnformation(orderLnformation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderLnformationByIds 批量删除OrderLnformation
// @Tags OrderLnformation
// @Summary 批量删除OrderLnformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OrderLnformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /orderLnformation/deleteOrderLnformationByIds [delete]
func (orderLnformationApi *OrderLnformationApi) DeleteOrderLnformationByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := orderLnformationService.DeleteOrderLnformationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrderLnformation 更新OrderLnformation
// @Tags OrderLnformation
// @Summary 更新OrderLnformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Order_Information.OrderLnformation true "更新OrderLnformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orderLnformation/updateOrderLnformation [put]
func (orderLnformationApi *OrderLnformationApi) UpdateOrderLnformation(c *gin.Context) {
	var orderLnformation Order_Information.OrderLnformation
	_ = c.ShouldBindJSON(&orderLnformation)
	if err := orderLnformationService.UpdateOrderLnformation(orderLnformation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrderLnformation 用id查询OrderLnformation
// @Tags OrderLnformation
// @Summary 用id查询OrderLnformation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Order_Information.OrderLnformation true "用id查询OrderLnformation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orderLnformation/findOrderLnformation [get]
func (orderLnformationApi *OrderLnformationApi) FindOrderLnformation(c *gin.Context) {
	var orderLnformation Order_Information.OrderLnformation
	_ = c.ShouldBindQuery(&orderLnformation)
	if err, reorderLnformation := orderLnformationService.GetOrderLnformation(orderLnformation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorderLnformation": reorderLnformation}, c)
	}
}

// GetOrderLnformationList 分页获取OrderLnformation列表
// @Tags OrderLnformation
// @Summary 分页获取OrderLnformation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Order_InformationReq.OrderLnformationSearch true "分页获取OrderLnformation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orderLnformation/getOrderLnformationList [get]
func (orderLnformationApi *OrderLnformationApi) GetOrderLnformationList(c *gin.Context) {
	var pageInfo Order_InformationReq.OrderLnformationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := orderLnformationService.GetOrderLnformationInfoList(pageInfo); err != nil {
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
