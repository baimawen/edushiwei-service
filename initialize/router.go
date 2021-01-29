package initialize

import (
	"edushiwei-service/global"
	"edushiwei-service/middleware"
	"edushiwei-service/router"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.COURSE_CONFIG.Local.Path, http.Dir(global.COURSE_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	global.COURSE_LOG.Info("use middlerware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.COURSE_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.COURSE_LOG.Info("register swagger handler")
	// 方便同意添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册记录功能路由 不做鉴权
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		// router.InitApiRouter(PrivateGroup) //注册功能路由
	}
	global.COURSE_LOG.Info("router register success")
	return Router
}
