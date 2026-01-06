// package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 单文件上传
	// 访问路径示例: POST http:/localhost:8080/upload
	r.POST("/upload", func(c *gin.Context) {
		// 获取上传的文件
		// func (c *gin.Context) FormFile(name string) (*multipart.FileHeader, error)
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get formFile err: %s", err.Error()))
			return
		}

		// 将文件保存到指定路径
		// func (c *gin.Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error
		dst := "./router/upload_file/single_file/" + file.Filename
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully.", file.Filename))
	})

	r.Run(":8080") // listen and serve on
}
