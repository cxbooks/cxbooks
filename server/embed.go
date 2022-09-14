package server

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed dist
var vue embed.FS

//go:embed dist/index.html
var vueIndex []byte

func embedVue(router *gin.Engine) {

	vuefs := mustFS(vue, `dist`)

	router.StaticFS("/ui", vuefs)

	// TODO 所有匹配不到API路由的请求将请求到 vue web
	// vue 生成的静态文件中只有一个 index.html 文件。vue中的路由本质上来说也是返回了index.html. js根据不用的router 返回了不同的内容
	// 这里暂时使用这种方式返回,这里有个缺陷就是可能吧其他未知路由路径重定向到了 index.html, 所以需要在 vue 代码中 处理404问题。
	//
	router.NoRoute(func(c *gin.Context) {

		if c.Request.Method == http.MethodGet &&
			!strings.ContainsRune(c.Request.URL.Path, '.') &&
			!strings.HasPrefix(c.Request.URL.Path, "/api/") {

			c.Header("content-type", "text/html;charset=utf-8")
			c.String(200, string(vueIndex))
		} else {
			c.JSON(404, ErrPageNotFound.Tr(ZH))
		}

	})
}

func mustFS(f fs.FS, path string) http.FileSystem {
	sub, err := fs.Sub(f, path)

	if err != nil {
		panic(err)
	}

	return http.FS(sub)
}
