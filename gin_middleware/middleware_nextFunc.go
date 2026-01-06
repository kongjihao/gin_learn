// package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func globalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行~")

		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")

		// 执行函数
		c.Next() // 执行后续的函数，直到所有函数执行完毕才会继续往下走

		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)

		t2 := time.Since(t)
		fmt.Println("time spend:", t2)
	}
}

func main() {
	r := gin.Default()
	// 注册中间件
	r.Use(globalMiddleware())
	// {}为了代码规范
	{
		r.GET("/test", func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Printf("从Context中取到的值: %v\n", req)
			// 页面接收
			c.JSON(http.StatusOK, gin.H{"request": req})
		})
	}

	r.Run(":8080")
}
