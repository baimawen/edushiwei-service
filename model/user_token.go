package model

import "edushiwei-service/global"

type UserToken struct {
	global.COURSE_MODEL
	UserId     int    `json:"userId" gorm:"type:int(10);not null;index:user_id;comment:用户编号"`
	Token      string `json:"token" gorm:"type:varchar(64);not null;comment:身份令牌"`
	ClientType int    `json:"clientType" gorm:"type:int(10);not null;comment:终端类型"`
	ClientIP   string `json:"clientIP" gorm:"type:varchar(64);not null;comment:终端IP"`
}
