package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建 Gin 引擎实例
	// gin.Default() 包含日志和恢复中间件
	// gin.New() 创建空白引擎，不包含任何中间件
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	r.Run()

}
