package model

import (
	"edushiwei-service/global"

	"github.com/shopspring/decimal"
)

type Order struct {
	global.COURSE_MODEL
	Sn         string          `json:"sn" gorm:"type:varchar(32);not null;comment:订单编号"`
	Subject    string          `json:"subject" gorm:"type:varchar(100);not null;index:sn;comment:订单标题"`
	Amount     decimal.Decimal `json:"amount" gorm:"type:decimal(10,2);not null;comment:订单金额"`
	OwnerId    int             `json:"ownerId" gorm:"type:int(10);not null;index:owner_id;comment:用户编号"`
	ItemId     int             `json:"itemId" gorm:"type:int(10);not null;index:item;comment:条目编号"`
	ItemType   int             `json:"itemType" gorm:"type:int(10);not null;index:item;comment:"条目类型"`
	ItemInfo   string          `json:"itemInfo" gorm:"type:varchar(3000);not null;comment:条目内容"`
	ClientType int             `json:"clientType" gorm:"type:int(0);not null;comment:终端类型"`
	ClientIP   string          `json:"clientIP" gorm:"type:varchar(64);not null;comment:终端IP"`
	status     int             `json:"status" gorm:"type:int(10);not null;comment:状态标识"`
}
