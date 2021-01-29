package model

import "edushiwei-service/global"

type Nav struct {
	global.COURSE_MODEL
	ParentId   int    `json:"parentId" gorm:"type:int(10);not null;default:0;comment:父级编号"`
	Level      int    `json:"level" gorm:"type:int(10);not null;comment:层级"`
	Name       string `json:"name" gorm:"type:varchar(30);not null;comment:名称"`
	Path       string `json:"path" gorm:"type:varchar(30);not null;comment:路径"`
	Target     string `json:"target" gorm:"type:varchar(30);not null;comment:打开方式"`
	Url        string `json:"url" gorm:"type:varchar(100);not null;comment:链接地址"`
	Position   int    `json:"position" gorm:"type:int(10);not null;comment:位置"`
	Priority   int    `json:"priority" gorm:"type:int(10);not null;comment:优先级"`
	Published  int    `json:"published" gorm:"type:int(10);not null;comment:发布标识"`
	ChildCount int    `json:"childCount" gorm:"type:int(10);not null;comment:子类数量"`
}
