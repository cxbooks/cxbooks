package server

import (
	"net"
	"time"

	"github.com/cxbooks/cxbooks/server/model"
	"github.com/cxbooks/cxbooks/server/zlog"
	"github.com/gin-gonic/gin"
)

const (
	XUserOauthListTag = `X-User-Oauth-list`
	XUserOauthTag     = `X-User-List-Tag`
	XUserInfoTag      = `X-User-Info-Key`
	XUserIDTag        = `X-User-ID-Key`
	XLangTag          = `x-Lang-Tag`

	UserSessionTag = `session-id`
)

// OauthMiddleware 认证中间件，如果seesion 合法将用户ID存入上下文
func OauthMiddleware(c *gin.Context) {

	session, err := getSession(c)

	if err != nil {
		zlog.E(`获取Session 失败：`, err.Error())
		c.JSON(401, ErrSession.Tr(ZH))
		c.Abort()
		return
	}

	//校验session 合法性
	//获取相对时间
	// diff := (time.Now().Unix() - session.Utime)
	diff := time.Since(session.UTime) / 1000 / 1000 / 1000

	if diff-session.Duration > 0 {

		zlog.E(`session has expired `, session.Session)
		// 删除
		c.SetCookie(UserSessionTag, "", -1, "/", "", true, true)
		session.Clean(srv.orm)
		c.JSON(401, ErrSession.Tr(ZH))
		c.Abort()
		return
	}

	//session 有效时间小于1分钟 刷新session 过去事情
	if diff > 60 {

		if err := srv.Store().Table(`sessions`).Where(`id = ?`, session.Session).Update(`utime`, time.Now()).Error; err != nil {
			zlog.E("刷新session失败: ", err)
		}
		zlog.D("刷新session")

	}

	c.Set(XUserInfoTag, session.UserInfo)
	c.Set(XUserIDTag, session.UserInfo.UserID)

	c.Next()
}


// parseSessionUser 获取session 关联的用户信息，
func getSession(c *gin.Context) (*model.Session, error) {
	id, err := c.Cookie(UserSessionTag)
	if err != nil {
		return nil, err
	}

	return model.FirstSessionByID(srv.orm, id)

}

// func CORSMiddleware(c *gin.Context) {

// 	c.Header("Access-Control-Allow-Origin", "192.168.1.201")
// 	// c.Header("Access-Control-Allow-Credentials", "true")
// 	// c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 	// c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

// 	if c.Request.Method == "OPTIONS" {
// 		c.AbortWithStatus(204)
// 		return
// 	}

// 	c.Next()
// }

func getHostName(c *gin.Context) string {
	host, _, err := net.SplitHostPort(c.Request.Host)

	if err != nil {
		zlog.I(`区分端口失败： `, c.Request.Host)
		host = c.Request.Host
	}

	return host
}
