package server

import "github.com/gin-gonic/gin"

func UserInfo(c *gin.Context) {

}

func UserMessages(c *gin.Context) {

}

// LoginReq  登陆
type LoginReq struct {
	Account        string `form:"account" json:"account" binding:"required"`
	Password       string `form:"password" json:"password" binding:"required"`
	RecaptchaToken string `form:"recaptcha_token" json:"recaptcha_token"`
}

func SignIn(c *gin.Context) {

	// ip := c.ClientIP()

	req := &LoginReq{}

	if err := c.ShouldBind(req); err != nil {
		c.JSON(400, ErrUserPassword.Tr(ZH))
		return
	}

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
