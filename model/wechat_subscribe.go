package model

import "edushiwei-service/global"

type WechatSubscribe struct {
	global.COURSE_MODEL
	UserId int    `json:"userId" gorm:"type:int(10);not null;comment:用户编号"`
	OpenId string `json:"openId" gorm:"type:varchar(64);not null;comment:开放ID"`
}
