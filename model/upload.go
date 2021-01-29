package model

import "edushiwei-service/global"

type Upload struct {
	global.COURSE_MODEL
	Type int    `json:"type" gorm:"type:int(10);not null;comment:条目类型`
	Name string `json:"name" gorm:"type:varchar(100);not null;comment:文件名"`
	Path string `json:"path" gorm:"type.varchar(100);not null;comment:路径"`
	Mime string `json:"mime" gorm:"type:varchar(100);not null;comment:mime"`
	Md5  string `json:"md5" gorm:"type:varchar(32);not null;index:md5;comment:md5"`
	Size int    `json:"size" gorm:"type:int(10); not null;comment:大小"`
}
