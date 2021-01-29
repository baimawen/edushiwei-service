package model

import "edushiwei-service/global"

type CourseFavorite struct {
	global.COURSE_MODEL
	CourseId int `json:"courseId" gorm:"type:int(10);not null;index:course_id;comment:课程编号"`
	UserId   int `json:"userId" gorm:"type:int(10);not null;index:user_id;comment:用户编号"`
}
