package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 表单参数：
	// 表单传输为post请求，http常见的传输格式为四种：
	// 	application/json
	// 	application/x-www-form-urlencoded
	// 	application/xml
	// 	multipart/form-data
	// 表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数

	r := gin.Default()

	// 访问路径示例: POST http:/localhost:8080/form
	// 表单参数: name=John
	r.POST("/form", func(c *gin.Context) {
		// func (c *gin.Context) DefaultPostForm(key string, defaultValue string) string
		types := c.DefaultPostForm("type", "post") // 如果type参数不存在，则使用默认值type = "post"
		username := c.PostForm("username")         // 获取表单中的username参数
		password := c.PostForm("userpassword")     // 获取表单中的userpassword参数

		c.String(http.StatusOK, fmt.Sprintf("username: %s, password: %s, type: %s", username, password, types))
	})

	r.Run(":8080") // listen and serve on
}
