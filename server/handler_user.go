package server

import (
	"github.com/cxbooks/cxbooks/server/model"
	"github.com/cxbooks/cxbooks/server/zlog"
	"github.com/gin-gonic/gin"
)

// UserInfo GET /api/user/info
// 获取用户信息
func UserInfoHandler(c *gin.Context) {
	zlog.D(`获取用户信息`)
	if info, ok := c.Get(XUserInfoTag); ok {

		c.JSON(200, SUCCESS.Tr(ZH).With(info))
		return
	}

	c.JSON(404, ErrNoFound.Tr(ZH))

}

func UserMessagesHandler(c *gin.Context) {
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
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	// RecaptchaToken string `form:"recaptcha_token" json:"recaptcha_token"`
}

func SignInHandler(c *gin.Context) {

	// ip := c.ClientIP()
	req := &struct {
		Account  string `form:"account" json:"account" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
		// RecaptchaToken string `form:"recaptcha_token" json:"recaptcha_token"`
	}{}

	if err := c.ShouldBind(req); err != nil {
		zlog.I(`用户登录参数异常`, err)
		c.JSON(200, ErrUserPassword.Tr(ZH))
		return
	}

	user, err := model.FirstUserByAccount(srv.orm, req.Account)

	if err != nil || !user.VerifyPassword(req.Password) {
		zlog.E(`查询用户失败：`, err)
		c.JSON(200, ErrUserPassword.Tr(ZH))
		return
	}

	if user.Locked {
		zlog.E(`用户或者客户端已经被锁定：`, err)
		c.JSON(200, ErrUserLocked.Tr(ZH))
		return
	}

	sess, err := user.CreateSession(srv.orm, c.GetHeader(`User-Agent`), c.GetHeader(`Referer`), c.ClientIP())
	if err != nil {
		zlog.E(`创建session失败：`, err)
		c.JSON(200, ErrInnerServer.Tr(ZH))
		return
	}

	zlog.I(`用户登录成功设置，浏览器cookies,Host: `, c.Request.Host)
	c.SetCookie(UserSessionTag, sess.Session, 1*24*3600, "/", getHostName(c), false, true)
	//TODO recording login log to DB
	c.JSON(200, SUCCESS.Tr(ZH).With(user.Masking()))
}

// SignOutHandler 用户登出
func SignOutHandler(c *gin.Context) {

	session, err := getSession(c)

	if err != nil {
		zlog.E(`获取Session 失败：`, err.Error())
		c.JSON(401, ErrSession.Tr(ZH))
		c.Abort()
		return
	}

	zlog.I(`用户登出：`, session.UserInfo.NickName)

	session.Clean(srv.Store())

	//清除cookies
	c.SetCookie(UserSessionTag, "", -1, "/", getHostName(c), true, true)

	c.JSON(200, SUCCESS.Tr(ZH))
}

func UserUpdateHandler(c *gin.Context) {

}
func UserResetHandler(c *gin.Context) {

}
func UserSendActiveHandler(c *gin.Context) {

}
