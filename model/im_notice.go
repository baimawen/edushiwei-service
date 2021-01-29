package model

import "edushiwei-service/global"

type ImNotice struct {
	global.COURSE_MODEL
	SenderId   int    `json:"senderId" gorm:"type:int(10);not null;index:sender_id;comment:发送方"`
	ReceiverId int    `json:"receiverId" gorm:"type:int(10);not null;index:receiver_id;comment:接收方"`
	ItemType   int    `json:"itemType" gorm:"type:int(10);not null;comment:条目类型"`
	ItemInfo   string `json:"itemInfo" gorm:"type:varchar(1500);not null;comment:条目内容"`
	Viewed     int    `json:"viewed" gorm:"type:int(10);not null;comment:已读标识"`
}
