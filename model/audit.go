package model

import "time"

type Audit struct {
	ID        uint      `gorm:"primarykey;type:int(10);not null;comment:主键编号"`
	UserId    int       `json:"userId" gorm:"type:int(10);not null;index:user_id;comment:用户编号"`
	UserName  string    `json:"userName" gorm:"type:varchar(30);not null;comment:用户名称"`
	UserIP    string    `json:"userIP" gorm:"type:varchar(30);not null;comment:用户IP"`
	ReqRouter string    `json:"reqRouter" gorm:"type:varchar(50);not null;comment:请求路由"`
	ReqPath   string    `json:"reqPath" gorm:"type:varchar(100);not null;comment:请求路径"`
	ReqData   string    `json:"reqData" gorm:"type:text;not null;comment:请求数据"`
	CreatedAt time.Time `gorm:"comment:创建时间;not null;autoCreateTime"`
}
