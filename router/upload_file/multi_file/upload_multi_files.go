package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 多文件上传处理，先限制下上传大小
	r.MaxMultipartMemory = 8 << 20 // 8 MiB, 多文件上传时gin有默认限制为32 MiB，通过设置MaxMultipartMemory可以修改该限制
	r.POST("/upload", func(c *gin.Context) {
		// func (c *gin.Context) MultipartForm() (*multipart.Form, error)
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "获取上传文件出错: %s", err.Error())
			return
		}

		// 获取所有图片
		files := form.File["files"]

		// 遍历所有文件并保存
		for _, file := range files {
			if err := c.SaveUploadedFile(file, "./router/upload_file/multi_file/"+file.Filename); err != nil {
				c.String(http.StatusInternalServerError, "保存上传文件出错: %s", err.Error())
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("%d 个文件上传成功！", len(files)))
	})

	r.Run(":8000") // listen and serve on

}
