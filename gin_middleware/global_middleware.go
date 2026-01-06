// package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func globalMiddleware01() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")

		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)

		t2 := time.Since(t)
		fmt.Println("spend:", t2)
	}
}

func main() {
	r := gin.Default()
	// 注册中间件
	r.Use(globalMiddleware01())
	// {}为了代码规范
	{
		r.GET("/test01", func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Printf("从Context中取到的值: %v\n", req)
			// 页面接收
			c.JSON(http.StatusOK, gin.H{"request": req})
		})
	}

	r.Run(":8080")
}
