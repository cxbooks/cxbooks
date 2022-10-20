package server

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cxbooks/cxbooks/server/model"
	"github.com/cxbooks/cxbooks/server/zlog"
	"github.com/gin-gonic/gin"
)

// UserInfo GET /api/user/info
// 获取用户信息
func GetBookStats(c *gin.Context) {
	zlog.D(`获取书籍概览信息`)

	stats, err := model.GetBookStats(srv.Store())

	if err != nil {
		zlog.E(`获取数据信息失败`)
		c.JSON(404, ErrNoFound.Tr(ZH))
	}

	c.JSON(200, SUCCESS.Tr(ZH).With(stats))

}

func MetaList(c *gin.Context) {

}
func MetaBooks(c *gin.Context) {

}

func AuthorBooksUpdate(c *gin.Context) {

}

func PubBooksUpdate(c *gin.Context) {

}

// Index 返回 首页 随机ID书籍和最近添加到书籍集。
func Index(c *gin.Context) {

	req := &struct {
		Random int `json:"random" binding:"required,gte=0,lt=30"`
		Recent int `json:"recent" binding:"required,gte=0,lt=30"`
	}{
		Random: 10,
		Recent: 12,
	}

	if err := c.ShouldBind(req); err != nil {
		zlog.I(`用户登录参数异常`, err)
		c.JSON(200, ErrReqArgs.Tr(ZH).With(err.Error()))
		return
	}

	randBooks, err := model.RandomBooks(srv.orm, req.Random)

	if err != nil {
		zlog.I(`用户登录参数异常`, err)
		c.JSON(200, ErrInnerServer.Tr(ZH).With(err.Error()))
		return
	}

	recentBooks, err := model.RecentBooks(srv.orm, req.Recent)

	if err != nil {
		zlog.I(`用户登录参数异常`, err)
		c.JSON(200, ErrInnerServer.Tr(ZH).With(err.Error()))
		return
	}

	data := map[string]interface{}{
		"random": randBooks.Books,
		"recent": recentBooks.Books,
	}

	c.JSON(200, SUCCESS.Tr(ZH).With(data))

}

// SearchBook 模糊搜索
func SearchBook(c *gin.Context) {

	req := &model.Query{}

	if err := c.ShouldBind(req); err != nil {
		zlog.I(`用户登录参数异常`, err)
		c.JSON(200, ErrReqArgs.Tr(ZH).With(err.Error()))
		return
	}

	//TODO过滤特殊字符

	recentBooks, err := model.SearchBooks(srv.orm, req)

	if err != nil {
		zlog.I(`用户登录参数异常`, err)
		c.JSON(200, ErrInnerServer.Tr(ZH).With(err.Error()))
		return
	}

	c.JSON(200, SUCCESS.Tr(ZH).With(recentBooks))

}

func RecentBook(c *gin.Context) {
	req := &struct {
		Recent int `json:"recent" binding:"required,gte=0"`
	}{
		Recent: 12,
	}

	if err := c.ShouldBind(req); err != nil {
		zlog.I(`用户登录参数异常`, err)
		c.JSON(200, ErrReqArgs.Tr(ZH).With(err.Error()))
		return
	}
	recentBooks, err := model.RecentBooks(srv.orm, req.Recent)

	if err != nil {
		zlog.I(`用户登录参数异常`, err)
		c.JSON(200, ErrInnerServer.Tr(ZH).With(err.Error()))
		return
	}

	c.JSON(200, SUCCESS.Tr(ZH).With(recentBooks))

}

func HotBook(c *gin.Context) {

}

func BookNav(c *gin.Context) {

}

func BookUpload(c *gin.Context) {

}

func BookDetail(c *gin.Context) {
	bookID := c.Param(`book_id`)
	id, err := strconv.Atoi(bookID)
	if err != nil {
		zlog.I(`用户登录参数异常`, err)
		c.JSON(200, ErrInnerServer.Tr(ZH).With(err.Error()))
		return
	}

	book, err := model.FirstBookByID(srv.orm, id)
	if err != nil {
		zlog.I(`用户登录参数异常`, err)
		c.JSON(200, ErrInnerServer.Tr(ZH).With(err.Error()))
		return
	}

	c.JSON(200, SUCCESS.Tr(ZH).With(book))

}

func BookDelete(c *gin.Context) {

}

func BookEdit(c *gin.Context) {

}

// BookDownload 下载书
func BookDownload(c *gin.Context) {

	bookID := c.Param(`book_id`)

	id, err := strconv.Atoi(bookID)

	if err != nil {
		zlog.I(`用户登录参数异常`, err)
		c.JSON(200, ErrInnerServer.Tr(ZH).With(err.Error()))
		return
	}

	book, err := model.FirstBookByID(srv.orm, id)
	if err != nil {
		zlog.I(`用户登录参数异常`, err)
		c.JSON(200, ErrInnerServer.Tr(ZH).With(err.Error()))
		return
	}

	if _, err := os.Stat(book.Path); os.IsNotExist(err) {
		// path/to/whatever does not exist
		zlog.E(`图书路径不存在: `, book.Path)
		os.Exit(1)
	}

	c.FileAttachment(book.Path, book.Title)

}
func BookPush(c *gin.Context) {

}
func BookRefer(c *gin.Context) {

}

func BookRead(c *gin.Context) {

}

// ProxyImageHandler 获取图书封面接口
func ProxyImageHandler(c *gin.Context) {

	coverPath := c.Param(`book_cover_path`)

	ext := filepath.Ext(c.Param(`book_cover_path`))

	if ext != ".png" && ext != ".jpeg" && ext != ".jpg" {
		zlog.E(`图片路径ID为空：`, ext)
		c.JSON(200, ErrReqArgs.Tr(ZH))
		return
	}

	image, err := srv.scanner.GetCoverData(filepath.Join(`/book/cover`, coverPath))
	if err != nil {
		zlog.E(`获取图片内容失败`, err.Error())
		c.JSON(http.StatusNotFound, ErrNoFound.Tr(ZH))
		return
	}

	c.Data(200, "image/"+ext, image)
}
