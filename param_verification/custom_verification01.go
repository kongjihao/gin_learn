// package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10" // 注意要用 v10 版本及以上版本，v8 已经过时，gin 也升级到 v10 了
)

/*
对绑定解析到结构体上的参数，自定义验证功能
比如我们要对 name 字段做校验，要不能为空，并且不等于 admin ，类似这种需求，就无法 binding 现成的方法
需要我们自己验证方法才能实现 官网示例（https://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Custom_Functions）
这里需要下载引入下 gopkg.in/go-playground/validator.v8
*/
type Person struct {
	Age int `form:"age" binding:"required,gt=10"`
	// 3、使用，在参数 binding 上使用自定义的校验方法函数注册时候的名称
	Name    string `form:"name" binding:"NotNullAndAdmin"`
	Address string `form:"address" binding:"required"`
}

// 1、自定义的校验方法 (符合 v10 的签名)
func nameNotNullAndAdmin(fl validator.FieldLevel) bool {
	// fl.Field() 直接获取当前字段值, 注意需要根据字段类型使用不同的方法获取,
	// 这里struct中 name 是 string 类型,所以用 String()
	value := fl.Field().String()

	// 验证逻辑
	return value != "" && value != "admin"
}

func main() {
	r := gin.Default()

	// 2、将我们自定义的校验方法注册到 validator中
	// 2、注册到 v10 引擎
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NotNullAndAdmin", nameNotNullAndAdmin)
	}

	/*
	   curl -X GET "http://127.0.0.1:8080/info?name=&age=12&address=beijing"
	   curl -X GET "http://127.0.0.1:8080/info?name=admin&age=12&address=beijing"
	   curl -X GET "http://127.0.0.1:8080/info?name=kjh&age=12&address=beijing"
	*/
	r.GET("/info", func(c *gin.Context) {
		var person Person
		if e := c.ShouldBindQuery(&person); e == nil {
			c.String(http.StatusOK, "%v", person)
		} else {
			c.String(http.StatusBadRequest, "person bind err: %v", e.Error())
		}
	})

	r.Run(":8080")
}
