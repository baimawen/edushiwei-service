package model

import "edushiwei-service/global"

type Slide struct {
	global.COURSE_MODEL
	Title     string `json:"title" gorm:"type:varchar(100);not null;comment:标题"`
	Cover     string `json:"cover" gorm:"type:varchar(100);not null;comment:封面"`
	Summary   string `json:"summary" gorm:"type:varchar(255);not null;comment:简介"`
	Content   string `json:"content" gorm:"type:varchar(255);not null;comment:内容"`
	Platform  int    `json:"platform" gorm:"type:int(10);not null;comment:平台类型"`
	Target    int    `json:"target" gorm:"type:int(10);not null;comment:目标类型"`
	Priority  int    `json:"priority" gorm:"type:int(10);not null;comment:优先级"`
	Published int    `json:"published" gorm:"type:int(10);not null;comment:发布状态"`
}
