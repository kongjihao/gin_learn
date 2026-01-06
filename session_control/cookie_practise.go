// package main

/*
模拟实现权限验证中间件:

有2个路由，login和home
login用于设置cookie
home是访问查看信息的请求
在请求home之前，先跑中间件代码，检验是否存在cookie
访问home，会显示错误，因为权限校验未通过
*/

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("username"); err == nil {
			if cookie != "" {
				c.Set("username", cookie)
				c.Next()
				return // 验证通过，继续处理请求
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})

		c.Abort() // 没有验证通过，中止后续处理
	}
}

func main() {
	r := gin.Default()

	r.GET("/home", AuthMiddleware(), func(c *gin.Context) {
		username, _ := c.Get("username")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s, Welcome to the home page!", username),
		})
	})

	r.GET("/login/:name", func(c *gin.Context) {
		username := c.Param("name")
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username cannot be empty"})
			return
		}
		c.SetCookie("username", username, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "Logged in successfully!",
		})
	})

	r.Run(":8080")
}
