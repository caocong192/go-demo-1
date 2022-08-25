package server

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/go-demo-1/controller"
	"github.com/go-demo-1/define"
	"io/fs"
	"log"
	"net/http"
	"strings"
)


//go:embed frontend/dist/*
var FS embed.FS


// Run 启动gin
func Run()  {
	r := gin.Default()

	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	r.StaticFS("/static", http.FS(staticFiles))

	r.POST("/api/v1/texts", controller.TextsController)
	r.GET("/api/v1/addresses", controller.AddressesController)
	r.GET("/uploads/:path",  controller.DownloadController)
	r.GET("/api/v1/qrcodes", controller.QrcodesController)
	r.POST("/api/v1/files", controller.FilesController)

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/static/") {
			reader, err := staticFiles.Open("index.html")
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()
			stat, err := reader.Stat()
			if err != nil {
				log.Fatal(err)
			}
			c.DataFromReader(http.StatusOK, stat.Size(), "text/html", reader, nil)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	r.Run(":" + define.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
