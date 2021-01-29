package model

import "edushiwei-service/global"

type ChapterLive struct {
	global.COURSE_MODEL
	CourseId  int `json:"courseId" gorm:"type:int(10);not null;index:course_id;comment:课程编号"`
	ChapterId int `json:"chapterId" gorm:"type:int(10);not null;index:chapter_id;comment:章节编号"`
	StartTime int `json:"startTime" gorm:"type:int(10);not null;comment:开始时间"`
	EndTime   int `json:"endTime" gorm:"type:int(10);not null;comment:结束时间"`
	UserLimit int `json:"userLimit" gorm:"type:int(10);not null;comment:用户限额"`
	Status    int `json:"status" gorm:"type:int(10);not null;comment:状态标识"`
}
