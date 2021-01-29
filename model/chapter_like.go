package model

import "edushiwei-service/global"

type ChapterLike struct {
	global.COURSE_MODEL
	ChapterId int `json:"chapterId" gorm:"type:int(10);not null;index:chapter_user;comment:课程编号"`
	UserId    int `json:"userId" gorm:"type:int(10);not null;index:chapter_user;comment:用户编号"`
}
