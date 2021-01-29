package model

import "edushiwei-service/global"

type Category struct {
	global.COURSE_MODEL
	ParentId  int    `json:"parentId" gorm:"type:int(10);not null;comment:父级编号"`
	Level     int    `json:"level" gorm:"type:int(10);not null;comment:层级"`
	Type      int    `json:"type" gorm:"type:int(10);not null;comment:类型"`
	Name      string `json:"name" gorm:"type:varchar(30);not null;comment:名称"`
	Path      string `json:"path" gorm:"type:varchar(30);not null;comment:路径"`
	Priority  int    `json:"priority" gorm:"type:int(10);not null;comment:优先级"`
	Published int    `json:"published" gorm:"type:int(10);not null;comment:发布标识"`
	ChildCout int    `json:"childCount" gorm:"type:int(10);not null;comment:节点数"`
}
