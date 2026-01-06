package main

// Gin 框架本身未内置 Session 功能，需结合gorilla/sessions实现
// 本代码基于 Gin+gorilla/sessions 实现用户登录态管理，包含登录、鉴权、退出登录三个核心接口。
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// 定义Session相关常量
const SessionName = "user_session" // Session名称

// 声明全局存储对象（需确保密钥保密，生产环境建议从配置文件读取）
// 用 gorilla/sessions 创建 Cookie 型 Session 存储，这个store并不是 “存储 Session 数据的容器”，而是Session 的管理器，它的核心作用是：
// 1. 把你要保存的 Session 数据（比如user_id=1001）加密 + 签名后，直接写入客户端浏览器的 Cookie 中；
// 2. 当客户端发起后续请求时，从 Cookie 中读取加密后的 Session 数据，解密 + 验证签名（确保数据没被篡改），还原出user_id等信息。
// “Cookie 型 Session” 就是把 Session 数据直接存在客户端 Cookie 里，服务端不保存任何 Session 数据，只靠密钥保证数据的安全性（加密防泄露、签名防篡改）。
var store = sessions.NewCookieStore( // 是 Session 的一种实现方式
	[]byte("your-secret-key-1"), // 签名密钥，任意设置
	// gorilla/sessions 的 Cookie 加密功能要求加密密钥长度必须是 16/24/32 字节（对应 AES-128/AES-192/AES-256 加密），否则将导致session.Save()时抛出加密密钥错误
	[]byte("1234567890123456"), // 修改为正好 16 个字节 (AES-128)
	// 或者： []byte("123456789012345678901234"), // 24字节
	// 或者： []byte("12345678901234567890123456789012"), // 32字节
)

// 初始化存储配置（可选，根据需求调整）,配置了 HttpOnly、SameSite 等安全属性，Session 默认有效期 1 小时
// init 函数会在包加载时自动调用
func init() {
	// 设置Cookie的HttpOnly属性（防止JS脚本访问，增强安全性）
	store.Options.HttpOnly = true
	// 设置Cookie的Secure属性（仅HTTPS环境下传输，生产环境建议开启）
	// store.Options.Secure = true
	// 设置Cookie的SameSite属性（防止CSRF攻击，可选SameSiteLaxMode/SameSiteStrictMode）
	store.Options.SameSite = http.SameSiteLaxMode
	// 设置Session默认有效期（单位：秒）
	store.Options.MaxAge = 3600 // 1小时
}

// GetSession 获取当前请求的Session对象
func GetSession(c *gin.Context) (*sessions.Session, error) {
	return store.Get(c.Request, SessionName)
}

// SetSessionValue 向Session中设置键值对
func SetSessionValue(c *gin.Context, key string, value interface{}) error {
	session, err := GetSession(c)
	if err != nil {
		return err
	}
	session.Values[key] = value
	// 保存更改到响应
	return session.Save(c.Request, c.Writer)
}

// GetSessionValue 从Session中获取值
func GetSessionValue(c *gin.Context, key string) (interface{}, error) {
	session, err := GetSession(c)
	if err != nil {
		return nil, err
	}
	return session.Values[key], nil
}

// DeleteSession 删除当前Session
func DeleteSession(c *gin.Context) error {
	session, err := GetSession(c)
	if err != nil {
		return err
	}
	// 设置MaxAge为-1删除Session
	session.Options.MaxAge = -1
	return session.Save(c.Request, c.Writer)
}

// 模拟用户登录，登录成功后存储用户ID到Session
func loginHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// 绑定并验证请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数错误"})
		return
	}

	// 模拟验证用户名密码（实际项目需连接数据库验证）
	if req.Username == "root" && req.Password == "admin" {
		// 存储用户ID到Session
		if err := SetSessionValue(c, "user_id", 1001); err != nil {
			c.JSON(500, gin.H{"error": "Session存储失败"})
			return
		}
		c.JSON(200, gin.H{"message": "登录成功"})
		return
	}

	c.JSON(401, gin.H{"error": "用户名或密码错误"})
}

// 需登录后才能访问的接口（如用户个人中心）
func profileHandler(c *gin.Context) {
	// 从Session中获取用户ID
	userID, err := GetSessionValue(c, "user_id")
	if err != nil {
		c.JSON(500, gin.H{"error": "获取Session失败"})
		return
	}
	if userID == nil {
		c.JSON(403, gin.H{"error": "请先登录"})
		return
	}

	// 模拟返回用户信息（实际项目需从数据库查询）
	c.JSON(200, gin.H{
		"message": "个人中心",
		"user_id": userID,
	})
}

// 用户退出登录，删除Session
func logoutHandler(c *gin.Context) {
	if err := DeleteSession(c); err != nil {
		c.JSON(500, gin.H{"error": "退出登录失败"})
		return
	}
	c.JSON(200, gin.H{"message": "退出登录成功"})
}

func main() {
	// 初始化Gin引擎
	r := gin.Default()

	// 注册路由
	// 最好用postman测试
	// curl -X POST http://127.0.0.1:8080/login -H 'content-type: application/json' -d '{"username":"root","password":"admin"}'
	r.POST("/login", loginHandler)    // 登录（存储Session）
	r.GET("/profile", profileHandler) // 个人中心（获取Session）
	r.POST("/logout", logoutHandler)  // 退出登录（删除Session）

	// 启动服务，监听8080端口
	if err := r.Run(":8080"); err != nil {
		panic("Gin服务启动失败: " + err.Error())
	}
}
