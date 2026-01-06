// package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/set_cookies", func(c *gin.Context) {
        // 设置多个 Cookie
        c.SetCookie("user_id", "12345", 3600, "/", "localhost", false, true)
        c.SetCookie("session_token", "abcdef", 3600, "/", "localhost", false, true)
        c.SetCookie("theme", "dark", 3600, "/", "localhost", false, true)

        c.JSON(200, gin.H{
            "message": "Cookies set successfully",
        })
    })

    r.GET("/get_cookies", func(c *gin.Context) {
        // 获取所有 Cookie
        userID, _ := c.Cookie("user_id")
        sessionToken, _ := c.Cookie("session_token")
        theme, _ := c.Cookie("theme")

        c.JSON(200, gin.H{
            "user_id":       userID,
            "session_token": sessionToken,
            "theme":         theme,
        })
    })

    r.Run(":8080")
}