package server

import (
	"github.com/cxbooks/cxbooks/server/model"
	"github.com/cxbooks/cxbooks/zlog"
	"github.com/gin-gonic/gin"
)

// UserInfo GET /api/user/info
// 获取用户信息
func UserInfo(c *gin.Context) {
	zlog.D(`获取用户信息`)
	if info, ok := c.Get(XUserInfoTag); ok {

		c.JSON(200, SUCCESS.Tr(ZH).With(info))
		return
	}

	c.JSON(404, ErrNoFound.Tr(ZH))

}

func UserMessages(c *gin.Context) {
	zlog.D(`获取用户信息`)
	id := c.GetInt64(XUserIDTag)

	if id == 0 {
		zlog.E(`用户获取失败`)
	}

	messages, err := model.FindMessages(srv.orm, id, 10)

	if err != nil {
		c.JSON(404, ErrNoFound.Tr(ZH))
	}
	c.JSON(200, SUCCESS.Tr(ZH).With(messages))
}

// LoginReq  登陆
type LoginReq struct {
	Account        string `form:"account" json:"account" binding:"required"`
	Password       string `form:"password" json:"password" binding:"required"`
	RecaptchaToken string `form:"recaptcha_token" json:"recaptcha_token"`
}

func SignIn(c *gin.Context) {

	// ip := c.ClientIP()

	req := struct {
		Account        string `form:"account" json:"account" binding:"required"`
		Password       string `form:"password" json:"password" binding:"required"`
		RecaptchaToken string `form:"recaptcha_token" json:"recaptcha_token"`
	}{}

	if err := c.ShouldBind(req); err != nil {
		c.JSON(400, ErrUserPassword.Tr(ZH))
		return
	}

	user, err := model.FirstUserByAccount(srv.orm, req.Account)

	if err != nil && user.VerifyPassword(req.Password) {
		zlog.E(`查询用户失败：`, err)
		c.JSON(200, ErrUserPassword.Tr(ZH))
		return
	}

	if user.Locked {
		zlog.E(`查询用户失败：`, err)
		c.JSON(200, ErrUserLocked.Tr(ZH))
		return
	}

	sess, err := user.CreateSession(srv.orm, c.GetHeader(`User-Agent`), c.GetHeader(`Referer`), c.ClientIP())
	if err != nil {
		zlog.E(`创建session失败：`, err)
		c.JSON(200, ErrInnerServer.Tr(ZH))
		return
	}

	c.SetCookie(UserSessionTag, sess.Session, 0, "/", c.Request.Host, true, true)
	//TODO recording login log to DB

	c.JSON(200, SUCCESS.Tr(ZH))
}

// func SignUp(c *gin.Context) {

// }

func SignOut(c *gin.Context) {

}
func UserUpdate(c *gin.Context) {

}
func UserReset(c *gin.Context) {

}
func UserSendActive(c *gin.Context) {

}
