package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// Example: http://localhost:8000/login/root/admin
	r.GET("/login/:user/:password", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.User == "root" && login.Password == "admin" {
			c.JSON(http.StatusOK, gin.H{"status": "200"})
			return
		} else {
			c.JSON(203, gin.H{"status": "203"})
			return
		}
	})

	r.Run(":8000")
}
