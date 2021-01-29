package router

import (
	v1 "edushiwei-service/api/v1"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("")
	{
		BaseRouter.GET("ping", v1.Ping)
	}
	return BaseRouter
}
