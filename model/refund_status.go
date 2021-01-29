package model

import "time"

type RefoundStatus struct {
	ID        uint      `gorm:"primarykey;type:int(10);not null;comment:主键编号"`
	RefundId  int       `json:"refundId" gorm:"type:int(10);not null;index:refund_id;comment:订单编号"`
	Status    int       `json:"status" gorm:"type:int(10);not null;comment:订单状态"`
	CreatedAt time.Time `gorm:"comment:创建时间;not null;autoCreateTime"`
}
