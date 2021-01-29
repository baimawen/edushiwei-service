package global

import (
	"time"
)

type COURSE_MODEL struct {
	ID        uint      `gorm:"primarykey;type:int(10);not null;comment:主键编号"`
	CreatedAt time.Time `gorm:"comment:创建时间;not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"comment:更新时间;not null;autoUpdateTime"`
	// DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Deleted int `gorm:"type:int(10);comment:删除标识;not null;default:0" json:"-"`
}
