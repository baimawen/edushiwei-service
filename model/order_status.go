package model

import "time"

type OrderStatus struct {
	ID        uint      `gorm:"primarykey;type:int(10);not null;comment:主键编号"`
	OrderId   int       `json:"orderId" gorm:"type:int(10);not null;index:order_id;comment:订单编号"`
	status    int       `json:"status" gorm:"type:int(10);not null;comment:订单状态"`
	CreatedAt time.Time `gorm:"comment:创建时间;not null;autoCreateTime"`
}
