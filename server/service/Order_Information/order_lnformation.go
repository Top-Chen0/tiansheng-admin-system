package Order_Information

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Order_Information"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    Order_InformationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Order_Information/request"
)

type OrderLnformationService struct {
}

// CreateOrderLnformation 创建OrderLnformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderLnformationService *OrderLnformationService) CreateOrderLnformation(orderLnformation Order_Information.OrderLnformation) (err error) {
	err = global.GVA_DB.Create(&orderLnformation).Error
	return err
}

// DeleteOrderLnformation 删除OrderLnformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderLnformationService *OrderLnformationService)DeleteOrderLnformation(orderLnformation Order_Information.OrderLnformation) (err error) {
	err = global.GVA_DB.Delete(&orderLnformation).Error
	return err
}

// DeleteOrderLnformationByIds 批量删除OrderLnformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderLnformationService *OrderLnformationService)DeleteOrderLnformationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Order_Information.OrderLnformation{},"id in ?",ids.Ids).Error
	return err
}

// UpdateOrderLnformation 更新OrderLnformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderLnformationService *OrderLnformationService)UpdateOrderLnformation(orderLnformation Order_Information.OrderLnformation) (err error) {
	err = global.GVA_DB.Save(&orderLnformation).Error
	return err
}

// GetOrderLnformation 根据id获取OrderLnformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderLnformationService *OrderLnformationService)GetOrderLnformation(id uint) (err error, orderLnformation Order_Information.OrderLnformation) {
	err = global.GVA_DB.Where("id = ?", id).First(&orderLnformation).Error
	return
}

// GetOrderLnformationInfoList 分页获取OrderLnformation记录
// Author [piexlmax](https://github.com/piexlmax)
func (orderLnformationService *OrderLnformationService)GetOrderLnformationInfoList(info Order_InformationReq.OrderLnformationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Order_Information.OrderLnformation{})
    var orderLnformations []Order_Information.OrderLnformation
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&orderLnformations).Error
	return err, orderLnformations, total
}
