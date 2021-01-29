package model

import (
	"edushiwei-service/global"

	"github.com/shopspring/decimal"
)

type Refund struct {
	global.COURSE_MODEL
	OwnerId    int             `json:"ownerId" gorm:"type:int(10);not null;index:owner_id;comment:用户编号"`
	OrderId    int             `json:"orderId" gorm:"type:int(10);not null;index:order_id;comment:订单编号"`
	TradeId    int             `json:"tradeId" gorm:"type:int(10);not null;index:trade_id;comment:交易编号"`
	Sn         string          `json:"sn" gorm:"type:varchar(32);not null;index:sn;comment:退款序号"`
	subject    string          `json:"subject" gorm:"type:varchar(100);not null;comment:退款主题"`
	Amount     decimal.Decimal `json:"amount" gorm:"type:decimal(10,2);not null;comment:退款金额"`
	Status     int             `json:"status" gorm:"type:int(10);not null;comment:状态类型"`
	ApplyNote  string          `json:"applyNote" gorm:"type:varchar(255);not null;comment:申请备注"`
	ReviewNote string          `json:"reviewNote" gorm:"type:varchar(255);not null;comment:审核备注"`
}
