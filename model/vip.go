package model

import (
	"edushiwei-service/global"

	"github.com/shopspring/decimal"
)

type Vip struct {
	global.COURSE_MODEL
	Title  string          `json:"title" gorm:"type:varchar(30);not null;comment:标题"`
	Expiry int             `json:"expiry" gorm:"type:int(10);not null;comment:有效期"`
	Price  decimal.Decimal `json:"price" gorm:"type:decimal(10,2);not null;comment:价格"`
}
