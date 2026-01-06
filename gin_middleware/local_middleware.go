// package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time spend:", t2)
	}
}

func main() {
	r := gin.Default()

	// Local middleware
	r.GET("/test", MiddleWare(), func(c *gin.Context) {
		// 取值，console打印
		value, _ := c.Get("request")
		fmt.Printf("走了中间件，从Context中取到的值: %v\n", value)
		// 页面接收
		c.JSON(http.StatusOK, gin.H{"request": value})
	})

	r.GET("/test01", func(c *gin.Context) {
		// 取值，console打印
		c.JSON(http.StatusOK, gin.H{"info": "没有走中间件"})
	})

	r.Run(":8080")
}
