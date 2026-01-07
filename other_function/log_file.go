// package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor() // 用于禁用控制台中的颜色输出。

	// Logging to a file, 真正线上项目中不推荐此种用法推荐使用第三方zap库或者中间件Logger()
	file, _ := os.Create("./other_function/gin_server_log.log")
	// 只将日志写入项目日志文件
	// gin.DefaultWriter = io.MultiWriter(file)

	// 如果需要同时将日志写入文件和控制台，使用以下代码
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout) // 注意这里是赋值，不是初始化

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.Run() //默认为localhost:8080
}
