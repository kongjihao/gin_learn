// package main

import "github.com/gin-gonic/gin"

// Gin 框架的内置数据验证能力，简化参数校验流程、减少冗余的if else判断，具体内容可拆解为以下几部分
// 明确 Gin 框架 “结构体参数验证” 的核心优势：无需开发者手动解析请求数据，通过结构体标签（Tag）定义验证规则，自动完成参数校验，让代码更简洁。
type Person struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Age      int    `json:"age" form:"age" binding:"gte=0,lte=130"`
	Birthday string `json:"birthday" form:"birthday" binding:"datetime=2006-01-02" time_utc:"1"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Phone    string `json:"phone" form:"phone" binding:"omitempty,e164"` // 可选字段，若提供则需符合 E.164 格式
}

/*
json:"name"
用于指定字段在 JSON 数据中的键名。例如，Name 字段在 JSON 数据中对应的键名是 "name"。
作用：在序列化和反序列化 JSON 时，Gin 会根据这个键名映射数据。

binding:"required"
表示该字段是必填项，不能为空。
如果请求中缺少该字段或字段值为空，验证会失败。

binding:"gte=0,lte=130"
gte=0：表示字段值必须大于或等于 0。
lte=130：表示字段值必须小于或等于 130。
作用：限制 Age 字段的值范围在 0 到 130 之间。
binding:"datetime=2006-01-02"

表示字段值必须是日期格式，且格式必须符合 2006-01-02（Go 的时间格式模板）。
作用：验证 Birthday 字段是否是有效的日期字符串，例如 2026-01-07。

binding:"required,email"
required：表示字段是必填项。
email：表示字段值必须是有效的电子邮件地址。
作用：验证 Email 字段是否是非空且符合电子邮件格式。

binding:"omitempty,e164"
omitempty：表示字段是可选的。如果字段为空，则跳过验证。
e164：表示字段值必须符合 E.164 国际电话号码格式（例如 +1234567890）。
作用：如果 Phone 字段提供了值，则验证其是否符合 E.164 格式；如果未提供值，则跳过验证。
*/

func main() {
	r := gin.Default()
	// 参数校验
	// 1. 定义结构体，使用标签指定验证规则
	// 2. 在处理函数中绑定请求参数到结构体，并自动进行验证
	// 3. 根据验证结果进行相应处理
	// 优点：简化代码、提高可读性、减少错误
	// 缺点：复杂验证逻辑可能需要自定义验证器
	// 适用场景：常见的参数校验需求，如用户注册、登录等
	// 不适用场景：极其复杂或动态变化的验证逻辑
	// 注意事项：合理设计结构体和标签，避免过度复杂化
	// 扩展：可以结合自定义验证器实现更复杂的验证需求
	// 性能：对于大多数应用场景性能足够，但极端高性能需求下需评估
	// 安全性：通过验证减少无效或恶意数据的影响
	// 可维护性：结构化的验证规则易于维护和更新
	// 测试：可以通过单元测试验证不同参数组合的行为
	// 文档：使用标签清晰表达验证规则，有助于生成文档
	// 社区支持：Gin 框架有广泛的社区支持和丰富的资源
	// 版本兼容性：确保使用的 Gin 版本支持所需的验证功能

	// 处理 GET 请求，绑定查询参数并验证
	r.POST("/person", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// 验证通过
		c.JSON(200, gin.H{"message": "Validation passed!", "data": person})
	})

	r.GET("/person", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindQuery(&person); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// 验证通过
		c.JSON(200, gin.H{"message": "Validation passed!", "data": person})
	})

	r.Run(":8080")
}

// GET 请求示例:
// http://localhost:8080/person?name=kjh&age=30&birthday=1993-05-15&email=zhangsan7@google.com&phone=%2B8615667890231

// POST 请求示例:
// curl -X POST http://localhost:8080/person \
//   -H "Content-Type: application/json" \
//   -d '{
//     "name": "kjh",
//     "age": 30,
//     "birthday": "1993-05-15",
//     "email": "zhangsan7@google.com",
// 		"phone": "+8615667890231"
//   }'
