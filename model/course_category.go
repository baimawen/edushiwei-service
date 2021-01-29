package model

import "time"

type CourseCategory struct {
	ID         uint      `gorm:"primarykey;type:int(10);not null;comment:主键编号"`
	CourseId   int       `json:"courseId" gorm:"type:int(10);not null;index:course_id;comment:课程编号"`
	CategoryId int       `json:"categoryId" gorm:"type:int(10);not null;index:category_id;comment:分类编号"`
	CreatedAt  time.Time `gorm:"comment:创建时间;not null;autoCreateTime"`
}
