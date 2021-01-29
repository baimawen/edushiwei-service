package model

import "edushiwei-service/global"

type ImUser struct {
	global.COURSE_MODEL
	Name        string `json:"name" gorm:"type:varchar(30);not null;comment:名称"`
	Avatar      string `json:"avatar" gorm:"type:varchar(100);not null;comment:头像"`
	Sign        string `json:"sign" gorm:"type:varchar(30);not null;comment:签名"`
	Skin        string `json:"skin" gorm:"type:varchar(100);not null;comment:皮肤"`
	Status      string `json:"status" gorm:"type:varchar(15);not null;comment:在线状态"`
	FriendCount int    `json:"friendCount" gorm:"type:int(10);not null;comment:好友数"`
	GroupCount  int    `json:"groupCount" gorm:"type:int(10);not null;comment:群组数"`
}
