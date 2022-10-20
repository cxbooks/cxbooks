package server

import (
	"os"

	"github.com/cxbooks/cxbooks/server/zlog"
	"github.com/gin-gonic/gin"
)

func AdminSSL(c *gin.Context) {

}
func AdminUsers(c *gin.Context) {

}
func AdminInstall(c *gin.Context) {

}

func AdminSettings(c *gin.Context) {

}

func AdminTestMail(c *gin.Context) {

}

func AdminBookList(c *gin.Context) {

}

func ScanList(c *gin.Context) {

}

// ScanRun 执行目录扫描
func ScanRun(c *gin.Context) {

	if srv.scanner.IsRunning() {
		zlog.I(`扫描未完成，放弃执行...`)
		c.JSON(200, ErrScannerIsRunning.Tr(ZH))
		return
	}

	req := &struct {
		Path      string `json:"path" binding:"required,startswith=/data/ebooks"`
		MaxThread int    `json:"max_thread" binding:"required,gte=1"`
	}{}

	if err := c.ShouldBind(req); err != nil {
		zlog.I(`用户登录参数异常`, err)
		c.JSON(200, ErrReqArgs.Tr(ZH).With(err.Error()))
		return
	}

	_, err := os.Stat(req.Path)
	if os.IsNotExist(err) {
		// path/to/whatever does not exist
		zlog.D(`扫描目录`, req.Path, `无法访问或者不存在 `, err)
		c.JSON(200, ErrScanPathNotExist.Tr(ZH))
		return
	}

	srv.scanner.Start(req.Path, req.MaxThread)

	c.JSON(200, SUCCESS.Tr(ZH))

}

// ScanStatus 获取当前扫描状态
func ScanStatus(c *gin.Context) {
	zlog.I(`获取当前扫描状态数据`)
	states := srv.scanner.Status()
	c.JSON(200, SUCCESS.Tr(ZH).With(states))

}

// ScanDelete 停止当前扫描
func ScanDelete(c *gin.Context) {

	zlog.I(`停止当前扫描`)
	srv.scanner.Stop()

	c.JSON(200, SUCCESS.Tr(ZH))

}

func ScanMark(c *gin.Context) {

}
