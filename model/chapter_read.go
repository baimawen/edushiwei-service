package model

import "edushiwei-service/global"

type ChapterRead struct {
	global.COURSE_MODEL
	CourseId  int    `json:"courseId" gorm:"type:int(10);not null;index:course_id;comment:课程编号"`
	ChapterId int    `json:"chapterId" gorm:"type:int(10);not null;index:chapter_id;comment:章节编号"`
	Content   string `json:"content" gorm:"type:text;not null;comment:内容"`
}
