# Cookie介绍

1. HTTP是无状态协议，服务器不能记录浏览器的访问状态，也就是说服务器不能区分两次请求是否由同一个客户端发出
2. Cookie就是解决HTTP协议无状态的方案之一，中文是小甜饼的意思
3. Cookie实际上就是**服务器保存在浏览器上的一段信息**。浏览器有了Cookie之后，每次向服务器发送请求时都会同时将该信息发送给服务器，服务器收到请求后，就可以根据该信息处理请求
4. **Cookie由服务器创建，并发送给浏览器，最终由浏览器保存**

# Cookie的用途

> 测试服务端发送cookie给客户端，客户端请求时携带cookie

## 访问别人的网站，cookie值有很多键值对，是怎么做到的

当你访问一个网站时，浏览器会通过 HTTP 请求和响应与服务器交换数据。服务器可以通过 HTTP 响应头中的 Set-Cookie 字段向浏览器发送一个或多个 Cookie。以下是实现多个键值对 Cookie 的机制：

1. Cookie 的结构
Cookie 是一种键值对的存储机制，格式如下：
``` bash
Set-Cookie: key1=value1; key2=value2; key3=value3; ...
```
每个键值对可以附加一些属性，例如：

- Expires 或 Max-Age：指定 Cookie 的过期时间。
- Domain：指定 Cookie 的作用域（哪个域名可以访问）。
- Path：指定 Cookie 的路径（哪个路径下的资源可以访问）。
- HttpOnly：限制 Cookie 只能通过 HTTP 协议访问，JavaScript 无法访问。
- Secure：限制 Cookie 只能通过 HTTPS 传输。

2. 服务器如何设置多个 Cookie
服务器可以通过多次 Set-Cookie 响应头来设置多个 Cookie。例如：
``` bash
Set-Cookie: user_id=12345; Path=/; HttpOnly
Set-Cookie: session_token=abcdef; Path=/; Secure
Set-Cookie: theme=dark; Path=/
```
浏览器会将这些 Cookie 存储起来，并在后续请求中通过 Cookie 请求头发送回服务器：
``` bash
Cookie: user_id=12345; session_token=abcdef; theme=dark
```

3. 在代码中设置多个 Cookie
以 Go 的 Gin 框架为例，服务器可以通过以下代码设置多个 Cookie：

``` go
package main

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
```

4. 浏览器如何管理多个 Cookie
浏览器会根据以下规则管理 Cookie：

作用域匹配：只有符合 Domain 和 Path 的 Cookie 才会发送给服务器。
过期时间：过期的 Cookie 会被自动删除。
优先级：如果多个 Cookie 的键相同，浏览器会使用最新的值。
5. 总结
通过 Set-Cookie 响应头，服务器可以设置多个键值对的 Cookie。浏览器会根据规则存储和发送这些 Cookie，从而实现状态管理和用户会话的功能。