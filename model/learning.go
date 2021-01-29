package model

import "edushiwei-service/global"

type Learning struct {
	global.COURSE_MODEL
	RequestId  string `json:"requestId" gorm:"type:varchar(64);not null;index:request_id;comment:请求编号"`
	CourseId   int    `json:"courseId" gorm:"type:int(10);not null;comment:课程编号"`
	ChapterId  int    `json:"chapterId" gorm:"type:int(10);not null;index:chapter_user;comment:章节编号"`
	UserId     int    `json:"userId" gorm:"type:int(10);not null;index:user_id;comment:用户编号"`
	PlanId     int    `json:"planId" gorm:"type:int(10);not null;comment:计划编号"`
	Duration   int    `json:"duration" gorm:"type:int(10);not null;comment:学习时长"`
	Position   int    `json:"position" gorm:"type:int(10);not null;comment:播放位置"`
	ClientType int    `json:"clientType" gorm:"type:int(10);not null;comment:终端类型"`
	ClientIP   string `json:"clientIP" gorm:"type:varchar(64);not null;comment:终端IP"`
	ActiveTime int    `json:"activeTime" gorm:"type:int(10);not null;comment:活跃时间"`
}
