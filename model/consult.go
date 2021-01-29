package model

import "edushiwei-service/global"

type Consult struct {
	global.COURSE_MODEL
	CourseId  int    `json:"courseId" gorm:"type:int(10);not null;index:course_id;comment:课程编号"`
	ChapterId int    `json:"chapterId" gorm:"type:int(10);not null;index:chapter_id;comment:章节编号"`
	OwnerId   int    `json:"ownerId" gorm:"type:int(10);not null;index:owner_id;comment:用户编号"`
	ReplierId int    `json:"replierId" gorm:"type:int(11);not null;comment:回复者编号"`
	Question  string `json:"question" gorm:"type:varchar(1500);not null;comment:问题"`
	Answer    string `json:"answer" gorm:"type:varchar(1500);not null;comment:答案"`
	Rating    int    `json:"rating" gorm:"type:int(10);not null;comment:评分"`
	Priority  int    `json:"priority" gorm:"type:int(10);not null;comment:优先级"`
	Private   int    `json:"private" gorm:"type:int(10);not null;comment:私密标识"`
	Published int    `json:"published" gorm:"type:int(10);not null;comment:发布标识"`
	LikeCount int    `json:"likeCount" gorm:"type:int(10);not null;comment:点赞数"`
	ReplyTime int    `json:"repluTime" gorm:"type:int(10);not null;comment:回复时间"`
}
