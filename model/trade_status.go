package model

import "time"

type TradeStatus struct {
	ID        uint      `gorm:"primarykey;type:int(10);not null;index:trade_id;comment:主键编号"`
	TradeId   int       `json:"tradeId" gorm:"type:int(10);not null;comment:订单编号"`
	Status    int       `json:"status" gorm:"type:int(10);not null;comment:订单状态"`
	CreatedAt time.Time `gorm:"comment:创建时间;not null;autoCreateTime"`
}
