package model

import (
	"edushiwei-service/global"
)

type User struct {
	global.COURSE_MODEL
	Name           string `json:"name" gorm:"type:varchar(30);not null;index:name;comment:名称"`
	Avatar         string `json:"avatar" gorm:"type:varchar(100);not null;comment:头像"`
	Title          string `json:"title" gorm:"type:varchar(30);not null;comment:头衔"`
	About          string `json:"about" gorm:"type:varchar(255);not null;comment:简介"`
	Area           string `json:"area" gorm:"type:varchar(30);not null;comment:地区"`
	Gender         int    `json:"gender" gorm:"type:int(10);not null;comment:性别"`
	Vip            int    `json:"vip" gorm:"type:int(10);not null;comment:会员标识"`
	Locked         int    `json:"locked" gorm:"type:int(10);not null;comment:锁定标识"`
	EduRole        int    `json:"eduRole" gorm:"type:int(10);not null;comment:教学角色"`
	AdminRole      int    `json:"adminRole" gorm:"type:int(10);not null;comment:后台角色"`
	CourseCount    int    `json:"courseCount" gorm:"type:int(10);not null;comment:课程数"`
	FavoriteCount  int    `json:"favoriteCount" gorm:"type:int(10);not null;comment:收藏数"`
	VipExpiryTime  int    `json:"vipExpiryTime" gorm:"type:int(10);not null;comment:会员期限"`
	LockExpiryTime int    `json:"lockExpiryTime" gorm:"type:int(10);not null;comment:锁定期限"`
	ActiveTime     int    `json:"activeTime" gorm:"type:int(10);not null;comment:活跃时间"`
}
