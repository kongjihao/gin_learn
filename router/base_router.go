// package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 设置gin为发布模式，还有debug和测试模式

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	r.POST("/xxxpost", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world post",
		})
	})

	r.PUT("/xxxput", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world put",
		})
	})

	r.Run(":8080") // listen and serve on
}
