package middleware

import (
	"github.com/gin-gonic/gin"
)

func Access() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Next()后就执行真实的路由函数，路由执行完成之后接着走time.Since(t)
		// gin设置响应头，设置跨域
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Action, Module, X-PINGOTHER, Content-Type, Content-Disposition")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.JSON(204, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
