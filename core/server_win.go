package core

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 使用endless包时会报错 “undefined: syscall.SIGUSR1”，原因是syscall包里的常量会根据当前操作系统做出选择，而win下的signal信号里没有这些信号。
func initServerWin(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
