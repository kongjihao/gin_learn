// package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 可以通过gin.Context的Param方法获取路径(API)参数
	// 例如：GET http:/localhost:8080/user/kongjihao/edit
	// name参数的值为kongjihao，action参数的值为action/edit
	// 注意：使用Param方法获取的参数值不包含/，如果参数值本身包含/，则需要手动处理
	// 例如上面的action参数值为edit，使用Param方法获取到的值为/edit，需要使用strings.Trim方法去除前后的/ (action = strings.Trim(action, "/"))

	r.GET("/user/:name/*action", func(c *gin.Context) { // 注意这里的*action表示action参数可以包含/，即可以匹配多个路径段
		name := c.Param("name")
		action := c.Param("action")

		action = strings.Trim(action, "/")

		c.JSON(http.StatusOK, gin.H{
			"name":   name,
			"action": action,
		})
	})

	r.Run(":8080") // listen and serve on
}
