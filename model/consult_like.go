package model

import "edushiwei-service/global"

type ConsultLike struct {
	global.COURSE_MODEL
	ConsultId int `json:"consultId" gorm:"type:int(10);not null;index:consult_user;comment:咨询编号"`
	UserId    int `json:"userId" gorm:"type:int(10);not null;index:consult_user;comment:用户编号"`
}
