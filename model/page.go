package model

import "edushiwei-service/global"

type Page struct {
	global.COURSE_MODEL
	Title     string `json:"title" gorm:"type:varchar(100);not null;comment:标题"`
	Content   string `json:"content" gorm:"type:text;not null;comment:内容"`
	Published int    `json:"published" gorm:"type:int(10);not null;comment:发布标识"`
}
