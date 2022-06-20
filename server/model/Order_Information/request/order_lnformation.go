package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Order_Information"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type OrderLnformationSearch struct{
    Order_Information.OrderLnformation
    request.PageInfo
}
