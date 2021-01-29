package model

type Area struct {
	ID   uint   `gorm:"primarykey;type:int(10);not null;comment:主键"`
	Type int    `json:"type" gorm:"type:int(10);not null;comment:类型"`
	Code string `json:"code" gorm:"type:varchar(30);not null;comment:编码"`
	Name string `json:"name" gorm:"type:varchar(30);not null;comment:名称"`
}
