package http

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")                                                                                   // 允许所有源跨域请求，或指定源如 "http://localhost:3000"
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")                                                                           // 允许发送 Cookie 等认证信息
		c.Writer.Header().Set("Access-Control-Allow-Headers", "x-token,Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization") // 允许的头部字段
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 允许发送 Cookie 等认证信息
		// 允许的方法
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // 对于预检请求，直接返回 204 No Content 不处理请求体。
			return
		}
		c.Next() // 继续执行其他中间件或路由处理函数。
	}
}
