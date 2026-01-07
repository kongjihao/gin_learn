package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10" // 注意要用 v10 版本及以上版本，v8 已经过时，gin 也升级到 v10 了
)

type Booking struct {
	// 预订入住和离店时间
	CheckIn time.Time `form:"check_in" time_format:"2006-01-02" binding:"required,bookabledate"`
	// binding:"required,gtfield=CheckIn" 表示必须大于 CheckIn 字段的值，且为必填参数
	CheckOut time.Time `form:"check_out" time_format:"2006-01-02" binding:"required,gtfield=CheckIn"`
}

// 自定义的校验方法 (符合 v10 的签名)
func bookableDate(fl validator.FieldLevel) bool {
	// 获取字段值，这里是 time.Time 类型
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		// 验证逻辑：必须是今天或未来的日期
		today := time.Now().Truncate(24 * time.Hour)  // Truncate
		return date.Equal(today) || date.After(today)
	}
	return false
}

func main() {
	r := gin.Default()

	// 将我们自定义的校验方法注册到 validator中
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	r.GET("/book", getBookable)

	r.Run(":8080")
}

func getBookable(c *gin.Context) {
	var book Booking

	if err := c.ShouldBindQuery(&book); err == nil {
		c.String(http.StatusOK, "Booking from %s to %s", book.CheckIn.Format("2006-01-02"), book.CheckOut.Format("2006-01-02"))
	} else {
		c.String(http.StatusBadRequest, "Booking bind err: %v", err.Error())
	}
}

/*
测试命令：
curl -X GET "http://localhost:8080/book?check_in=2026-11-07&check_out=2026-11-20"
curl -X GET "http://localhost:8080/book?check_in=2019-09-07&check_out=2026-11-20" // check_in 早于今天
curl -X GET "http://localhost:8080/book?check_in=2026-11-07&check_out=2026-11-01" // check_out 早于 check_in
*/
