package model

import "edushiwei-service/global"

type Content struct {
	global.COURSE_MODEL
	UserId     int    `json:"userId" gorm:"type:int(10);not null;index:user_id;comment:用户编号"`
	UnionId    string `json:"unionId" gorm:"type:varchar(64);not null;index:union_provider;comment:union_id"`
	OpenId     string `json:"openId" gorm:"type:varchar(64);index:open_provider;not null;comment:开放ID"`
	OpenName   string `json:"openName" gorm:"type:varchar(30);not null;comment:开放名称"`
	OpenAvatar string `json:"openAvatar" gorm:"type:varchar(150);not null;comment:开放头像"`
	Provider   int    `json:"provider" gorm:"type:int(10);index:open_provider;index:union_provider;not null;comment:提供方"`
}
