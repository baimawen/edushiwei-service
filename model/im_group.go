package model

import "edushiwei-service/global"

type ImGroup struct {
	global.COURSE_MODEL
	OwnerId   int    `json:"ownerId" gorm:"type:int(10);not null;comment:群主编号"`
	CourseId  int    `json:"courseId" gorm:"type:int(10);not null;comment:课程编号"`
	Type      int    `json:"type" gorm:"type:int(11);not null;comment:类型"`
	Name      string `json:"name" gorm:"type:varchar(100);not null;comment:名称"`
	Avatar    string `json:"avatar" gorm:"type:varchar(100);not null;comment:头像"`
	About     string `json:"about" gorm:"type:varchar(255);not null;comment:简介"`
	Published int    `json:"published" gorm:"type:int(10);not null;comment:发布标识"`
	UserCount int    `json:"userCount" gorm:"type:int(10);not null;comment:成员数"`
	MsgCout   int    `json:"msgCount" gorm:"type:int(10);not null;comment:消息数"`
}
