package v1

import (
	"edushiwei-service/model/response"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	response.Ok(c)
}
