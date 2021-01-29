package model

import "edushiwei-service/global"

type ChapterUser struct {
	global.COURSE_MODEL
	CourseId  int `json:"courseId" gorm:"type:int(10);not null;index:course_id;comment:课程编号"`
	ChapterId int `json:"chapterId" gorm:"type:int(10);not null;index:chapter_id;comment:章节编号"`
	UserId    int `json:"userId" gorm:"type:int(10);not null;comment:用户编号"`
	PlanId    int `json:"planId" gorm:"type:int(10);not null;comment:计划编号"`
	Duration  int `json:"duration" gorm:"type:int(10);not null;comment:学习时长"`
	Position  int `json:"position" gorm:"type:int(10);not null;comment:播放位置"`
	Progress  int `json:"progress" gorm:"type:int(10);not null;comment:学习进度"`
	Consumed  int `json:"consumed" gorm:"type:int(10);not null;comment:消费标识"`
}
