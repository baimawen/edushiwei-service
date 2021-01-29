package model

import "edushiwei-service/global"

type Role struct {
	global.COURSE_MODEL
	Type      int    `json:"type" gorm:"type:int(10);not null;comment:类型"`
	Name      string `json:"name" gorm:"type:varchar(30);not null;comment:名称"`
	Summary   string `json:"summary" gorm:"type:varchar(255);not null;comment:简介"`
	Routes    string `json:"routes" gorm:"type:text;not null;comment:权限路由"`
	UserCount int    `json:"userCount" gorm:"type:int(10);not null;comment:成员数量"`
}
