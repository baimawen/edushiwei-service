package model

import "edushiwei-service/global"

type Help struct {
	global.COURSE_MODEL
	CategoryId int    `json:"categoryId" gorm:"type:int(10);not null;comment:分类编号"`
	Title      string `json:"title" gorm:"type:varchar(100);not null;comment:标题"`
	Content    string `json:"content" gorm:"type:text;not null;comment:内容"`
	Priority   int    `json:"priority" gorm:"type:int(10);not null;comment:优先级"`
	Published  int    `json:"published" gorm:"type:int(10);not null;comment:发布标识"`
}
