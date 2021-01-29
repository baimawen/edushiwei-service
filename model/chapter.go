package model

import "edushiwei-service/global"

type Chapter struct {
	global.COURSE_MODEL
	ParentId      int    `json:"parentId" gorm:"type:int(10);not null;index:parent_id;comment:父级编号"`
	CourseId      int    `json:"courseId" gorm:"type:int(10);not null;index:course_id;comment:课程编号"`
	Title         string `json:"title" gorm:"type:varchar(100);not null;comment:标题"`
	Summary       string `json:"summary" gorm:"type:varchar(255);not null;comment:简介"`
	Priority      int    `json:"priority" gorm:"type:int(10);not null;comment:优先级"`
	Free          int    `json:"free" gorm:"type:int(10);not null;comment:免费标识"`
	Model         int    `json:"model" gorm:"type:int(10);not null;comment:模式类型"`
	Attrs         string `json:"attrs" gorm:"type:varchar(1000);not null;comment:扩展属性"`
	Published     int    `json:"published" gorm:"type:int(10);not null;comment:发布标识"`
	ResourceCount int    `json:"resourceCount" gorm:"type:int(10);not null;comment:"资料数""`
	LessonCount   int    `json:"lessonCount" gorm:"type:int(10);not null;comment:课时数"`
	UserCount     int    `json:"userCount" gorm:"type:int(10);not null;comment:学员数"`
	ConsultCount  int    `json:"consultCount" gorm:"type:int(10);not null;comment:咨询数"`
	LikeCount     int    `json:"likeCount" gorm:"type:int(10);not null;comment:点赞数"`
}
