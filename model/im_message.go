package model

import "edushiwei-service/global"

type ImMessage struct {
	global.COURSE_MODEL
	ChatId       string `json:"chatId" gorm:"type:varchar(30);not null;index:chat_id;comment:对话编号"`
	SenderId     int    `json:"sendId" gorm:"type:int(10);not null;comment:发送方"`
	ReceiverId   int    `json:"receiverId" gorm:"type:int(10);not null;index:receiver_id;comment:接收方"`
	ReceiverType int    `json:"receiverType" gorm:"type:int(11);not null;index:receiver_type;comment:接收方类型"`
	Content      string `json:"content" gorm:"type:varchar(3000);not null;comment:内容"`
	Viewed       int    `json:"viewed" gorm:"type:int(10);not null;comment:已读标识"`
}
