package model

import (
	"edushiwei-service/global"

	"github.com/shopspring/decimal"
)

type Course struct {
	global.COURSE_MODEL
	Title          string          `json:"title" gorm:"type:varchar(100);not null;comment:标题"`
	Cover          string          `json:"cover" gorm:"type:varchar(100);not null;comment:封面"`
	Summary        string          `json:"summary" gorm:"type:varchar(255);not null;comment:简介"`
	Keywords       string          `json:"keywords" gorm:"type:varchar(100);not null;comment:关键字"`
	Details        string          `json:"details" gorm:"type:text;not null;comment:详情"`
	CategoryId     int             `json:"categoryId" gorm:"type:int(10);not null;comment:分类编号"`
	TeacherId      int             `json:"teacherId" gorm:"type:int(10);not null;comment:讲师编号"`
	MarketPrice    decimal.Decimal `json:"marketPrice" gorm:"type:decimal(10,2);not null;comment:市场价格"`
	VipPrice       decimal.Decimal `json:"vipPrice" gorm:"type:decimal(10,2);not null;comment:会员价格"`
	StudyExpiry    int             `json:"studyExpiry" gorm:"type:int(10);not null;comment:学习期限"`
	RefundExpiry   int             `json:"refundExpiry" gorm:"type:int(10);not null;comment:退款期限"`
	Rating         float64         `json:"rating" gorm:"type:float(10,2);not null;comment:用户评分"`
	Score          float64         `json:"score" gorm:"type:float(10,2);not null;comment:综合得分"`
	Model          int             `json:"model" gorm:"type:int(10);not null;comment:模型"`
	Level          int             `json:"level" gorm:"type:int(10);not null;comment:难度"`
	Attrs          string          `json:"attrs" gorm:"type:varchar(1000);not null;comment:扩展属性"`
	Featured       int             `json:"featured" gorm:"type:int(10);not null;comment:推荐标识"`
	Published      int             `json:"published" gorm:"type:int(10);not null;comment:发布标识"`
	ResourceCount  int             `json:"resourceCount" gorm:"type:int(10);not null;comment:资料数"`
	UserCount      int             `json:"userCount" gorm:"type:int(10);not null;comment:学员数"`
	LessonCount    int             `json:"lessonCount" gorm:"type:int(10);not null;comment:课时数"`
	PackageCount   int             `json:"packageCount" gorm:"type:int(10);not null;comment:套餐数"`
	ReviewCount    int             `json:"reviewCount" gorm:"type:int(10);not null;comment:评价数"`
	ConsultCount   int             `json:"consultCount" gorm:"type:int(10);not null;comment:咨询数"`
	FavoriteCoount int             `json:"favoriteCount" gorm:"type:int(10);not null;comment:收藏数"`
}
