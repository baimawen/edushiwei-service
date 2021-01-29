package core

import (
	"edushiwei-service/global"
	"edushiwei-service/initialize"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.COURSE_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.COURSE_CONFIG.System.Addr)
	s := initServerWin(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.COURSE_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	欢迎使用 识微云课堂
	当前版本:V0.1.0
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.COURSE_LOG.Error(s.ListenAndServe().Error())
}
