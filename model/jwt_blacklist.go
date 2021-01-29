package model

import "edushiwei-service/global"

type JwtBlacklist struct {
	global.COURSE_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
