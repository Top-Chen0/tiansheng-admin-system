package Order_Information

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderLnformationRouter struct {
}

// InitOrderLnformationRouter 初始化 OrderLnformation 路由信息
func (s *OrderLnformationRouter) InitOrderLnformationRouter(Router *gin.RouterGroup) {
	orderLnformationRouter := Router.Group("orderLnformation").Use(middleware.OperationRecord())
	orderLnformationRouterWithoutRecord := Router.Group("orderLnformation")
	var orderLnformationApi = v1.ApiGroupApp.Order_InformationApiGroup.OrderLnformationApi
	{
		orderLnformationRouter.POST("createOrderLnformation", orderLnformationApi.CreateOrderLnformation)   // 新建OrderLnformation
		orderLnformationRouter.DELETE("deleteOrderLnformation", orderLnformationApi.DeleteOrderLnformation) // 删除OrderLnformation
		orderLnformationRouter.DELETE("deleteOrderLnformationByIds", orderLnformationApi.DeleteOrderLnformationByIds) // 批量删除OrderLnformation
		orderLnformationRouter.PUT("updateOrderLnformation", orderLnformationApi.UpdateOrderLnformation)    // 更新OrderLnformation
	}
	{
		orderLnformationRouterWithoutRecord.GET("findOrderLnformation", orderLnformationApi.FindOrderLnformation)        // 根据ID获取OrderLnformation
		orderLnformationRouterWithoutRecord.GET("getOrderLnformationList", orderLnformationApi.GetOrderLnformationList)  // 获取OrderLnformation列表
	}
}
