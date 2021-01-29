package service

import (
	"edushiwei-service/global"
	"edushiwei-service/model"
	"errors"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func FindUserByUuid(uuid string) (err error, user *model.SysUser) {
	var u model.SysUser
	if err = global.COURSE_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
