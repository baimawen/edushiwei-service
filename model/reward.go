package model

import (
	"edushiwei-service/global"

	"github.com/shopspring/decimal"
)

type Reward struct {
	global.COURSE_MODEL
	Title string          `json:"title" gorm:"type:varchar(100);not null;comment:标题"`
	Price decimal.Decimal `json:"price" gorm:"type:decimal(10,2);not null;comment:价格"`
}
