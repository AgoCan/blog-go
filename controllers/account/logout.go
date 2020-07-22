package account

import (
	"blog-go/utils/session"

	"github.com/gin-gonic/gin"
)

// LogoutHandler 用户注销
func LogoutHandler(c *gin.Context) {
	sess := session.Get(c.Request)
	session.Remove(sess, c.Writer)
}
