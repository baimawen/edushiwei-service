package model

import (
	"edushiwei-service/global"

	"github.com/shopspring/decimal"
)

type Trade struct {
	global.COURSE_MODEL
	OwnerId   int             `json:"ownerId" gorm:"type:int(10);not null;index:owner_id;comment:用户编号"`
	OrderId   int             `json:"orderId" gorm:"type:int(10);not null;index:order_id;comment:订单编号"`
	Sn        string          `json:"sn" gorm:"type:varchar(32);not null;index:sn;comment:交易序号"`
	Subject   string          `json:"subject" gorm:"type:varchar(100);not null;comment:交易主题"`
	Amount    decimal.Decimal `json:"amount" gorm:"type:decimal(10,2);not null;comment:交易金额"`
	Channel   int             `json:"channel" gorm:"type:int(10);not null;comment:平台类型"`
	ChannelSn string          `json:"channelSn" gorm:"type:varchar(64);not null;comment:平台序号"`
	Status    int             `json:"status" gorm:"type:int(10);not null;comment:状态类型"`
}
