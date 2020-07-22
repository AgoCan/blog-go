package auth

import (
	"blog-go/utils/response"
	"blog-go/utils/session"
	"fmt"

	"github.com/gin-gonic/gin"
)

func isLogin(c *gin.Context) bool {
	statusTml, ok := c.Get(loginStatus)
	if !ok {
		return false
	}
	status, ok := statusTml.(int)
	if !ok {
		return false
	}
	if status == 0 {
		return false
	}
	return true
}

func processRequest(c *gin.Context) {
	s := session.Global.Get(c.Request)
	if s == nil {
		return
	}
	c.Set(loginStatus, int(1))
	c.Set(loginUsername, s.CAttr("UserName"))
}

// Middleware 认证中间件
func Middleware(c *gin.Context) {
	processRequest(c)
	isLogin := isLogin(c)
	if !isLogin {
		response.Error(c, response.ErrCodeNotLogin)
		//中断当前请求
		c.Abort()
		return
	}
	fmt.Println(123)
	c.Next()
}
