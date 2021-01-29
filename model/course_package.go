package model

import "time"

type CoursePackage struct {
	ID        uint      `gorm:"primarykey;type:int(10);not null;comment:主键编号"`
	CourseId  int       `json:"courseId" gorm:"type:int(10);not null;index:course_id;comment:课程编号"`
	PackageId int       `json:"packageId" gorm:"type:int(10);not null;index:package_id;comment:套餐编号"`
	CreatedAt time.Time `gorm:"comment:创建时间;not null;autoCreateTime"`
}
