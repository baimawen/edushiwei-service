package model

import "edushiwei-service/global"

type CourseUser struct {
	global.COURSE_MODEL
	CourseId   int `json:"courseId" gorm:"type:int(10);not null;index:course_id;index:course_user;comment:课程编号"`
	UserId     int `json:"userId" gorm:"type:int(10);not null;index:user_id;index:course_user;comment:用户编号"`
	PlanId     int `json:"planId" gorm:"type:int(10);not null;comment:计划编号"`
	RoleType   int `json:"roleType" gorm:"type:int(10);not null;comment:角色类型"`
	SourceType int `json:"sourceType" gorm:"type:int(10);not null;comment:来源类型"`
	Duration   int `json:"duration" gorm:"type:int(10);not null;comment:学习时长"`
	Progress   int `json:"progress" gorm:"type:int(10);not null;comment:学习进度"`
	Reviewed   int `json:"reviewed" gorm:"type:int(10);not null;comment:评价标识"`
	ExpiryTime int `json:"expiryTime" gorm:"type:int(10);not null;comment:过期时间"`
}
