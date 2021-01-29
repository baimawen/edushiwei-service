package model

import "edushiwei-service/global"

type ImGroupUser struct {
	global.COURSE_MODEL
	GroupId  int `json:"groupId" gorm:"type:int(10);not null;index:group_id;index:group_user;comment:群组编号"`
	UserId   int `json:"userId" gorm:"type:int(10);not null;index:user_id;index:group_user;comment:用户编号"`
	Priority int `json:"priority" gorm:"type:int(10);not null;comment:优先级"`
}
