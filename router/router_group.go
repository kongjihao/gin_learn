// package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 路由组是为了管理一些相同的URL
	// 访问路径示例: GET http:/localhost:8080/api/v1/user
	// 访问路径示例: GET http:/localhost:8080/api/v1/product

	// 创建一个路由组/api/v1，处理GET请求
	// curl -X GET http://localhost:8080/api/v1/user
	// curl -X GET http://localhost:8080/api/v1/product
	v1 := r.Group("/api/v1")
	// {}是书写规范
	{
		v1.GET("/user", func(c *gin.Context) {
			c.String(http.StatusOK, "User Endpoint")
		})

		v1.GET("/product", func(c *gin.Context) {
			c.String(http.StatusOK, "Product Endpoint")
		})
	}

	// 创建第二个路由组，处理POST请求
	// curl -X POST http://localhost:8080/api/v2/login
	// curl -X POST http://localhost:8080/api/v2/submit
	v2 := r.Group("/api/v2")
	{
		// func (group *gin.RouterGroup) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	r.Run(":8080") // listen and serve on
}

func login(c *gin.Context) {
	name := c.DefaultPostForm("name", "Guest")
	c.String(http.StatusOK, fmt.Sprintf("Login, %s!", name))
}

func submit(c *gin.Context) {
	data := c.DefaultQuery("data", "No Data")
	c.String(http.StatusOK, fmt.Sprintf("Submitted: %s", data))
}
