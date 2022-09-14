// router 整个网站的路由
package server

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func RegRoute(router *gin.Engine) {

	admin := router.Group("/api/admin", OauthMiddleware)

	admin.POST("/ssl", AdminSSL)
	admin.POST("/users", AdminUsers)
	admin.POST("/install", AdminInstall)
	admin.POST("/settings", AdminSettings)
	admin.POST("/testmail", AdminTestMail)
	admin.POST("/book/list", AdminBookList)

	admin.POST("/api/admin/scan/list", ScanList)
	admin.POST("/api/admin/scan/run", ScanRun)
	admin.POST("/api/admin/scan/status", ScanStatus)
	admin.POST("/api/admin/scan/delete", ScanDelete)
	admin.POST("/api/admin/scan/mark", ScanMark)
	admin.POST("/api/admin/import/run", ImportRun)
	admin.POST("/api/admin/import/status", ImportStatus)

	// (r"/api/welcome", Welcome),

	router.POST("api/user/login", SignInHandler)

	user := router.Group("/api/user", OauthMiddleware)

	user.GET("/info", UserInfoHandler)
	user.GET("/messages", UserMessagesHandler)
	user.POST("/logout", SignOutHandler)
	// user.POST("/sign_up", SignUp)
	user.POST("/update", UserUpdateHandler)
	user.POST("/reset", UserResetHandler)
	user.POST("/active/send", UserSendActiveHandler)

	// (r"/api/active/(.*)/(.*)", UserActive),
	// (r"/api/done/", Done),

	// opds := router.Group("/opds")

	// opds.POST("/", OpdsIndex)
	// opds.POST("/nav/(.*)", OpdsNav)
	// opds.POST("/category/(.*)/(.*)", OpdsCategory)
	// opds.POST("/categorygroup/(.*)/(.*)", OpdsCategoryGroup)
	// opds.POST("/search/(.*)", OpdsSearch)

	// router.GET("/api/(author|publisher|tag|rating|series)", MetaList)
	// router.GET("/api/(author|publisher|tag|rating|series)/(.*)", MetaBooks)
	// router.GET("/api/author/(.*)/update", AuthorBooksUpdate)
	// router.GET("/api/publisher/(.*)/update", PubBooksUpdate)

	router.GET("/api/book/index", OauthMiddleware, Index)
	router.GET("/api/book/search", OauthMiddleware, SearchBook)
	router.GET("/api/book/recent", OauthMiddleware, RecentBook)
	router.GET("/api/book/hot", OauthMiddleware, HotBook)
	router.GET("/api/book/nav", OauthMiddleware, BookNav)
	router.GET("/api/book/upload", OauthMiddleware, BookUpload)
	router.GET("/api/books/:book_id", OauthMiddleware, BookDetail)
	router.GET("/api/books/:book_id/delete", OauthMiddleware, BookDelete)
	router.GET("/api/books/:book_id/edit", OauthMiddleware, BookEdit)
	router.GET(`/api/books/:book_id/download`, OauthMiddleware, BookDownload)
	router.GET("/api/books/:book_id/push", OauthMiddleware, BookPush)
	router.GET("/api/books/:book_id/refer", OauthMiddleware, BookRefer)
	router.GET("/read/:book_id", OauthMiddleware, BookRead)

	//  (r"/get/pcover", ProxyImageHandler),
	//     (r"/get/progress/([0-9]+)", ProgressHandler),
	//     (r"/get/extract/(.*)", web.StaticFileHandler, {"path": CONF["extract_path"]}),
	//     (r"/get/(.*)/(.*)", ImageHandler),
	//     (r"/(.*)", web.StaticFileHandler, static_config),
}

func initGinRoute(level zapcore.Level) *gin.Engine {

	if level == zapcore.DebugLevel {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// log.SetFlags(log.LstdFlags) // gin will disable log flags

	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	

	embedVue(router)

	RegRoute(router)

	return router
}
