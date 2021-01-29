package model

import "edushiwei-service/global"

type ChapterVod struct {
	global.COURSE_MODEL
	CourseId      int    `json:"courseId" gorm:"type:int(10);not null;index:course_id;comment:课程编号"`
	ChapterId     int    `json:"chapterId" gorm:"type:int(10);not null;index:chapter_id;comment:章节编号"`
	FileId        string `json:"fileId" gorm:"type:varchar(32);not null;inex:file_id;comment:文件编号"`
	FileTranscode string `json:"fileTranscode" gorm:"type:varchar(1500);not null;comment:文件属性"`
}
