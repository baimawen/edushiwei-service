package model

import "edushiwei-service/global"

type Accout struct {
	global.COURSE_MODEL
	Email    string `json:"email" gorm:"type:varchar(30);not null;index:email;comment:邮箱"`
	Phone    string `json:"phone" gorm:"type:varchar(30);not null;index:phone;comment:手机"`
	Password string `json:"password" gorm:"type:varchar(32);not null;comment:密码"`
	Slat     string `json:"slat" gorm:"type:varchar(32);not null;comment:密盐"`
}
