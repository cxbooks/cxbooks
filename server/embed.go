package server

import (
	"embed"
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//go:embed dist
var vue embed.FS

func embedVue(router *gin.Engine) {

	handler := VueHandler()

	router.HEAD("/*filepath", handler)

	//所有匹配不到API路由的请求将请求到 vue web
	router.NoRoute(func(c *gin.Context) {

		if c.Request.Method != `GET` {
			noRouteHandler(c)
			c.Abort()
			return
		}

		handler(c)

	})
}

func VueHandler() gin.HandlerFunc {

	vuefs := mustFS(vue, `dist`)

	fileServer := http.StripPrefix("/", http.FileServer(vuefs))

	return func(c *gin.Context) {
		if _, noListing := vuefs.(*onlyFilesFS); noListing {
			c.Writer.WriteHeader(http.StatusNotFound)
		}

		file := c.Request.URL.Path
		// Check if file exists and/or if we have permission to access it
		f, err := vuefs.Open(file)
		if err != nil {
			noRouteHandler(c)

			c.Abort()
			return
		}
		f.Close()

		fileServer.ServeHTTP(c.Writer, c.Request)
	}

}

func noRouteHandler(c *gin.Context) {
	c.JSON(404, ErrPageNotFound.Tr(ZH))
}

func mustFS(f fs.FS, path string) http.FileSystem {
	sub, err := fs.Sub(f, path)

	if err != nil {
		panic(err)
	}

	return http.FS(sub)
}

type onlyFilesFS struct {
	fs http.FileSystem
}

type neuteredReaddirFile struct {
	http.File
}

// Dir returns a http.FileSystem that can be used by http.FileServer(). It is used internally
// in router.Static().
// if listDirectory == true, then it works the same as http.Dir() otherwise it returns
// a filesystem that prevents http.FileServer() to list the directory files.
func Dir(root string, listDirectory bool) http.FileSystem {
	fs := http.Dir(root)
	if listDirectory {
		return fs
	}
	return &onlyFilesFS{fs}
}

// Open conforms to http.Filesystem.
func (fs onlyFilesFS) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

// Readdir overrides the http.File default implementation.
func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	// this disables directory listing
	return nil, nil
}
