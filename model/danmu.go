package model

import "edushiwei-service/global"

type Danmu struct {
	global.COURSE_MODEL
	CourseId  int    `json:"courseId" gorm:"type:int(10);not null;comment:课程编号"`
	ChapterId int    `json:"chapterId" gorm:"type:int(10);not null;index:chapter_id;comment:章节编号"`
	OwnerId   int    `json:"ownerId" gorm:"type:int(10);not null;index:owver_id;comment:用户编号"`
	Time      int    `json:"time" gorm:"type:int(10);not null;comment:时间轴"`
	Text      string `json:"text" gorm:"type:varchar(255);not null;comment:内容"`
	Color     string `json:"color" gorm:"type:varchar(30);not null;comment:颜色"`
	Size      int    `json:"size" gorm:"type:int(10);not null;comment:字号"`
	Position  int    `json:"position" gorm:"type:int(10);not null;comment:位置"`
	Published int    `json:"published" gorm:"type:int(10);not null;comment:发布标识"`
}
