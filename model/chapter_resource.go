package model

import (
	"edushiwei-service/global"
)

type ChapterResource struct {
	global.COURSE_MODEL
	CourseId  int `json:"courseId" gorm:"type:int(10);not null;index:course_id;comment:课程编号"`
	ChapterId int `json:"chapterId" gorm:"type:int(10);not null;index:chapter_id;comment:章节编号"`
	UploadId  int `json:"uploadId" gorm:"type:int(10);not null;comment:上传编号"`
}
