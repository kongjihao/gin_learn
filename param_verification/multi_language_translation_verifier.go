// package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	zh_tw_translations "gopkg.in/go-playground/validator.v9/translations/zh_tw"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

type User struct {
	Username string `form:"user_name" validate:"required"`
	Tagline  string `form:"tag_line" validate:"required,lt=10"`
	Tagline2 string `form:"tag_line2" validate:"required,gt=1"`
}

func main() {
	en := en.New()
	zh := zh.New()
	zh_tw := zh_Hant_TW.New()
	Uni = ut.New(en, zh, zh_tw)
	Validate = validator.New()

	route := gin.Default()
	route.GET("/testing", startPage)
	route.POST("/testing", startPage)
	route.Run(":8080")
}

func startPage(c *gin.Context) {
	//这部分应放到中间件中
	locale := c.DefaultQuery("locale", "zh")
	trans, _ := Uni.GetTranslator(locale)

	// 每次请求都重新注册翻译
	switch locale {
	case "zh":
		zh_translations.RegisterDefaultTranslations(Validate, trans)
		// 移除自定义覆盖，使用默认的中文翻译
	case "en":
		en_translations.RegisterDefaultTranslations(Validate, trans)
		// 移除自定义覆盖，使用默认的英文翻译
	case "zh_tw":
		zh_tw_translations.RegisterDefaultTranslations(Validate, trans)
		// 移除自定义覆盖，使用默认的繁体中文翻译
	default:
		zh_translations.RegisterDefaultTranslations(Validate, trans)
		// 移除自定义覆盖，使用默认的中文翻译
	}

	// 移除自定义的required翻译，让它使用各个语言的默认翻译
	// 这样不同语言就会返回不同的错误信息

	//这块应该放到公共验证方法中
	user := User{}
	c.ShouldBind(&user)
	fmt.Println("用户数据:", user)
	fmt.Println("当前语言:", locale)

	err := Validate.Struct(user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		fmt.Println("验证错误:", sliceErrs)
		c.JSON(200, gin.H{
			"user":   user,
			"errors": sliceErrs,
			"locale": locale,
		})
		return
	}

	c.JSON(200, gin.H{
		"user":    user,
		"message": "验证通过",
		"locale":  locale,
	})
}

/*
测试链接：

1. 中文错误信息：
http://localhost:8080/testing?user_name=&tag_line=92&tag_line2=32&locale=zh
返回示例：
{
  "errors": [
    "Username为必填字段"
  ],
  "locale": "zh",
  "user": {
    "Tagline": "92",
    "Tagline2": "32",
    "Username": ""
  }
}

2. 英文错误信息：
http://localhost:8080/testing?user_name=&tag_line=92&tag_line2=32&locale=en
返回示例：
{
  "errors": [
    "Username is a required field"
  ],
  "locale": "en",
  "user": {
    "Tagline": "92",
    "Tagline2": "32",
    "Username": ""
  }
}

3. 繁体中文错误信息：
http://localhost:8080/testing?user_name=&tag_line=92&tag_line2=32&locale=zh_tw
返回示例：
{
  "errors": [
    "Username為必填欄位"
  ],
  "locale": "zh_tw",
  "user": {
    "Tagline": "92",
    "Tagline2": "32",
    "Username": ""
  }
}

4. 验证通过的情况：
http://localhost:8080/testing?user_name=枯藤&tag_line=9&tag_line2=3&locale=zh
返回：
{
  "locale": "zh",
  "message": "验证通过",
  "user": {
    "Tagline": "9",
    "Tagline2": "3",
    "Username": "枯藤"
  }
}
*/

// package main

// import (
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/locales/en"
// 	"github.com/go-playground/locales/zh"
// 	"github.com/go-playground/locales/zh_Hant_TW"
// 	ut "github.com/go-playground/universal-translator"
// 	"gopkg.in/go-playground/validator.v9" // 这是v9版本，推荐用最新v10版本
// 	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
// 	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
// 	zh_tw_translations "gopkg.in/go-playground/validator.v9/translations/zh_tw"
// )

// var (
// 	Uni      *ut.UniversalTranslator
// 	Validate *validator.Validate
// )

// type User struct {
// 	Username string `form:"user_name" validate:"required"`
// 	Tagline  string `form:"tag_line" validate:"required,lt=10"`
// 	Tagline2 string `form:"tag_line2" validate:"required,gt=1"`
// }

// func main() {
// 	en := en.New()
// 	zh := zh.New()
// 	zh_tw := zh_Hant_TW.New()
// 	Uni = ut.New(en, zh, zh_tw)
// 	Validate = validator.New()

// 	route := gin.Default()
// 	route.GET("/testing", startPage)
// 	route.POST("/testing", startPage)
// 	route.Run(":8080")
// }

// func startPage(c *gin.Context) {
// 	//这部分应放到中间件中
// 	locale := c.DefaultQuery("locale", "zh")
// 	trans, _ := Uni.GetTranslator(locale)
// 	switch locale {
// 	case "zh":
// 		zh_translations.RegisterDefaultTranslations(Validate, trans)
// 		break
// 	case "en":
// 		en_translations.RegisterDefaultTranslations(Validate, trans)
// 		break
// 	case "zh_tw":
// 		zh_tw_translations.RegisterDefaultTranslations(Validate, trans)
// 		break
// 	default:
// 		zh_translations.RegisterDefaultTranslations(Validate, trans)
// 		break
// 	}

// 	//自定义错误内容
// 	Validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
// 		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
// 	}, func(ut ut.Translator, fe validator.FieldError) string {
// 		t, _ := ut.T("required", fe.Field())
// 		return t
// 	})

// 	//这块应该放到公共验证方法中
// 	user := User{}
// 	c.ShouldBind(&user)
// 	fmt.Println(user)
// 	err := Validate.Struct(user)
// 	if err != nil {
// 		errs := err.(validator.ValidationErrors)
// 		sliceErrs := []string{}
// 		for _, e := range errs {
// 			sliceErrs = append(sliceErrs, e.Translate(trans))
// 		}
// 		c.String(200, fmt.Sprintf("%#v", sliceErrs))
// 	}
// 	c.String(200, fmt.Sprintf("%#v", user))
// }

/*
正确的链接：http://localhost:8080/testing?user_name=枯藤&tag_line=92&tag_line2=32&locale=zh

http://localhost:8080/testing?user_name=枯藤&tag_line=9&tag_line2=3&locale=en 返回英文的验证错误信息

http://localhost:8080/testing?user_name=枯藤&tag_line=9&tag_line2=3&locale=zh 返回中文的验证错误信息
*/
