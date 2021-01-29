package model

import "edushiwei-service/global"

type ImFriendUser struct {
	global.COURSE_MODEL
	UserId   int `json:"userId" gorm:"type:int(10);not null;index:user_friend;comment:用户编号"`
	FriendId int `json:"friendId" gorm:"type:int(10);not null;index:user_friend;comment:目标编号"`
	GroupId  int `json:"groupId" gorm:"type:int(10);not null;comment:分组编号"`
	MsgCout  int `json:"msgCout" gorm:"type:int(10);not null;comment:消息数"`
}
