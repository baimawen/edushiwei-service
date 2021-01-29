package model

import "time"

type CourseRelated struct {
	ID        uint      `gorm:"primarykey;type:int(10);not null;comment:主键编号"`
	CourseId  int       `json:"courseId" gorm:"type:int(10);not null;index:course_id;comment:课程编号"`
	RelatedId int       `json:"relatedId" gorm:"type:int(10);not null;index:related_id;comment:相关编号"`
	CreatedAt time.Time `gorm:"comment:创建时间;not null;autoCreateTime"`
}
