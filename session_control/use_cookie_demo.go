// package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("username")
		if err != nil {
			cookie = "NotSet" // 如果没有cookie，则返回NotSet,并且设置cookie，只有第一次没有cookie时会返回NotSet，之后就会返回设置的cookie值

			// 给客户端设置cookie
			// func (c *gin.Context) SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool)
			// name string, cookie的名字
			// value string, cookie的值
			// maxAge int, 单位为秒，cookie的过期时间
			// path string, cookie所在目录
			// domain string, 域名
			// secure bool, 是否智能通过https访问
			// httpOnly bool, 是否允许别人通过js获取自己的cookie

			c.SetCookie("username", "kjh", 3600, "/", "localhost", false, true)

		}
		fmt.Printf("cookie的值是： %s\n", cookie)
	})

	r.Run(":8080")
}
