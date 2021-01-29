package model

import "time"

type CourseRating struct {
	CourseId int     `json:"courseId" gorm:"type:int(10);not null;comment:主键编号"`
	Rating   float64 `json:"rating" gorm:"type:float(10,2);not null;comment:综合评分"`
	Rating1  float64 `json:"rating" gorm:"type:float(10,2);not null;comment:维度1评分"`

	Rating2   float64   `json:"rating" gorm:"type:float(10,2);not null;comment:维度2评分"`
	Rating3   float64   `json:"rating" gorm:"type:float(10,2);not null;comment:维度3评分"`
	CreatedAt time.Time `gorm:"comment:创建时间;not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"comment:更新时间;not null;autoUpdateTime"`
}
