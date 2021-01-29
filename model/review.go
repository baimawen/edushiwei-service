package model

import "edushiwei-service/global"

type Review struct {
	global.COURSE_MODEL
	CourseId int     `json:"courseId" gorm:"type:int(10);not null;index:course_id;comment:课程编号"`
	OwnerId  int     `json:"ownerId" gorm:"type:int(10);not null;index:owner_id;comment:用户编号"`
	Content  string  `json:"content" gorm:"type:varchar(1000);not null;comment:内容"`
	Reply    string  `json:"reply" gorm:"type:varchar(1000);not null;comment:回复"`
	Rating   float64 `json:"rating" gorm:"type:float(10,2);not null;comment:综合评分"`
	Rating1  float64 `json:"rating" gorm:"type:float(10,2);not null;comment:维度1评分"`

	Rating2   float64 `json:"rating" gorm:"type:float(10,2);not null;comment:维度2评分"`
	Rating3   float64 `json:"rating" gorm:"type:float(10,2);not null;comment:维度3评分"`
	Anonymous int     `json:"anonymous" gorm:"type:int(10);not null;comment:匿名标识"`
	Published int     `json:"published" gorm:"type:int(10);not null;comment:发布标识"`
	LikeCount int     `json:"likeCount" gorm:"type:int(10);not null;comment:点赞数"`
}
