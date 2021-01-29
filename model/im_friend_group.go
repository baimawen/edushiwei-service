package model

import "edushiwei-service/global"

type ImFriendGroup struct {
	global.COURSE_MODEL
	UserId    int    `json:"userId" gorm:"type:int(10);not null;index:user_id;comment:用户编号"`
	Name      string `json:"name" gorm:"type:varchar(100);not null;comment:名称"`
	Priority  int    `json:"priority" gorm:"type:int(10);not null;comment:优先级"`
	UserCount int    `json:"userCount" gorm:"type:int(10);not null;comment:成员数"`
}
