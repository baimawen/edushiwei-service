package main

import (
	"edushiwei-service/core"
	"edushiwei-service/global"
	"edushiwei-service/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.COURSE_VP = core.Viper()          // 初始化Viper
	global.COURSE_LOG = core.Zap()           // 初始化zap日志库
	global.COURSE_DB = initialize.Gorm()     // gorm连接数据库
	initialize.MysqlTables(global.COURSE_DB) // 初始化表
	// 程序结束前关闭数据库连接
	db, _ := global.COURSE_DB.DB()
	defer db.Close()

	core.RunWindowsServer()
}
