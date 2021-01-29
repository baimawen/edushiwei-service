package model

import "edushiwei-service/global"

type Topic struct {
	global.COURSE_MODEL
	Title       string `json:"title" gorm:"type:varchar(100);not null;comment:标题"`
	Keywords    string `json:"keywords" gorm:"type:varchar(100);not null;comment:关键字"`
	Summary     string `json:"summary" gorm:"type:varchar(255);not null;comment:简介"`
	CourseCount int    `json:"courseCount" gorm:"type:int(10);not null;comment:课程数量"`
	Published   int    `json:"published" gorm:"type:int(10);not null;comment:发布标识"`
}
