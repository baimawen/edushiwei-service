package model

import "edushiwei-service/global"

type ReviewLike struct {
	global.COURSE_MODEL
	ReviewId int `json:"reviewId" gorm:"type:int(10);not null;index:review_user;comment:评价编号"`
	UserId   int `json:"userId" gorm:"type:int(10);not null;index:review_user;comment:用户编号"`
}
