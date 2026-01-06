package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义程序计时中间件，然后定义2个路由，执行函数后应该打印统计的执行时间，如下：
func myTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		fmt.Printf("请求处理时间: %v\n", latency)
	}
}

func main() {
	r := gin.Default()
	r.Use(myTime())
	respGroup := r.Group("/v1")
	{
		respGroup.GET("/ping", func(c *gin.Context) {
			time.Sleep(2 * time.Second)
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		respGroup.GET("/hello", func(c *gin.Context) {
			time.Sleep(1 * time.Second)
			c.JSON(http.StatusOK, gin.H{
				"message": "hello world",
			})
		})
	}

	r.Run(":8080")
}
