package model

import "edushiwei-service/global"

type Task struct {
	global.COURSE_MODEL
	ItemId   int `json:"itemId" gorm:"type:int(10);not null;comment:条目编号"`
	ItemType int `json:"itemType" gorm:"type:int(10);not null;comment:条目类型"`
	ItemInfo int `json:"itemInfo" gorm:"type:varchar(3000);not null;comment:条目内容"`
	Status   int `json:"status" gorm:"type:int(10);not null;comment:状态"`
	Priority int `json:"priority" gorm:"type:int(10);not null;comment:优先级"`
	TryCount int `json:"tryCount" gorm:"type:int(10);not null;comment:重试数"`
}
