// package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 从查询参数获取值
	// 访问路径示例: GET http:/localhost:8080/user?name=John

	r.GET("/user", func(c *gin.Context) {
		// 如果name参数不存在，则使用默认值name = "Guest"
		// func (c *gin.Context) DefaultQuery(key string, defaultValue string) string
		name := c.DefaultQuery("name", "Guest")
		c.String(http.StatusOK, fmt.Sprintf("Hello, %s!", name))
	})

	r.Run(":8080") // listen and serve on
}
