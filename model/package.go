package model

import (
	"edushiwei-service/global"

	"github.com/shopspring/decimal"
)

type Package struct {
	global.COURSE_MODEL
	Title       string          `json:"title" gorm:"type:varchar(100);not null;comment:标题"`
	Summary     string          `json:"summary" gorm:"type:varchar(255);not null;comment:简介"`
	MarketPrice decimal.Decimal `json:"marketPrice" gorm:"type:decimal(10,2);not null;comment:市场价格"`
	VipPrice    decimal.Decimal `json:"vipPrice" gorm:"type:decimal(10,2);not null;comment:会员价格"`
	CourseCount int             `json:"courseCount" gorm:"type:int(10);not null;comment:课程数量"`
	Published   int             `json:"published" gorm:"type:int(10);not null;comment:发布标识"`
}
