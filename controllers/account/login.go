package account

import (
	"blog-go/utils/response"
	"blog-go/utils/session"
	"fmt"

	"github.com/gin-gonic/gin"
)

// LoginHandler 用户登陆
func LoginHandler(c *gin.Context) {
	var err error
	var user Userinfo
	sess := session.Get(c.Request)
	//err = c.BindJSON(&user)
	user.Username = c.Query("username")
	user.Password = c.Query("password")
	fmt.Println(user)
	if err != nil {
		panic("报错")
	}
	if len(user.Username) == 0 || len(user.Password) == 0 {
		response.Error(c, response.ErrCodeLoginParameter)
		return
	}
	if user.Username == "hank997" && user.Password == "abc" {
		sess = session.NewSessionOptions(&session.SessOptions{
			CAttrs: map[string]interface{}{"UserName": user.Username},
			Attrs:  map[string]interface{}{"Count": 1},
		})
		session.Add(sess, c.Writer)
		response.Success(c, "login sussed")
	} else {
		response.Error(c, response.ErrCodeLoginParameter)
	}
}
