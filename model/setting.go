package model

type Setting struct {
	ID        uint   `gorm:"primarykey;type:int(10);not null;comment:主键编号"`
	Section   string `json:"section" gorm:"type:varchar(50);not null;index:section_key;comment:配置组"`
	ItemKey   string `json:"itemKey" gorm:"type:varchar(50);not null;index:section_key;comment:配置项"`
	ItemValue string `json:"itemValue" gorm:"type:text;not null;comment:配置值"`
}
