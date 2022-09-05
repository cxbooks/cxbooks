// router 整个网站的路由
package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func RegRoute(router *gin.Engine) {

	admin := router.Group("/api/admin")

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

	user := router.Group("/api/user")

	user.GET("/info", UserInfo)
	user.GET("/messages", UserMessages)
	user.POST("/sign_in", SignIn)
	user.POST("/sign_out", SignOut)
	// user.POST("/sign_up", SignUp)

	user.POST("/update", UserUpdate)
	user.POST("/reset", UserReset)
	user.POST("/active/send", UserSendActive)

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

	router.GET("/api/book/index", Index)
	router.GET("/api/book/search", SearchBook)
	router.GET("/api/book/recent", RecentBook)
	router.GET("/api/book/hot", HotBook)
	router.GET("/api/book/nav", BookNav)
	router.GET("/api/book/upload", BookUpload)
	router.GET("/api/books/:book_id", BookDetail)
	router.GET("/api/books/:book_id/delete", BookDelete)
	router.GET("/api/books/:book_id/edit", BookEdit)
	router.GET(`/api/books/:book_id/download`, BookDownload)
	router.GET("/api/books/:book_id/push", BookPush)
	router.GET("/api/books/:book_id/refer", BookRefer)
	router.GET("/read/:book_id", BookRead)

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

	log.SetFlags(log.LstdFlags) // gin will disable log flags

	router := gin.New()
	router.Use(gin.Recovery())

	RegRoute(router)

	return router
}
