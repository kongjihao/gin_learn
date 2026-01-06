package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 有的用户上传文件需要限制上传文件的类型以及上传文件的大小
	// 访问路径示例: POST http:/localhost:8080/upload
	r.POST("/upload", func(c *gin.Context) {
		// 限制上传文件的大小为8MB
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 8<<20) // 8MB, 1MB 等于 2的20次方字节

		// 获取上传的文件
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "获取上传文件出错: %s", err.Error())
			return
		}

		// 限制上传文件的类型为图片类型
		allowedTypes := map[string]bool{
			"image/jpeg": true,
			"image/png":  true,
			"image/jpg":  true,
			"image/gif":  true,
		}
		if !allowedTypes[file.Header.Get("Content-Type")] {
			c.String(http.StatusBadRequest, "不允许的文件类型: %s", file.Header.Get("Content-Type"))
			return
		}

		// 保存上传的文件到指定路径
		if err := c.SaveUploadedFile(file, "./router/upload_file/single_file/"+file.Filename); err != nil {
			c.String(http.StatusInternalServerError, "保存上传文件出错: %s", err.Error())
			return
		}

		c.String(http.StatusOK, "文件 %s 上传成功", file.Filename)
	})

	r.Run(":8080") // listen and serve on
}
